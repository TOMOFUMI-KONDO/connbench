package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"time"

	"github.com/TOMOFUMI-KONDO/connbench"
)

const (
	addr = ":44300"
)

func main() {
	listener, err := tls.Listen("tcp", addr, genTLSCfg())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Listening %s\n", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	acceptedAt := time.Now()
	fmt.Printf("AcceptedAt: %s\n", acceptedAt)

	if _, err := conn.Write(int64ToBytes(acceptedAt.UnixNano())); err != nil {
		panic(err)
	}
}

func int64ToBytes(n int64) []byte {
	bytes := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(bytes, n)
	return bytes
}

func genTLSCfg() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}

	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}

	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{connbench.NextProto},
	}
}
