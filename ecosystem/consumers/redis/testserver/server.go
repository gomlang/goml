package testserver

import (
	"bufio"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io"
	"math/big"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Server interface{ peer() }
type state struct {
	listener        net.Listener
	config          *tls.Config
	mode, cert, key string
	stop            chan struct{}
	stalled         chan struct{}
	done            chan struct{}
	once            sync.Once
	mutex           sync.Mutex
	connection      net.Conn
	count           int
	failure         string
}

func (*state) peer() {}
func Start(mode string) Server {
	public, private, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	template := &x509.Certificate{SerialNumber: big.NewInt(1), Subject: pkix.Name{CommonName: "localhost"}, DNSNames: []string{"localhost"}, IPAddresses: []net.IP{net.ParseIP("127.0.0.1")}, NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(time.Hour), IsCA: true, BasicConstraintsValid: true, KeyUsage: x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign, ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth}}
	der, err := x509.CreateCertificate(rand.Reader, template, template, public, private)
	if err != nil {
		panic(err)
	}
	key, err := x509.MarshalPKCS8PrivateKey(private)
	if err != nil {
		panic(err)
	}
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: key})
	pair, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{pair}, MinVersion: tls.VersionTLS12}
	if mode == "mutual" || mode == "pool_mutual" {
		roots := x509.NewCertPool()
		roots.AppendCertsFromPEM(certPEM)
		config.ClientCAs = roots
		config.ClientAuth = tls.RequireAndVerifyClientCert
	}
	peer := &state{listener: listener, config: config, mode: mode, cert: string(certPEM), key: string(keyPEM), stop: make(chan struct{}), stalled: make(chan struct{}, 1), done: make(chan struct{})}
	go peer.serve()
	return peer
}
func Port(server Server) int           { return server.(*state).listener.Addr().(*net.TCPAddr).Port }
func Certificate(server Server) string { return server.(*state).cert }
func Key(server Server) string         { return server.(*state).key }
func Wait(server Server) bool {
	select {
	case <-server.(*state).stalled:
		return true
	case <-time.After(8 * time.Second):
		return false
	}
}
func Count(server Server) int {
	s := server.(*state)
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.count
}
func Close(server Server) string {
	s := server.(*state)
	s.once.Do(func() {
		close(s.stop)
		s.listener.Close()
		s.mutex.Lock()
		if s.connection != nil {
			s.connection.Close()
		}
		s.mutex.Unlock()
	})
	select {
	case <-s.done:
	case <-time.After(10 * time.Second):
		return "TLS peer did not terminate"
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.failure
}
func (s *state) fail(err error) { s.mutex.Lock(); s.failure = err.Error(); s.mutex.Unlock() }
func command(reader *bufio.Reader) ([]string, error) {
	header, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	if !strings.HasPrefix(header, "*") || !strings.HasSuffix(header, "\r\n") {
		return nil, fmt.Errorf("invalid array header")
	}
	count, err := strconv.Atoi(strings.TrimSuffix(header[1:], "\r\n"))
	if err != nil || count < 1 || count > 32 {
		return nil, fmt.Errorf("invalid count")
	}
	result := make([]string, count)
	for index := range result {
		header, err = reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		if !strings.HasPrefix(header, "$") || !strings.HasSuffix(header, "\r\n") {
			return nil, fmt.Errorf("invalid bulk header")
		}
		length, err := strconv.Atoi(strings.TrimSuffix(header[1:], "\r\n"))
		if err != nil || length < 0 || length > 1024 {
			return nil, fmt.Errorf("invalid length")
		}
		data := make([]byte, length+2)
		if _, err = io.ReadFull(reader, data); err != nil {
			return nil, err
		}
		if string(data[length:]) != "\r\n" {
			return nil, fmt.Errorf("missing CRLF")
		}
		result[index] = string(data[:length])
	}
	return result, nil
}
func (s *state) serve() {
	defer close(s.done)
	raw, err := s.listener.Accept()
	if err != nil {
		select {
		case <-s.stop:
			return
		default:
			s.fail(err)
			return
		}
	}
	s.mutex.Lock()
	s.connection = raw
	s.mutex.Unlock()
	defer raw.Close()
	raw.SetDeadline(time.Now().Add(8 * time.Second))
	if s.mode == "handshake_timeout" {
		<-s.stop
		return
	}
	var conn net.Conn = raw
	if s.mode != "plain" && s.mode != "pool_plain" {
		secure := tls.Server(raw, s.config)
		if err := secure.Handshake(); err != nil {
			if s.mode != "untrusted" && s.mode != "name" {
				s.fail(err)
			}
			return
		}
		conn = secure
	}
	reader := bufio.NewReader(conn)
	for {
		args, err := command(reader)
		if err != nil {
			if err != io.EOF {
				select {
				case <-s.stop:
				default:
					s.fail(err)
				}
			}
			return
		}
		s.mutex.Lock()
		s.count++
		s.mutex.Unlock()
		if len(args) == 2 && args[0] == "HELLO" && args[1] == "3" {
			if s.mode == "hello_timeout" {
				s.stalled <- struct{}{}
				<-s.stop
				return
			}
			_, err = io.WriteString(conn, "%1\r\n+proto\r\n:3\r\n")
		} else if len(args) == 1 && args[0] == "PING" {
			_, err = io.WriteString(conn, "+PONG\r\n")
		} else if len(args) == 2 && args[0] == "ECHO" && args[1] == "stall" {
			_, err = io.WriteString(conn, "$10\r\nx")
			if err != nil {
				s.fail(err)
				return
			}
			s.stalled <- struct{}{}
			<-s.stop
			return
		} else if len(args) == 2 && args[0] == "ECHO" && args[1] == "secure\x00echo" {
			_, err = fmt.Fprintf(conn, "$%d\r\n%s\r\n", len(args[1]), args[1])
		} else {
			s.fail(fmt.Errorf("unexpected command %q", args))
			return
		}
		if err != nil {
			s.fail(err)
			return
		}
	}
}
