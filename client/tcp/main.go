package main

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/TOMOFUMI-KONDO/connbench/client"
)

const (
	addr  = "localhost:44300"
	times = 100
)

func main() {
	for i := 0; i < times; i++ {
		startAt := time.Now()
		fmt.Printf("StartAt: %s\n", startAt)

		conn, err := tls.Dial("tcp", addr, client.GenTLSCfg())
		if err != nil {
			panic(err)
		}

		err = handleConn(conn, startAt)
		if err != nil {
			panic(err)
		}
	}
}

func handleConn(conn *tls.Conn, startAt time.Time) error {
	defer conn.Close()

	if _, err := conn.Write(client.Int64ToBytes(startAt.UnixNano())); err != nil {
		return err
	}

	return nil
}
