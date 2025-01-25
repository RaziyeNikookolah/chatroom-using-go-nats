package nats

import (
	"log"
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type NATS struct {
	serverURL string
	conn      *nats.Conn
	js        jetstream.JetStream
	jc        nats.JetStreamContext
}

// Singleton instance and once for thread-safe initialization
var (
	instance *NATS
	once     sync.Once
	err      error
)

// GetInstance returns a singleton NATS instance
func GetInstance() (*NATS, error) {
	once.Do(func() {
		instance, err = newNATS(nats.GetDefaultOptions().Url)
	})
	return instance, err
}

// Private function to initialize NATS connection
func newNATS(serverURL string) (*NATS, error) {
	nc, err := nats.Connect(serverURL, nats.Name("NATS Client"))
	if err != nil {
		return nil, err
	}

	js, err := jetstream.New(nc)
	if err != nil {
		nc.Close()
		return nil, err
	}

	jc, err := nc.JetStream()
	if err != nil {
		nc.Close()
		return nil, err
	}

	return &NATS{
		serverURL: serverURL,
		conn:      nc,
		js:        js,
		jc:        jc,
	}, nil
}

// Run initializes and starts the message broker
func Run(url string) {
	natsInstance, err := GetInstance()
	if err != nil {
		log.Fatalf("Failed to initialize NATS: %v", err)
		return
	}

	log.Println("Broker is started..", natsInstance.serverURL)
}
