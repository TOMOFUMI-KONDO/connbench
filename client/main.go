package main

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/TOMOFUMI-KONDO/connbench"
)

const (
	addr  = "localhost:44300"
	times = 10
)

func main() {
	var sum int64 = 0
	for i := 0; i < times; i++ {
		start := time.Now()

		conn, err := tls.Dial("tcp", addr, genTLSCfg())
		if err != nil {
			panic(err)
		}
		handleConn(conn)

		duration := time.Since(start)

		fmt.Println(duration)
		sum += duration.Milliseconds()
	}
	fmt.Printf("Average[ms]: %d", sum/times)
}

func handleConn(conn *tls.Conn) {
	defer conn.Close()
}

func genTLSCfg() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{connbench.NextProto},
	}
}
