package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type NATS struct {
	serverURL string
	conn      *nats.Conn
	js        jetstream.JetStream
	jc        nats.JetStreamContext
}

// NewNATS creates a new NATS instance and connects to the NATS server.
func NewNATS(serverURL string) (*NATS, error) {
	// Connect to the NATS server
	nc, err := nats.Connect(serverURL, nats.Name("NATS Client"))
	if err != nil {
		return nil, err
	}

	// Create the JetStream context
	js, err := jetstream.New(nc)
	if err != nil {
		nc.Close() // Close the connection if JetStream setup fails
		return nil, err
	}

	// Get the JetStreamContext from the NATS connection
	jc, err := nc.JetStream()
	if err != nil {
		nc.Close() // Close the connection if JetStreamContext creation fails
		return nil, err
	}

	// Return a new NATS struct
	return &NATS{
		serverURL: serverURL,
		conn:      nc,
		js:        js,
		jc:        jc,
	}, nil
}
