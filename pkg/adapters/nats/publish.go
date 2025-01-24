package nats

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func (n *NATS) Publish(subject, msg string) {
	print(11111111, "\n")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := n.js.Publish(ctx, subject, []byte(msg))
	if err != nil {
		log.Printf("Failed to publish message to subject %s: %v", subject, err)
		return
	}

	log.Printf("Message successfully published to subject: %s", subject)
}
func (n *NATS) SetupStream() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if n.jc == nil {
		log.Fatal("NATS client is not initialized")
	}
	_, err := n.js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:        "chatroom",
		Subjects:    []string{"chatroom"},
		Description: "chatroom in nats",
		MaxBytes:    1024 * 1024 * 1024,
		// Enable persistence for at-least-once delivery
		Storage:   jetstream.FileStorage,  // Use file storage for durability
		Retention: jetstream.LimitsPolicy, // Keep messages based on limits

		// Enable acknowledgments and redelivery
		Discard:  jetstream.DiscardOld,
		Replicas: 1,     // Set replication factor (increase for HA)
		NoAck:    false, // Require acknowledgment from consumers
	})
	if err != nil {
		log.Fatalf("Failed to create stream: %v", err)
	}
	// log.Println("Stream 'chatroom' is ready")
}
func (n *NATS) SubscribeToChat(ctx context.Context, username string) {
	if n.jc == nil {
		log.Fatal("NATS client is not initialized")
	}
	_, err := n.jc.Subscribe("chatroom", func(msg *nats.Msg) {
		// Checking for inactivity timeout
		select {
		case <-ctx.Done():
			fmt.Println("User inactive for 5 minutes, stopping listener.")
			return
		default:
			fmt.Printf("Received message: %s\n", string(msg.Data))
			msg.Ack()
		}
	}, nats.Durable(username), nats.ManualAck())
	if err != nil {
		log.Fatal(err)
	}
	// i think here by store user status in redis when user is inactive or when get ack as a active user we can have active users list
	// here i think we should send data to client and not to print here
	// Waiting for messages until context is canceled
	<-ctx.Done()
	fmt.Println("Goroutine for receiving messages has stopped.")
}
