package native

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io"
	"math/big"
	"net"
	"sync"
	"time"
)

var listener net.Listener
var certificatePEM []byte
var keyPEM []byte
var connections = map[net.Conn]bool{}
var mutex sync.Mutex
var workers sync.WaitGroup
var stopped chan struct{}

func Start(requireClient, stall bool) uint16 {
	Stop()
	public, private, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	template := &x509.Certificate{
		SerialNumber: big.NewInt(1), Subject: pkix.Name{CommonName: "localhost"},
		DNSNames:  []string{"localhost"},
		NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true, IsCA: true,
	}
	der, err := x509.CreateCertificate(rand.Reader, template, template, public, private)
	if err != nil {
		panic(err)
	}
	certificatePEM = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})
	key, err := x509.MarshalPKCS8PrivateKey(private)
	if err != nil {
		panic(err)
	}
	keyPEM = pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: key})
	pair, err := tls.X509KeyPair(certificatePEM, keyPEM)
	if err != nil {
		panic(err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{pair}, NextProtos: []string{"goml-test"}, MinVersion: tls.VersionTLS12, MaxVersion: tls.VersionTLS12}
	if requireClient {
		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(certificatePEM)
		config.ClientAuth = tls.RequireAndVerifyClientCert
		config.ClientCAs = pool
	}
	listener, err = net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	stopped = make(chan struct{})
	current := listener
	stop := stopped
	workers.Add(1)
	go func() {
		defer workers.Done()
		for {
			raw, err := current.Accept()
			if err != nil {
				return
			}
			mutex.Lock()
			select {
			case <-stop:
				mutex.Unlock()
				raw.Close()
				return
			default:
			}
			connections[raw] = true
			mutex.Unlock()
			workers.Add(1)
			go func() {
				defer workers.Done()
				defer raw.Close()
				defer func() { mutex.Lock(); delete(connections, raw); mutex.Unlock() }()
				if stall {
					<-stop
					return
				}
				conn := tls.Server(raw, config)
				if conn.Handshake() != nil {
					return
				}
				_, _ = io.Copy(conn, conn)
			}()
		}
	}()
	return uint16(listener.Addr().(*net.TCPAddr).Port)
}

func Certificate() []byte { return append([]byte(nil), certificatePEM...) }
func PrivateKey() []byte  { return append([]byte(nil), keyPEM...) }

func Stop() {
	if listener == nil {
		return
	}
	listener.Close()
	close(stopped)
	mutex.Lock()
	for conn := range connections {
		conn.Close()
	}
	mutex.Unlock()
	workers.Wait()
	listener = nil
}
