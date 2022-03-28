package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"time"

	"github.com/TOMOFUMI-KONDO/connbench/client"
	"github.com/lucas-clemente/quic-go"
)

var (
	addr  string
	times int
)

func init() {
	flag.StringVar(&addr, "addr", "localhost:44300", "server address")
	flag.IntVar(&times, "times", 100, "number of times to try connection establishment")
	flag.Parse()
}

func main() {
	for i := 0; i < times; i++ {
		startAt := time.Now()
		fmt.Printf("StartAt: %s\n", startAt)

		sess, err := quic.DialAddr(addr, client.GenTLSCfg(), nil)
		if err != nil {
			panic(err)
		}

		err = handleSess(sess, startAt)
		if err != nil {
			panic(err)
		}
	}
}

func handleSess(sess quic.Session, startAt time.Time) error {
	stream, err := sess.OpenStreamSync(context.Background())
	if err != nil {
		return err
	}
	defer stream.Close()

	if _, err := stream.Write(client.Int64ToBytes(startAt.UnixNano())); err != nil {
		return err
	}

	_, err = io.ReadAll(stream)
	if err != nil {
		return err
	}

	return nil
}
