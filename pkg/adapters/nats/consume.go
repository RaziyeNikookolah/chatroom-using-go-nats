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
	stream, err := n.js.Stream(ctx, subject)
	if err != nil {
		log.Fatalf("Failed to retrieve stream: %v", err)
	}

	_, err = stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Name:       subject + "-consumer",
		Durable:    subject + "-consumer",       // Persistent consumer for the stream
		AckPolicy:  jetstream.AckExplicitPolicy, // Require explicit acks
		AckWait:    30 * time.Second,            // Time before redelivery if not acked
		MaxDeliver: -1,
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	// log.Println("Consumer 'chatroom' is ready")
}
func (n *NATS) GetAllMessages(subject string) {
	ctx := context.Background()

	// Consumer ro misazim ke az aval hame payam ha ro daryaft kone
	stream, err := n.js.Stream(ctx, subject)
	if err != nil {
		log.Fatalf("Failed to retrieve stream: %v", err)
	}

	_, err = stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Name:          subject + "-consumer",
		Durable:       subject + "-consumer",
		AckPolicy:     jetstream.AckExplicitPolicy,
		DeliverPolicy: jetstream.DeliverAllPolicy,
		AckWait:       30 * time.Second,
		MaxDeliver:    -1,
		ReplayPolicy:  jetstream.ReplayInstantPolicy,
	})
	if err != nil {
		log.Fatalf("Failed to create or get consumer: %v", err)
	}

	consumer, err := n.js.Consumer(ctx, "chatroom", subject+"-consumer")
	if err != nil {
		log.Fatalf("Failed to get consumer: %v", err)
	}

	// Daryaft hame payamha
	msgs, err := consumer.Fetch(1000, jetstream.FetchMaxWait(10*time.Second))
	if err != nil {
		log.Fatalf("Failed to fetch messages: %v", err)
	}

	// Iteration ruye hame message-ha
	for msg := range msgs.Messages() {
		log.Printf("Consumed message: %s\n", string(msg.Data()))
		// meta, err := msg.Metadata()
		// if err != nil {
		// 	log.Printf("Failed to get message metadata: %v", err)
		// 	continue
		// }
		// // Get the ack timestamp
		// ackTime := meta.Timestamp
		// log.Printf("Message received at: %s\n", ackTime.Format(time.RFC3339))
		msg.Ack()
	}
	// consumerLister := stream.ConsumerNames(ctx)

	// // Print consumer names
	// log.Println("Consumers in stream:")
	// for c := range consumerLister.Name() {
	// 	log.Println(" -", c)
	// }
	err = stream.DeleteConsumer(ctx, subject+"-consumer")
	if err != nil {
		log.Fatalf("Failed to delete consumer: %v", err)
	}
}
