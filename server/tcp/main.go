package main

import (
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"github.com/TOMOFUMI-KONDO/connbench/server"
)

const (
	addr = ":44300"
)

var (
	times int64 = 1
	sum   int64 = 0
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

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	acceptedAt := time.Now()

	buf := make([]byte, binary.MaxVarintLen64)
	if _, err := conn.Read(buf); err != nil {
		panic(err)
	}

	startAtUnix, _ := binary.Varint(buf)
	startAt := time.Unix(0, startAtUnix)

	duration := acceptedAt.Sub(startAt)
	sum += duration.Microseconds()

	fmt.Printf("Duration: %d[μs] (%dth)\n", sum/times, times)

	times++
}
