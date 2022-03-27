package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/TOMOFUMI-KONDO/connbench/client"
	"github.com/lucas-clemente/quic-go"
)

const (
	addr  = "localhost:44300"
	times = 100
)

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
