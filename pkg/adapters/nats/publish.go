package nats

import (
	"context"
	"log"
	"time"

	"github.com/nats-io/nats.go/jetstream"
)

func (n *NATS) Publish(subject, msg string) {
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
	ctx := context.Background()
	_, err := n.js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:        "chatroom",
		Subjects:    []string{"chatroom.>"},
		Description: "chatroom in nats",
		MaxBytes:    1024 * 1024 * 1024,
	})
	if err != nil {
		log.Fatalf("Failed to create stream: %v", err)
	}
	log.Println("Stream 'chatroom' is ready")
}
