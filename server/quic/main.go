package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"time"

	"github.com/TOMOFUMI-KONDO/connbench/server"
	"github.com/lucas-clemente/quic-go"
)

const (
	addr = ":44300"
)

var (
	times int64 = 1
	sum   int64 = 0
)

func main() {
	listener, err := quic.ListenAddr(addr, server.GenTLSCfg(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Listening %s\n", addr)

	for {
		sess, err := listener.Accept(context.Background())
		if err != nil {
			panic(err)
		}

		go handleSess(sess)
	}
}

func handleSess(sess quic.Session) {
	stream, err := sess.AcceptStream(context.Background())
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	acceptedAt := time.Now()

	buf := make([]byte, binary.MaxVarintLen64)
	if _, err := stream.Read(buf); err != io.EOF && err != nil {
		panic(err)
	}

	startAtUnix, _ := binary.Varint(buf)
	startAt := time.Unix(0, startAtUnix)

	duration := acceptedAt.Sub(startAt)
	sum += duration.Microseconds()

	fmt.Printf("Duration: %d[μs] (%dth)\n", sum/times, times)

	times++
}
