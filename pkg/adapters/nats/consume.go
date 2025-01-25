package nats

import (
	"context"
	"log"
	"time"

	"github.com/nats-io/nats.go/jetstream"
)

func (n *NATS) Consume(consumerName string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create or retrieve the consumer
	consumer, err := n.js.Consumer(ctx, "chatroom", consumerName)
	if err != nil {
		log.Fatalf("Failed to get or create consumer %s: %v", consumerName, err)
	}

	log.Printf("Listening for messages on subject: %s", "chatroom")

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
func (n *NATS) SetupConsumer() {
	ctx := context.Background()
	stream, err := n.js.Stream(ctx, "chatroom")
	if err != nil {
		log.Fatalf("Failed to retrieve stream: %v", err)
	}

	_, err = stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Name:       "chatroom-consumer",
		Durable:    "chatroom-consumer",         // Persistent consumer for the stream
		AckPolicy:  jetstream.AckExplicitPolicy, // Require explicit acks
		AckWait:    30 * time.Second,            // Time before redelivery if not acked
		MaxDeliver: -1,
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	// log.Println("Consumer 'chatroom' is ready")
}
func (n *NATS) GetAllMessages() ([]string, error) {
	ctx := context.Background()

	// Consumer ro misazim ke az aval hame payam ha ro daryaft kone
	stream, err := n.js.Stream(ctx, "chatroom")
	if err != nil {
		log.Fatalf("Failed to retrieve stream: %v", err)
	}

	_, err = stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Name:          "chatroom-consumer",
		Durable:       "",
		AckPolicy:     jetstream.AckExplicitPolicy,
		DeliverPolicy: jetstream.DeliverAllPolicy,
		AckWait:       30 * time.Second,
		MaxDeliver:    -1,
		ReplayPolicy:  jetstream.ReplayInstantPolicy,
	})
	if err != nil {
		log.Fatalf("Failed to create or get consumer: %v", err)
	}

	consumer, err := n.js.Consumer(ctx, "chatroom", "chatroom-consumer")
	if err != nil {
		log.Fatalf("Failed to get consumer: %v", err)
	}

	// Daryaft hame payamha
	msgs, err := consumer.Fetch(1000, jetstream.FetchMaxWait(10*time.Second))
	if err != nil {
		log.Fatalf("Failed to fetch messages: %v", err)
	}

	var messages []string
	for msg := range msgs.Messages() {
		// log.Printf("Consumed message: %s\n", string(msg.Data()))
		messages = append(messages, string(msg.Data()))

	}

	err = stream.DeleteConsumer(ctx, "chatroom-consumer")
	if err != nil {
		log.Fatalf("Failed to delete consumer: %v", err)
	}
	return messages, nil
}
