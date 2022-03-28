package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/TOMOFUMI-KONDO/connbench/client"
	"github.com/lucas-clemente/quic-go"
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

		sess, err := quic.DialAddr(addr, client.GenTLSCfg(), nil)
		if err != nil {
			panic(err)
		}

		stream, err := sess.OpenStreamSync(context.Background())
		if err != nil {
			panic(err)
		}

		if err = client.HandleConn(stream, startAt, durations, &idx); err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}
}
