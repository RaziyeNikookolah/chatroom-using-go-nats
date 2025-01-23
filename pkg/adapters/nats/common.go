package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type NATS struct {
	serverURL string
	conn      *nats.Conn
	js        jetstream.JetStream
}

func NewNATS(serverURL string) (*NATS, error) {
	nc, err := nats.Connect(serverURL, nats.Name("NATS Client"))
	if err != nil {
		return nil, err
	}

	js, err := jetstream.New(nc)
	if err != nil {
		return nil, err
	}

	return &NATS{serverURL: serverURL, conn: nc, js: js}, nil
}
