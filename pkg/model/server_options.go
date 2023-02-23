package model

import (
	"time"
)

type ServerOptions struct {
	HandshakeTimeout time.Duration
	ReadBufferSize   int
	WriteBufferSize  int
}

func GetDefaultOptions() ServerOptions {
	return ServerOptions{
		HandshakeTimeout: time.Second * 5,
		ReadBufferSize:   1024 * 2,
		WriteBufferSize:  1024 * 2,
	}
}
