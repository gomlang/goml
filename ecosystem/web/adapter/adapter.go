package adapter

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Failure struct {
	Code    int
	Message string
}

func FailureCode(f Failure) int                   { return f.Code }
func FailureMessage(f Failure) string             { return f.Message }
func NewFailure(code int, message string) Failure { return Failure{code, message} }
func Success() Failure                            { return Failure{} }

func failed(err error) Failure {
	if err == nil {
		return Success()
	}
	code := 6
	var network net.Error
	if errors.As(err, &network) && network.Timeout() {
		code = 5
	}
	var limit *http.MaxBytesError
	if errors.As(err, &limit) {
		code = 3
	}
	if errors.Is(err, context.Canceled) {
		code = 4
	}
	if errors.Is(err, context.DeadlineExceeded) {
		code = 5
	}
	return Failure{code, err.Error()}
}

type request struct {
	req    *http.Request
	ctx    context.Context
	mu     sync.Mutex
	closed atomic.Bool
}
type Request interface{ requestHandle() }

func (*request) requestHandle() {}

type writer struct {
	out       http.ResponseWriter
	ctx       context.Context
	mu        sync.Mutex
	closed    bool
	committed bool
	count     int64
	limit     int64
	head      bool
}
type Writer interface{ writerHandle() }

func (*writer) writerHandle() {}

type Response struct {
	status  int
	headers http.Header
	body    []byte
	stream  func(Writer) Failure
}
type Capture struct {
	status  int
	headers http.Header
	body    []byte
}
type server struct {
	http          *http.Server
	address       string
	cancel        context.CancelFunc
	done          chan struct{}
	mu            sync.Mutex
	closing       bool
	shutdownDone  chan struct{}
	shutdownError Failure
	serveError    error
}
type Server interface{ serverHandle() }

func (*server) serverHandle() {}

func RequestMethod(r Request) string   { return r.(*request).req.Method }
func RequestPath(r Request) string     { return r.(*request).req.URL.EscapedPath() }
func RequestQuery(r Request) string    { return r.(*request).req.URL.RawQuery }
func RequestRemote(r Request) string   { return r.(*request).req.RemoteAddr }
func RequestHost(r Request) string     { return r.(*request).req.Host }
func RequestNames(r Request) []string  { n, _ := headerPairs(r.(*request).req.Header); return n }
func RequestValues(r Request) []string { _, v := headerPairs(r.(*request).req.Header); return v }

func headerPairs(h http.Header) ([]string, []string) {
	keys := make([]string, 0, len(h))
	for key := range h {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var names, values []string
	for _, key := range keys {
		for _, value := range h[key] {
			names = append(names, key)
			values = append(values, value)
		}
	}
	return names, values
}

func contextFailure(ctx context.Context) Failure {
	if deadline, ok := ctx.Deadline(); ok && !time.Now().Before(deadline) {
		return Failure{5, "request deadline exceeded"}
	}
	return failed(ctx.Err())
}

func Check(r Request) Failure {
	if r.(*request).closed.Load() {
		return Failure{7, "request has finished"}
	}
	return contextFailure(r.(*request).ctx)
}

func Wait(r Request, nanos int64) Failure {
	if err := Check(r); err.Code != 0 {
		return err
	}
	if nanos < 0 {
		return Failure{1, "negative wait duration"}
	}
	timer := time.NewTimer(time.Duration(nanos))
	defer timer.Stop()
	select {
	case <-r.(*request).ctx.Done():
		return contextFailure(r.(*request).ctx)
	case <-timer.C:
		return Check(r)
	}
}

func Read(r Request, size int) ([]byte, Failure) {
	if size < 1 || size > 65536 {
		return nil, Failure{1, "read size must be in 1..65536"}
	}
	r.(*request).mu.Lock()
	defer r.(*request).mu.Unlock()
	if err := Check(r); err.Code != 0 {
		return nil, err
	}
	data := make([]byte, size)
	n, err := r.(*request).req.Body.Read(data)
	if check := r.(*request).ctx.Err(); check != nil {
		return nil, contextFailure(r.(*request).ctx)
	}
	if err == io.EOF {
		err = nil
	}
	return data[:n], failed(err)
}

func CloseBody(r Request) Failure { return failed(r.(*request).req.Body.Close()) }

func NewResponse(status int, names, values []string, body []byte, stream func(Writer) Failure) Response {
	h := make(http.Header)
	for i, name := range names {
		if i < len(values) {
			h.Add(name, values[i])
		}
	}
	return Response{status, h, bytes.Clone(body), stream}
}

func Write(w Writer, data []byte, flush bool) Failure {
	if len(data) > 65536 {
		return Failure{1, "write chunk exceeds 65536 bytes"}
	}
	w.(*writer).mu.Lock()
	defer w.(*writer).mu.Unlock()
	if w.(*writer).closed {
		return Failure{7, "response stream has finished"}
	}
	if err := w.(*writer).ctx.Err(); err != nil {
		return failed(err)
	}
	if int64(len(data)) > w.(*writer).limit-w.(*writer).count {
		return Failure{3, "response exceeds capture limit"}
	}
	w.(*writer).count += int64(len(data))
	if !w.(*writer).head {
		n, err := w.(*writer).out.Write(data)
		if err != nil {
			return failed(err)
		}
		if n != len(data) {
			return failed(io.ErrShortWrite)
		}
		if flush {
			if err := http.NewResponseController(w.(*writer).out).Flush(); err != nil {
				return failed(err)
			}
		}
	}
	return Success()
}

func WriterCheck(w Writer) Failure {
	w.(*writer).mu.Lock()
	defer w.(*writer).mu.Unlock()
	if w.(*writer).closed {
		return Failure{7, "response stream has finished"}
	}
	return contextFailure(w.(*writer).ctx)
}

func validHeader(name, value string) bool {
	if name == "" {
		return false
	}
	for _, c := range []byte(name) {
		if !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || strings.ContainsRune("!#$%&'*+-.^_`|~", rune(c))) {
			return false
		}
	}
	for _, c := range []byte(value) {
		if c == 127 || c < 32 && c != 9 {
			return false
		}
	}
	return true
}

func transportHeader(name string) bool {
	switch strings.ToLower(name) {
	case "content-length", "transfer-encoding", "connection", "trailer", "upgrade", "proxy-connection":
		return true
	}
	return false
}

func emit(w *writer, response Response) Failure {
	if response.status < 200 || response.status > 599 {
		return Failure{1, "response status must be in 200..599"}
	}
	for name, values := range response.headers {
		for _, value := range values {
			if !validHeader(name, value) || transportHeader(name) {
				return Failure{1, "invalid response header"}
			}
		}
	}
	if (response.status == 204 || response.status == 304) && (len(response.body) != 0 || response.stream != nil) {
		return Failure{1, "status does not allow a response body"}
	}
	if err := w.ctx.Err(); err != nil {
		return contextFailure(w.ctx)
	}
	for name, values := range response.headers {
		w.out.Header()[name] = append([]string(nil), values...)
	}
	w.out.WriteHeader(response.status)
	w.committed = true
	if w.head {
		return Success()
	}
	if response.stream != nil {
		if err := http.NewResponseController(w.out).Flush(); err != nil {
			return failed(err)
		}
		return response.stream(w)
	}
	if int64(len(response.body)) > w.limit {
		return Failure{3, "response exceeds capture limit"}
	}
	for len(response.body) > 0 {
		n := min(len(response.body), 65536)
		if failure := Write(w, response.body[:n], false); failure.Code != 0 {
			return failure
		}
		response.body = response.body[n:]
	}
	return Success()
}

func execute(out http.ResponseWriter, req *http.Request, handler func(Request) Response, timeout time.Duration, bodyLimit, captureLimit int64) (failure Failure) {
	ctx, cancel := context.WithTimeout(req.Context(), timeout)
	defer cancel()
	req = req.WithContext(ctx)
	req.Body = http.MaxBytesReader(out, req.Body, bodyLimit)
	current := &request{req: req, ctx: ctx}
	writer := &writer{out: out, ctx: ctx, limit: captureLimit, head: req.Method == "HEAD"}
	controller := http.NewResponseController(out)
	_ = controller.EnableFullDuplex()
	deadline, _ := ctx.Deadline()
	_ = controller.SetReadDeadline(deadline)
	_ = controller.SetWriteDeadline(deadline)
	cancelled := make(chan struct{})
	stop := context.AfterFunc(ctx, func() {
		defer close(cancelled)
		_ = controller.SetReadDeadline(time.Now())
		_ = controller.SetWriteDeadline(time.Now())
		_ = req.Body.Close()
	})
	defer func() {
		if !stop() {
			<-cancelled
		}
		current.closed.Store(true)
		_ = req.Body.Close()
		writer.mu.Lock()
		writer.closed = true
		writer.mu.Unlock()
		if recovered := recover(); recovered != nil {
			failure = Failure{8, fmt.Sprintf("handler panicked: %v", recovered)}
		}
		if failure.Code != 0 {
			if !writer.committed {
				_ = controller.SetWriteDeadline(time.Now().Add(100 * time.Millisecond))
				status := 500
				if failure.Code == 3 {
					status = 413
				}
				if failure.Code == 5 {
					status = 504
				}
				if failure.Code == 4 {
					status = 499
				}
				out.Header().Set("Content-Type", "text/plain; charset=utf-8")
				if failure.Code == 4 || failure.Code == 5 {
					out.Header().Set("Connection", "close")
				}
				out.WriteHeader(status)
				_, _ = io.WriteString(out, http.StatusText(status))
			} else if _, memory := out.(*httptest.ResponseRecorder); !memory {
				panic(http.ErrAbortHandler)
			}
		}
		_ = controller.SetReadDeadline(time.Time{})
		_ = controller.SetWriteDeadline(time.Time{})
	}()
	if req.ContentLength > bodyLimit {
		return Failure{3, "request body exceeds limit"}
	}
	failure = emit(writer, handler(current))
	return failure
}

func Dispatch(handler func(Request) Response, method, target string, names, values []string, body []byte, timeout int64, bodyLimit, captureLimit int64) (Capture, Failure) {
	if handler == nil || timeout <= 0 || bodyLimit < 0 || captureLimit < 0 {
		return Capture{}, Failure{1, "invalid dispatcher options"}
	}
	if !strings.HasPrefix(target, "/") || strings.HasPrefix(target, "//") || strings.Contains(target, "#") {
		return Capture{}, Failure{1, "target must be an origin-form path"}
	}
	req, err := http.NewRequest(method, "http://dispatch.invalid"+target, bytes.NewReader(body))
	if err != nil {
		return Capture{}, failed(err)
	}
	for i, name := range names {
		if i >= len(values) || !validHeader(name, values[i]) {
			return Capture{}, Failure{1, "invalid request header"}
		}
		req.Header.Add(name, values[i])
	}
	recorder := httptest.NewRecorder()
	failure := execute(recorder, req, handler, time.Duration(timeout), bodyLimit, captureLimit)
	result := recorder.Result()
	if failure.Code != 0 {
		return Capture{}, failure
	}
	return Capture{result.StatusCode, result.Header.Clone(), bytes.Clone(recorder.Body.Bytes())}, Success()
}

func CaptureStatus(c Capture) int      { return c.status }
func CaptureNames(c Capture) []string  { n, _ := headerPairs(c.headers); return n }
func CaptureValues(c Capture) []string { _, v := headerPairs(c.headers); return v }
func CaptureBody(c Capture) []byte     { return bytes.Clone(c.body) }

func Start(address string, handler func(Request) Response, timeout int64, bodyLimit int64, maxHeaders, maxConcurrent int, headerTimeout int64) (Server, Failure) {
	if handler == nil || timeout <= 0 || bodyLimit < 0 || maxHeaders < 1024 || maxConcurrent < 1 || headerTimeout <= 0 {
		return nil, Failure{1, "invalid server options"}
	}
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, failed(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	value := &server{address: listener.Addr().String(), cancel: cancel, done: make(chan struct{}), shutdownDone: make(chan struct{})}
	slots := make(chan struct{}, maxConcurrent)
	value.http = &http.Server{
		ReadHeaderTimeout: time.Duration(headerTimeout),
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    maxHeaders,
		BaseContext:       func(net.Listener) context.Context { return ctx },
		Handler: http.HandlerFunc(func(out http.ResponseWriter, req *http.Request) {
			select {
			case slots <- struct{}{}:
				defer func() { <-slots }()
			default:
				http.Error(out, "server concurrency limit reached", http.StatusServiceUnavailable)
				return
			}
			_ = execute(out, req, handler, time.Duration(timeout), bodyLimit, 1<<63-1)
		}),
	}
	go func() {
		err := value.http.Serve(listener)
		value.mu.Lock()
		if !errors.Is(err, http.ErrServerClosed) {
			value.serveError = err
		}
		value.mu.Unlock()
		close(value.done)
	}()
	return value, Success()
}

func Address(s Server) string {
	if s == nil {
		return ""
	}
	return s.(*server).address
}
func Shutdown(s Server, nanos int64) Failure {
	if s == nil || nanos <= 0 {
		return Failure{1, "shutdown requires a live server and positive timeout"}
	}
	value := s.(*server)
	value.mu.Lock()
	if value.closing {
		value.mu.Unlock()
		timer := time.NewTimer(time.Duration(nanos))
		defer timer.Stop()
		select {
		case <-value.shutdownDone:
			return value.shutdownError
		case <-timer.C:
			return Failure{5, "timed out waiting for concurrent shutdown"}
		}
	}
	value.closing = true
	value.mu.Unlock()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(nanos))
	err := value.http.Shutdown(ctx)
	cancel()
	if err != nil {
		value.cancel()
		_ = value.http.Close()
	}
	value.cancel()
	<-value.done
	value.mu.Lock()
	if err == nil {
		err = value.serveError
	}
	value.shutdownError = failed(err)
	value.mu.Unlock()
	close(value.shutdownDone)
	return value.shutdownError
}

func IsClosed(s Server) bool {
	if s == nil {
		return true
	}
	s.(*server).mu.Lock()
	defer s.(*server).mu.Unlock()
	return s.(*server).closing
}
