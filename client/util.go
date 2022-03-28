package client

import (
	"crypto/tls"
	"encoding/binary"
	"io"
	"time"

	"github.com/TOMOFUMI-KONDO/connbench"
)

func GenTLSCfg() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{connbench.NextProto},
	}
}

func HandleConn(conn io.ReadWriteCloser, startAt time.Time, durations []time.Duration, idx *int) error {
	defer conn.Close()

	if _, err := conn.Write([]byte("GET\n")); err != nil {
		return err
	}

	buf := make([]byte, binary.MaxVarintLen64)
	if _, err := conn.Read(buf); err != io.EOF && err != nil {
		return err
	}

	acceptedAtUnix, _ := binary.Varint(buf)
	acceptedAt := time.Unix(0, acceptedAtUnix)

	duration := acceptedAt.Sub(startAt)
	durations[*idx] = duration
	*idx++

	return nil
}
