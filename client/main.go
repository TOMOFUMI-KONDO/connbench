package main

import (
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/TOMOFUMI-KONDO/connbench"
)

const (
	addr  = "localhost:44300"
	times = 1000
)

func main() {
	var sum int64 = 0
	for i := 0; i < times; i++ {
		startAt := time.Now()

		conn, err := tls.Dial("tcp", addr, genTLSCfg())
		if err != nil {
			panic(err)
		}

		acceptedAt, err := handleConn(conn)
		if err != nil {
			panic(err)
		}

		duration := acceptedAt.Sub(startAt)
		fmt.Println(duration)
		sum += duration.Microseconds()
	}
	fmt.Printf("Average: %d[μs]\n", sum/times)
}

func handleConn(conn *tls.Conn) (*time.Time, error) {
	defer conn.Close()

	buf := make([]byte, binary.MaxVarintLen64)
	if _, err := conn.Read(buf); err != nil {
		return nil, err
	}

	acceptedAtUnix, _ := binary.Varint(buf)
	acceptedAt := time.Unix(0, acceptedAtUnix)

	return &acceptedAt, nil
}

func genTLSCfg() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{connbench.NextProto},
	}
}
