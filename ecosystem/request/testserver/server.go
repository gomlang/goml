package testserver

import (
	"compress/gzip"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"encoding/pem"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Server interface{ server() }
type state struct {
	http      *httptest.Server
	cert, key string
	entered   chan struct{}
	released  chan struct{}
	release   sync.Once
}

func (*state) server() {}
func Start(secure, mutual bool) Server {
	entered := make(chan struct{}, 1024)
	released := make(chan struct{})
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case entered <- struct{}{}:
		default:
		}
		w.Header().Set("X-Remote", r.RemoteAddr)
		switch r.URL.Path {
		case "/held":
			select {
			case <-r.Context().Done():
				return
			case <-released:
				io.WriteString(w, "released")
			}
		case "/close":
			w.Header().Set("Connection", "close")
			io.WriteString(w, "closed-response")
		case "/echo":
			body, _ := io.ReadAll(r.Body)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"method": r.Method, "path": r.URL.Path, "query": r.URL.RawQuery, "body": string(body), "authorization": r.Header.Get("Authorization"), "proxy_authorization": r.Header.Get("Proxy-Authorization"), "cookie": r.Header.Get("Cookie"), "custom": strings.Join(r.Header.Values("X-Custom"), "|"), "content_type": r.Header.Get("Content-Type"), "user_agent": r.Header.Get("User-Agent")})
		case "/multipart":
			if err := r.ParseMultipartForm(1 << 20); err != nil {
				w.WriteHeader(400)
				return
			}
			defer r.MultipartForm.RemoveAll()
			file, header, err := r.FormFile("upload")
			if err != nil {
				w.WriteHeader(400)
				return
			}
			defer file.Close()
			content, _ := io.ReadAll(file)
			json.NewEncoder(w).Encode(map[string]string{"name": r.FormValue("name"), "filename": header.Filename, "mime": header.Header.Get("Content-Type"), "body": string(content)})
		case "/json":
			w.Header().Set("Content-Type", "application/json")
			io.Copy(w, r.Body)
		case "/headers":
			w.Header().Add("X-Many", "first")
			w.Header().Add("X-Many", "second")
			io.WriteString(w, "headers")
		case "/redirect":
			code, _ := strconv.Atoi(r.URL.Query().Get("code"))
			if code < 300 || code > 399 {
				code = 302
			}
			w.Header().Set("Location", r.URL.Query().Get("to"))
			w.WriteHeader(code)
		case "/loop":
			w.Header().Set("Location", "/loop")
			w.WriteHeader(302)
		case "/gzip", "/gzip-small", "/gzip-empty":
			w.Header().Set("Content-Encoding", "gzip")
			z := gzip.NewWriter(w)
			if r.URL.Path == "/gzip" {
				io.WriteString(z, strings.Repeat("compressed-", 1024))
			} else if r.URL.Path == "/gzip-small" {
				io.WriteString(z, "ok")
			}
			z.Close()
		case "/chunked":
			w.WriteHeader(200)
			w.(http.Flusher).Flush()
			io.WriteString(w, "chunk-one")
			w.(http.Flusher).Flush()
			io.WriteString(w, "chunk-two")
		case "/large":
			io.WriteString(w, strings.Repeat("x", 4096))
		case "/status":
			w.WriteHeader(422)
			io.WriteString(w, "unprocessable")
		case "/invalid":
			w.Write([]byte{255, 254, 0})
		case "/cookie/set":
			http.SetCookie(w, &http.Cookie{Name: "session", Value: "goml", Path: "/", HttpOnly: true})
			http.SetCookie(w, &http.Cookie{Name: "domain", Value: "rejected", Path: "/", Domain: r.URL.Query().Get("domain")})
		case "/cookie/get":
			io.WriteString(w, r.Header.Get("Cookie"))
		case "/slow":
			select {
			case <-r.Context().Done():
				return
			case <-time.After(2 * time.Second):
				io.WriteString(w, "late")
			}
		case "/slowbody":
			w.WriteHeader(200)
			w.(http.Flusher).Flush()
			select {
			case <-r.Context().Done():
				return
			case <-time.After(2 * time.Second):
				io.WriteString(w, "late")
			}
		default:
			io.WriteString(w, "ok")
		}
	})
	server := httptest.NewUnstartedServer(handler)
	server.Config.ErrorLog = log.New(io.Discard, "", 0)
	result := &state{http: server, entered: entered, released: released}
	if secure {
		public, private, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			panic(err)
		}
		template := &x509.Certificate{SerialNumber: big.NewInt(1), Subject: pkix.Name{CommonName: "localhost"}, NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(time.Hour), DNSNames: []string{"localhost"}, IPAddresses: []net.IP{net.ParseIP("127.0.0.1")}, KeyUsage: x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign, ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth}, IsCA: true, BasicConstraintsValid: true}
		der, err := x509.CreateCertificate(rand.Reader, template, template, public, private)
		if err != nil {
			panic(err)
		}
		key, err := x509.MarshalPKCS8PrivateKey(private)
		if err != nil {
			panic(err)
		}
		result.cert = string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}))
		result.key = string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: key}))
		pair, err := tls.X509KeyPair([]byte(result.cert), []byte(result.key))
		if err != nil {
			panic(err)
		}
		config := &tls.Config{Certificates: []tls.Certificate{pair}, MinVersion: tls.VersionTLS12}
		if mutual {
			roots := x509.NewCertPool()
			roots.AppendCertsFromPEM([]byte(result.cert))
			config.ClientCAs = roots
			config.ClientAuth = tls.RequireAndVerifyClientCert
		}
		server.TLS = config
		server.EnableHTTP2 = true
		server.StartTLS()
	} else {
		server.Start()
	}
	return result
}
func URL(value Server) string         { return value.(*state).http.URL }
func Certificate(value Server) string { return value.(*state).cert }
func Key(value Server) string         { return value.(*state).key }
func Close(value Server)              { value.(*state).http.Close() }

func Release(value Server) {
	state := value.(*state)
	state.release.Do(func() { close(state.released) })
}

func Wait(value Server) bool {
	select {
	case <-value.(*state).entered:
		return true
	case <-time.After(3 * time.Second):
		return false
	}
}

func StartProxy(secure bool) Server {
	transport := &http.Transport{DisableKeepAlives: true}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "CONNECT" {
			remote, err := net.DialTimeout("tcp", r.Host, 3*time.Second)
			if err != nil {
				http.Error(w, err.Error(), 502)
				return
			}
			defer remote.Close()
			local, _, err := w.(http.Hijacker).Hijack()
			if err != nil {
				return
			}
			defer local.Close()
			io.WriteString(local, "HTTP/1.1 200 Connection Established\r\n\r\n")
			done := make(chan struct{})
			go func() {
				io.Copy(remote, local)
				remote.Close()
				close(done)
			}()
			io.Copy(local, remote)
			local.Close()
			<-done
			return
		}
		request := r.Clone(r.Context())
		request.RequestURI = ""
		response, err := transport.RoundTrip(request)
		if err != nil {
			http.Error(w, err.Error(), 502)
			return
		}
		defer response.Body.Close()
		for name, values := range response.Header {
			for _, value := range values {
				w.Header().Add(name, value)
			}
		}
		w.WriteHeader(response.StatusCode)
		io.Copy(w, response.Body)
	})
	server := httptest.NewUnstartedServer(handler)
	result := &state{http: server, entered: make(chan struct{}, 1)}
	if secure {
		server.StartTLS()
		result.cert = string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: server.Certificate().Raw}))
	} else {
		server.Start()
	}
	return result
}
