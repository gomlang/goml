package adapter

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func response(status int, body string) Response {
	return NewResponse(status, nil, nil, []byte(body), nil)
}
func start(t *testing.T, handler func(Request) Response, timeout time.Duration, bodyLimit int64, concurrent int) Server {
	t.Helper()
	server, failure := Start("127.0.0.1:0", handler, int64(timeout), bodyLimit, 4096, concurrent, int64(time.Second))
	if failure.Code != 0 {
		t.Fatal(failure)
	}
	t.Cleanup(func() { _ = Shutdown(server, int64(time.Second)) })
	return server
}
func get(t *testing.T, address string) (int, string) {
	t.Helper()
	client := &http.Client{Timeout: 2 * time.Second}
	reply, err := client.Get(address)
	if err != nil {
		t.Fatal(err)
	}
	defer reply.Body.Close()
	body, err := io.ReadAll(reply.Body)
	if err != nil {
		t.Fatal(err)
	}
	return reply.StatusCode, string(body)
}

func TestCaptureSnapshotsAndPanicRecovery(t *testing.T) {
	payload := []byte("original")
	ready := NewResponse(201, []string{"X-Many", "X-Many"}, []string{"one", "two"}, payload, nil)
	payload[0] = 'X'
	captured, failure := Dispatch(func(Request) Response { return ready }, "GET", "/", nil, nil, nil, int64(time.Second), 10, 100)
	if failure.Code != 0 || captured.status != 201 || string(captured.body) != "original" || len(captured.headers.Values("X-Many")) != 2 {
		t.Fatal(captured, failure)
	}
	copy := CaptureBody(captured)
	copy[0] = 'X'
	if string(CaptureBody(captured)) != "original" {
		t.Fatal("capture aliases body")
	}
	_, failure = Dispatch(func(Request) Response { panic("secret") }, "GET", "/", nil, nil, nil, int64(time.Second), 10, 100)
	if failure.Code != 8 {
		t.Fatal(failure)
	}
	server := start(t, func(Request) Response { panic("private information") }, time.Second, 100, 2)
	status, body := get(t, "http://"+Address(server)+"/")
	if status != 500 || strings.Contains(body, "private") {
		t.Fatal(status, body)
	}
}

func TestStreamingArrivesBeforeProducerCompletes(t *testing.T) {
	release := make(chan struct{})
	entered := make(chan struct{})
	server := start(t, func(Request) Response {
		return NewResponse(200, []string{"Content-Type"}, []string{"text/event-stream"}, nil, func(writer Writer) Failure {
			if failure := Write(writer, []byte("data: first\n\n"), true); failure.Code != 0 {
				return failure
			}
			close(entered)
			<-release
			return Write(writer, []byte("data: second\n\n"), true)
		})
	}, time.Second, 100, 2)
	client := &http.Client{Timeout: 2 * time.Second}
	reply, err := client.Get("http://" + Address(server) + "/")
	if err != nil {
		close(release)
		t.Fatal(err)
	}
	defer reply.Body.Close()
	reader := bufio.NewReader(reply.Body)
	first, err := reader.ReadString('\n')
	if err != nil || first != "data: first\n" {
		close(release)
		t.Fatal(first, err)
	}
	<-entered
	close(release)
	rest, err := io.ReadAll(reader)
	if err != nil || string(rest) != "\ndata: second\n\n" {
		t.Fatal(string(rest), err)
	}
}

func TestStreamingClientDisconnectCancelsProducer(t *testing.T) {
	done := make(chan Failure, 1)
	server := start(t, func(request Request) Response {
		return NewResponse(200, nil, nil, nil, func(writer Writer) Failure {
			if failure := Write(writer, []byte("first\n"), true); failure.Code != 0 {
				done <- failure
				return failure
			}
			failure := Wait(request, int64(5*time.Second))
			done <- failure
			return failure
		})
	}, 10*time.Second, 100, 2)
	reply, err := http.Get("http://" + Address(server) + "/")
	if err != nil {
		t.Fatal(err)
	}
	first := make([]byte, 6)
	if _, err := io.ReadFull(reply.Body, first); err != nil {
		t.Fatal(err)
	}
	_ = reply.Body.Close()
	select {
	case failure := <-done:
		if failure.Code != 4 {
			t.Fatal(failure)
		}
	case <-time.After(time.Second):
		t.Fatal("disconnect did not cancel request")
	}
}

func TestStreamingBackpressureDoesNotBufferUnboundedly(t *testing.T) {
	var produced atomic.Int64
	done := make(chan struct{})
	server := start(t, func(Request) Response {
		return NewResponse(200, nil, nil, nil, func(writer Writer) Failure {
			defer close(done)
			block := bytes.Repeat([]byte("x"), 65536)
			for i := 0; i < 32768; i++ {
				if failure := Write(writer, block, true); failure.Code != 0 {
					return failure
				}
				produced.Add(int64(len(block)))
			}
			return Success()
		})
	}, 2*time.Second, 100, 2)
	conn, err := net.Dial("tcp", Address(server))
	if err != nil {
		t.Fatal(err)
	}
	if tcp, ok := conn.(*net.TCPConn); ok {
		_ = tcp.SetReadBuffer(1024)
	}
	_, _ = io.WriteString(conn, "GET / HTTP/1.1\r\nHost: localhost\r\n\r\n")
	deadline := time.Now().Add(time.Second)
	for produced.Load() == 0 && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	time.Sleep(30 * time.Millisecond)
	if count := produced.Load(); count == 0 || count > 32<<20 {
		t.Fatalf("unexpected buffered output: %d", count)
	}
	_ = conn.Close()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("blocked writer was not cancelled")
	}
}

func TestChunkedRequestLimitAndReadDeadline(t *testing.T) {
	server := start(t, func(request Request) Response {
		var result []byte
		for {
			chunk, failure := Read(request, 4)
			if failure.Code != 0 {
				if failure.Code == 3 {
					return response(413, "limit")
				}
				return response(400, "read")
			}
			if len(chunk) == 0 {
				break
			}
			result = append(result, chunk...)
		}
		return response(200, string(result))
	}, 50*time.Millisecond, 5, 2)
	conn, err := net.Dial("tcp", Address(server))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	_, _ = io.WriteString(conn, "POST / HTTP/1.1\r\nHost: localhost\r\nTransfer-Encoding: chunked\r\n\r\n3\r\nabc\r\n3\r\ndef\r\n0\r\n\r\n")
	reply, err := http.ReadResponse(bufio.NewReader(conn), nil)
	if err != nil {
		t.Fatal(err)
	}
	defer reply.Body.Close()
	if reply.StatusCode != 413 {
		t.Fatal(reply.StatusCode)
	}
	slow, err := net.Dial("tcp", Address(server))
	if err != nil {
		t.Fatal(err)
	}
	defer slow.Close()
	_ = slow.SetDeadline(time.Now().Add(time.Second))
	_, _ = io.WriteString(slow, "POST / HTTP/1.1\r\nHost: localhost\r\nContent-Length: 4\r\n\r\na")
	reply, err = http.ReadResponse(bufio.NewReader(slow), nil)
	if err == nil {
		defer reply.Body.Close()
		if reply.StatusCode != 504 {
			t.Fatal(reply.StatusCode)
		}
	}
}

func TestConcurrencyLimitAndForcedShutdown(t *testing.T) {
	entered := make(chan struct{}, 1)
	done := make(chan Failure, 1)
	server := start(t, func(request Request) Response {
		entered <- struct{}{}
		failure := Wait(request, int64(10*time.Second))
		done <- failure
		return response(200, "done")
	}, 20*time.Second, 100, 1)
	clientDone := make(chan struct{})
	go func() {
		defer close(clientDone)
		client := &http.Client{Timeout: 2 * time.Second}
		reply, _ := client.Get("http://" + Address(server))
		if reply != nil {
			_ = reply.Body.Close()
		}
	}()
	<-entered
	status, _ := get(t, "http://"+Address(server))
	if status != 503 {
		t.Fatal(status)
	}
	before := time.Now()
	failure := Shutdown(server, int64(20*time.Millisecond))
	if failure.Code != 5 || time.Since(before) > time.Second {
		t.Fatal(failure, time.Since(before))
	}
	select {
	case failure = <-done:
		if failure.Code != 4 {
			t.Fatal(failure)
		}
	case <-time.After(time.Second):
		t.Fatal("shutdown did not cancel handler")
	}
	<-clientDone
	if !IsClosed(server) {
		t.Fatal("server not closed")
	}
}

func TestConcurrentShutdownAndGracefulInflightCompletion(t *testing.T) {
	entered := make(chan struct{})
	release := make(chan struct{})
	server := start(t, func(Request) Response { close(entered); <-release; return response(200, "done") }, time.Second, 100, 2)
	clientDone := make(chan string, 1)
	go func() {
		reply, err := http.Get("http://" + Address(server))
		if err != nil {
			clientDone <- err.Error()
			return
		}
		defer reply.Body.Close()
		body, err := io.ReadAll(reply.Body)
		if err != nil {
			clientDone <- err.Error()
			return
		}
		clientDone <- string(body)
	}()
	<-entered
	failures := make(chan Failure, 8)
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); failures <- Shutdown(server, int64(time.Second)) }()
	}
	deadline := time.Now().Add(time.Second)
	for !IsClosed(server) && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	close(release)
	wg.Wait()
	close(failures)
	for failure := range failures {
		if failure.Code != 0 {
			t.Fatal(failure)
		}
	}
	if body := <-clientDone; body != "done" {
		t.Fatal(body)
	}
}

func TestCancellationDeadlineCleanupDoesNotPoisonKeepalive(t *testing.T) {
	var calls atomic.Int64
	server := start(t, func(request Request) Response {
		calls.Add(1)
		if RequestPath(request) == "/timeout" {
			_ = Wait(request, int64(100*time.Millisecond))
		}
		return response(200, "alive")
	}, 10*time.Millisecond, 100, 4)
	transport := &http.Transport{MaxConnsPerHost: 1}
	defer transport.CloseIdleConnections()
	client := &http.Client{Transport: transport, Timeout: time.Second}
	for i := 0; i < 25; i++ {
		reply, err := client.Get("http://" + Address(server) + "/timeout")
		if err == nil {
			_, _ = io.Copy(io.Discard, reply.Body)
			_ = reply.Body.Close()
		}
		reply, err = client.Get("http://" + Address(server) + "/healthy?i=" + strconv.Itoa(i))
		if err != nil {
			t.Fatal(i, err)
		}
		body, err := io.ReadAll(reply.Body)
		_ = reply.Body.Close()
		if err != nil || reply.StatusCode != 200 || string(body) != "alive" {
			t.Fatal(i, reply.StatusCode, string(body), err)
		}
	}
}

func TestEscapedHandlesAndConcurrentRead(t *testing.T) {
	var saved Request
	var sink Writer
	captured, failure := Dispatch(func(request Request) Response {
		saved = request
		var wg sync.WaitGroup
		chunks := make(chan []byte, 20)
		for i := 0; i < 20; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				chunk, err := Read(request, 1)
				if err.Code != 0 {
					t.Error(err)
				}
				chunks <- chunk
			}()
		}
		wg.Wait()
		close(chunks)
		count := 0
		for chunk := range chunks {
			count += len(chunk)
		}
		if count != 20 {
			t.Error(count)
		}
		return NewResponse(200, nil, nil, nil, func(writer Writer) Failure { sink = writer; return Write(writer, []byte("ok"), false) })
	}, "POST", "/", nil, nil, bytes.Repeat([]byte("x"), 20), int64(time.Second), 100, 100)
	if failure.Code != 0 || string(captured.body) != "ok" {
		t.Fatal(failure)
	}
	if Check(saved).Code != 7 || WriterCheck(sink).Code != 7 || Write(sink, []byte("late"), false).Code != 7 {
		t.Fatal("escaped handle remains active")
	}
}

func TestCancelledContextAndCaptureBounds(t *testing.T) {
	_, failure := Dispatch(func(Request) Response {
		return NewResponse(200, nil, nil, nil, func(writer Writer) Failure { return Write(writer, bytes.Repeat([]byte("x"), 11), false) })
	}, "GET", "/", nil, nil, nil, int64(time.Second), 100, 10)
	if failure.Code != 3 {
		t.Fatal(failure)
	}
	_, failure = Dispatch(func(Request) Response {
		return NewResponse(200, nil, nil, nil, func(writer Writer) Failure { return Write(writer, bytes.Repeat([]byte("x"), 65537), false) })
	}, "GET", "/", nil, nil, nil, int64(time.Second), 100, 100000)
	if failure.Code != 1 {
		t.Fatal(failure)
	}
	server := start(t, func(request Request) Response { _ = Wait(request, int64(time.Second)); return response(200, "late") }, time.Second, 100, 2)
	ctx, cancel := context.WithCancel(context.Background())
	request, _ := http.NewRequestWithContext(ctx, "GET", "http://"+Address(server), nil)
	cancel()
	if _, err := http.DefaultClient.Do(request); err == nil {
		t.Fatal("cancelled request unexpectedly succeeded")
	}
}
