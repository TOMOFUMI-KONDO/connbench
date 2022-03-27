package client

import (
	"crypto/tls"
	"encoding/binary"

	"github.com/TOMOFUMI-KONDO/connbench"
)

func GenTLSCfg() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{connbench.NextProto},
	}
}

func Int64ToBytes(n int64) []byte {
	bytes := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(bytes, n)
	return bytes
}
