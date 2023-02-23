package fac

import (
	"io"

	"github.com/mis-hashemi/goose/pkg/model"
)

type GooseServerFactory interface {
	// Listen will set up a server on given url (127.0.0.1:8080)
	Listen(url string, options ...model.ServerOption) (*model.WebSocketServer, error)
	// Listen will set up a server on given writer and reader
	ListenOn(writer io.Writer, reader io.Reader, options ...model.ServerOption) (*model.WebSocketServer, error)
	// StartTestServer will set up a server for testing purpose. Note that it does not require any
	StartTestServer(options ...model.ServerOption) (*model.WebSocketServer, error)
}
