package adapter

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)

type Failure interface{ failure() }
type failureState struct {
	kind    int
	message string
}

func (*failureState) failure()              {}
func fail(kind int, message string) Failure { return &failureState{kind, message} }
func FailureKind(value Failure) int {
	if value == nil {
		return 0
	}
	return value.(*failureState).kind
}
func FailureMessage(value Failure) string {
	if value == nil {
		return ""
	}
	return value.(*failureState).message
}
func failure(err error) Failure {
	if err == nil {
		return nil
	}
	kind := 4
	var network net.Error
	if errors.Is(err, context.DeadlineExceeded) || errors.As(err, &network) && network.Timeout() {
		kind = 2
	} else if errors.Is(err, context.Canceled) {
		kind = 3
	}
	return fail(kind, err.Error())
}

type Control interface{ control() }
type controlState struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func (*controlState) control() {}
func NewControl(nanoseconds int64) Control {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(nanoseconds))
	return &controlState{ctx, cancel}
}
func Cancel(value Control) { value.(*controlState).cancel() }

type Client interface{ client() }
type clientState struct {
	http      *http.Client
	transport *http.Transport
	ctx       context.Context
	cancel    context.CancelFunc
	mu        sync.Mutex
	closed    bool
}

func (*clientState) client() {}

type hostJar struct{ jar *cookiejar.Jar }

func (j hostJar) Cookies(u *url.URL) []*http.Cookie { return j.jar.Cookies(u) }
func (j hostJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	accepted := make([]*http.Cookie, 0, len(cookies))
	for _, cookie := range cookies {
		if cookie.Domain == "" {
			accepted = append(accepted, cookie)
		}
	}
	j.jar.SetCookies(u, accepted)
}
func NewClient(connectMS, idleMS int64, maxIdle int, maxHeaders int64, gzip, cookies bool, ca, certificate, key, proxy string, environmentProxy bool) (Client, Failure) {
	if connectMS <= 0 || connectMS > 9223372036854 || idleMS <= 0 || idleMS > 9223372036854 || maxIdle < 1 || maxHeaders < 1 {
		return nil, fail(1, "invalid transport limits")
	}
	config := &tls.Config{MinVersion: tls.VersionTLS12}
	if ca != "" {
		roots, err := x509.SystemCertPool()
		if err != nil {
			roots = x509.NewCertPool()
		}
		if !roots.AppendCertsFromPEM([]byte(ca)) {
			return nil, fail(1, "invalid root certificate PEM")
		}
		config.RootCAs = roots
	}
	if certificate != "" || key != "" {
		pair, err := tls.X509KeyPair([]byte(certificate), []byte(key))
		if err != nil {
			return nil, fail(1, "invalid client identity PEM")
		}
		config.Certificates = []tls.Certificate{pair}
	}
	transport := &http.Transport{
		DialContext:     (&net.Dialer{Timeout: time.Duration(connectMS) * time.Millisecond, KeepAlive: 30 * time.Second}).DialContext,
		TLSClientConfig: config, TLSHandshakeTimeout: time.Duration(connectMS) * time.Millisecond,
		ForceAttemptHTTP2: true, DisableCompression: !gzip, MaxIdleConns: maxIdle,
		MaxIdleConnsPerHost: maxIdle, IdleConnTimeout: time.Duration(idleMS) * time.Millisecond,
		MaxResponseHeaderBytes: maxHeaders, ExpectContinueTimeout: time.Second,
	}
	if environmentProxy {
		transport.Proxy = http.ProxyFromEnvironment
	}
	if proxy != "" {
		parsed, err := url.Parse(proxy)
		if err != nil || parsed.Host == "" || parsed.Scheme != "http" && parsed.Scheme != "https" {
			return nil, fail(1, "proxy must be an absolute HTTP(S) URL")
		}
		transport.Proxy = http.ProxyURL(parsed)
	}
	client := &http.Client{Transport: transport, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	if cookies {
		jar, _ := cookiejar.New(nil)
		client.Jar = hostJar{jar}
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &clientState{http: client, transport: transport, ctx: ctx, cancel: cancel}, nil
}
func Close(value Client) {
	c := value.(*clientState)
	c.mu.Lock()
	if !c.closed {
		c.closed = true
		c.cancel()
	}
	c.mu.Unlock()
	c.transport.CloseIdleConnections()
}
func CloseIdle(value Client) { value.(*clientState).transport.CloseIdleConnections() }
func IsClosed(value Client) bool {
	c := value.(*clientState)
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.closed
}

func ParseURL(text, base string) (string, string, bool, Failure) {
	u, err := url.Parse(text)
	if err != nil {
		return "", "", false, fail(1, "invalid URL")
	}
	if base != "" {
		parent, err := url.Parse(base)
		if err != nil {
			return "", "", false, fail(1, "invalid base URL")
		}
		u = parent.ResolveReference(u)
	}
	u.Scheme = strings.ToLower(u.Scheme)
	if (u.Scheme != "http" && u.Scheme != "https") || u.Hostname() == "" || u.Opaque != "" || u.User != nil {
		return "", "", false, fail(1, "expected absolute HTTP(S) URL without credentials")
	}
	port := u.Port()
	if port == "" {
		if u.Scheme == "https" {
			port = "443"
		} else {
			port = "80"
		}
	}
	host := strings.ToLower(u.Hostname())
	u.Host = strings.ToLower(u.Host)
	u.Fragment = ""
	return u.String(), u.Scheme + "://" + net.JoinHostPort(host, port), u.Scheme == "https", nil
}
func Query(text string, names, values []string) (string, Failure) {
	u, err := url.Parse(text)
	if err != nil || len(names) != len(values) {
		return "", fail(1, "invalid query")
	}
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return "", fail(1, "invalid URL query")
	}
	for i, name := range names {
		q.Add(name, values[i])
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}
func EncodeForm(names, values []string) string {
	q := url.Values{}
	for i, name := range names {
		q.Add(name, values[i])
	}
	return q.Encode()
}

type Response interface{ response() }
type responseState struct {
	status        int
	version       string
	names, values []string
	body          []byte
}

func (*responseState) response()             {}
func ResponseStatus(value Response) int      { return value.(*responseState).status }
func ResponseVersion(value Response) string  { return value.(*responseState).version }
func ResponseNames(value Response) []string  { return value.(*responseState).names }
func ResponseValues(value Response) []string { return value.(*responseState).values }
func ResponseBody(value Response) []byte     { return value.(*responseState).body }
func Do(client Client, control Control, method, address string, names, values []string, body []byte, limit int64) (Response, Failure) {
	c, ctl := client.(*clientState), control.(*controlState)
	c.mu.Lock()
	closed := c.closed
	c.mu.Unlock()
	if closed {
		return nil, fail(7, "client is closed")
	}
	ctx, cancel := context.WithCancel(ctl.ctx)
	stop := context.AfterFunc(c.ctx, cancel)
	defer stop()
	defer cancel()
	if c.ctx.Err() != nil {
		cancel()
	}
	req, err := http.NewRequestWithContext(ctx, method, address, bytes.NewReader(body))
	if err != nil {
		return nil, fail(1, "invalid request")
	}
	for i, name := range names {
		req.Header.Add(name, values[i])
	}
	response, err := c.http.Do(req)
	if err != nil {
		return nil, failure(err)
	}
	defer response.Body.Close()
	if method != "HEAD" && response.ContentLength > limit {
		return nil, fail(5, "response body exceeds limit")
	}
	payload, err := io.ReadAll(io.LimitReader(response.Body, limit+1))
	if err != nil {
		return nil, failure(err)
	}
	if int64(len(payload)) > limit {
		return nil, fail(5, "response body exceeds limit")
	}
	result := &responseState{status: response.StatusCode, version: response.Proto, body: payload}
	keys := make([]string, 0, len(response.Header))
	for name := range response.Header {
		keys = append(keys, name)
	}
	sort.Strings(keys)
	for _, name := range keys {
		for _, value := range response.Header[name] {
			result.names = append(result.names, strings.ToLower(name))
			result.values = append(result.values, value)
		}
	}
	return result, nil
}
