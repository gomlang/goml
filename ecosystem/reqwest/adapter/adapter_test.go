package adapter

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"
)

func testClient(t *testing.T) Client {
	t.Helper()
	client, problem := NewClient(1000, 1000, 8, 4096, true, false, "", "", "", "", false)
	if problem != nil {
		t.Fatal(FailureMessage(problem))
	}
	t.Cleanup(func() { Close(client) })
	return client
}

func TestCancellationAndCloseInterruptActiveRequests(t *testing.T) {
	for _, closeClient := range []bool{false, true} {
		for iteration := 0; iteration < 12; iteration++ {
			entered := make(chan struct{})
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { close(entered); <-r.Context().Done() }))
			client := testClient(t)
			control := NewControl(int64(5 * time.Second))
			done := make(chan Failure, 1)
			go func() { _, problem := Do(client, control, "GET", server.URL, nil, nil, nil, 100); done <- problem }()
			select {
			case <-entered:
			case <-time.After(3 * time.Second):
				t.Fatal("request did not start")
			}
			if closeClient {
				Close(client)
			} else {
				Cancel(control)
			}
			select {
			case problem := <-done:
				if FailureKind(problem) != 3 {
					t.Fatalf("expected cancellation: %v", problem)
				}
			case <-time.After(3 * time.Second):
				t.Fatal("request did not stop")
			}
			Cancel(control)
			Close(client)
			server.Close()
		}
	}
}

func TestHeadAndBodyBounds(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "4096")
		if r.Method != "HEAD" {
			io.WriteString(w, strings.Repeat("x", 4096))
		}
	}))
	defer server.Close()
	client := testClient(t)
	control := NewControl(int64(5 * time.Second))
	defer Cancel(control)
	response, problem := Do(client, control, "HEAD", server.URL, nil, nil, nil, 0)
	if problem != nil || len(ResponseBody(response)) != 0 {
		t.Fatalf("HEAD should ignore representation length: %v", problem)
	}
	_, problem = Do(client, control, "GET", server.URL, nil, nil, nil, 10)
	if FailureKind(problem) != 5 {
		t.Fatalf("body limit: %v", problem)
	}
}

func TestConcurrentReuseAndClose(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "ok") }))
	defer server.Close()
	client := testClient(t)
	var workers sync.WaitGroup
	for i := 0; i < 32; i++ {
		workers.Add(1)
		go func() {
			defer workers.Done()
			control := NewControl(int64(5 * time.Second))
			defer Cancel(control)
			for j := 0; j < 8; j++ {
				response, problem := Do(client, control, "GET", server.URL, nil, nil, nil, 10)
				if problem == nil && !bytes.Equal(ResponseBody(response), []byte("ok")) {
					t.Error("invalid response")
				}
				if problem != nil && FailureKind(problem) != 3 && FailureKind(problem) != 7 {
					t.Errorf("unexpected error %v", problem)
				}
			}
		}()
	}
	CloseIdle(client)
	Close(client)
	Close(client)
	workers.Wait()
	if !IsClosed(client) {
		t.Fatal("client remains open")
	}
}

func TestURLAndMultipartValidation(t *testing.T) {
	for _, text := range []string{"file:///tmp/file", "http://user:pass@example.com", "/relative", "https://"} {
		if _, _, _, problem := ParseURL(text, ""); problem == nil {
			t.Fatalf("accepted %q", text)
		}
	}
	text, origin, secure, problem := ParseURL("../next", "https://EXAMPLE.com/base/file")
	if problem != nil || text != "https://example.com/next" || origin != "https://example.com:443" || !secure {
		t.Fatalf("URL normalization: %s %s", text, origin)
	}
	source := []byte("original")
	part := NewPart("file", "data.bin", "application/octet-stream", source, true)
	source[0] = 'x'
	body, contentType, problem := Multipart([]Part{part}, 4096)
	if problem != nil || !bytes.Contains(body, []byte("original")) || !strings.HasPrefix(contentType, "multipart/form-data; boundary=") {
		t.Fatal("multipart snapshot failed")
	}
	if _, _, problem := Multipart([]Part{NewPart("field", "bad\r\nname", "", nil, true)}, 4096); FailureKind(problem) != 1 {
		t.Fatal("header injection accepted")
	}
}
