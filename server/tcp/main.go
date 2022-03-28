package main

import (
	"crypto/tls"
	"fmt"

	"github.com/TOMOFUMI-KONDO/connbench/server"
)

const (
	addr = ":44300"
)

func main() {
	listener, err := tls.Listen("tcp", addr, server.GenTLSCfg())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Listening %s\n", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go server.HandleConn(conn)
	}
}
