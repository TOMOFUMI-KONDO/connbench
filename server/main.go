package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"

	"github.com/TOMOFUMI-KONDO/connbench"
)

const (
	addr = ":44300"
)

func main() {
	cfg, err := genTLSCfg()
	if err != nil {
		panic(err)
	}

	listener, err := tls.Listen("tcp", addr, cfg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Listening %s\n", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Println("connection accepted.")

	_, err := conn.Write([]byte("hello"))
	if err != nil {
		fmt.Println(err)
	}
}

func genTLSCfg() (*tls.Config, error) {
	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, fmt.Errorf("failed to rsa.GenerateKey(); %w", err)
	}

	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		return nil, fmt.Errorf("failed to x509.CreateCertificate(); %w", err)
	}

	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return nil, fmt.Errorf("failed to tls.X509KeyPair(); %w", err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{connbench.NextProto},
	}, nil
}
