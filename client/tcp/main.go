package main

import (
	"crypto/tls"
	"flag"
	"time"

	"github.com/TOMOFUMI-KONDO/connbench/client"
)

var (
	addr  string
	times int

	durations []time.Duration
	idx       = 0
)

func init() {
	flag.StringVar(&addr, "addr", "localhost:44300", "server address")
	flag.IntVar(&times, "times", 100, "number of times to try connection establishment")
	flag.Parse()
}

func main() {
	durations = make([]time.Duration, times)

	for i := 0; i < times; i++ {
		startAt := time.Now()

		conn, err := tls.Dial("tcp", addr, client.GenTLSCfg())
		if err != nil {
			panic(err)
		}

		if err = client.HandleConn(conn, startAt, durations, &idx); err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}

	if err := client.Render(durations, "tcp"); err != nil {
		panic(err)
	}
}
