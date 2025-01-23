package nats

import (
	"context"
	"log"
	"time"

	"github.com/nats-io/nats.go/jetstream"
)

func (n *NATS) Consume(subject, consumerName string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create or retrieve the consumer
	consumer, err := n.js.Consumer(ctx, subject, consumerName)
	if err != nil {
		log.Fatalf("Failed to get or create consumer %s: %v", consumerName, err)
	}

	log.Printf("Listening for messages on subject: %s", subject)

	for {
		// Pull messages from the stream with a timeout
		_, err := consumer.Fetch(10, jetstream.FetchMaxWait(5*time.Second))
		if err != nil {
			log.Printf("Error fetching messages: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}
	}
}
func (n *NATS) SetupConsumer(subject string) {
	ctx := context.Background()
	stream, err := n.js.Stream(ctx, "test")
	if err != nil {
		log.Fatalf("Failed to retrieve stream: %v", err)
	}

	_, err = stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Name:    "chatroom",
		Durable: "chatroom", // Persistent consumer for the stream
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	log.Println("Consumer 'chatroom' is ready")
}
