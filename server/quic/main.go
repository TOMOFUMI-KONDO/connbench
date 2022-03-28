package main

import (
	"context"
	"fmt"

	"github.com/TOMOFUMI-KONDO/connbench/server"
	"github.com/lucas-clemente/quic-go"
)

const (
	addr = ":44300"
)

func main() {
	listener, err := quic.ListenAddr(addr, server.GenTLSCfg(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Listening %s\n", addr)

	for {
		sess, err := listener.Accept(context.Background()) // here is time-consuming
		if err != nil {
			panic(err)
		}

		stream, err := sess.AcceptStream(context.Background())
		if err != nil {
			panic(err)
		}

		go server.HandleConn(stream)
	}
}
