package nats

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func (n *NATS) Publish(ctx context.Context, username, message string) (output string) {
	msgStruct := struct {
		Sender    string `json:"sender"`
		Content   string `json:"content"`
		Timestamp string `json:"timestamp"`
	}{
		Sender:    username,
		Content:   message,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	msgBytes, err := json.Marshal(msgStruct)
	if err != nil {
		output = fmt.Sprintf("Error marshaling message: %v", err)
		return
	}

	_, err = n.js.Publish(ctx, "chatroom", msgBytes)
	if err != nil {
		output = fmt.Sprintf("Error sending message: %v", err)
	} else {
		output = "Message sent successfully"
	}
	return
}

func (n *NATS) SubscribeToChat(ctx context.Context, username string) {
	if n.jc == nil {
		log.Fatal("NATS client is not initialized")
	}

	// Subscribe without acknowledging messages
	_, err := n.jc.Subscribe("chatroom", func(msg *nats.Msg) {
		// Simply consume the message without acknowledgment
		select {
		case <-ctx.Done():
			fmt.Println("User inactive for 5 minutes, stopping listener.")
			return
		default:
			// No msg.Ack() here to leave messages unacknowledged
		}
	}, nats.Durable(username), nats.ManualAck())
	if err != nil {
		log.Fatal(err)
	}

	// Simulate listening period before cancelling context
	go func() {
		<-ctx.Done()
		fmt.Println("Goroutine for receiving messages has stopped.")
	}()

	// Cancel context after a delay (simulate ending subscription)
	_, cancel := context.WithCancel(ctx)
	defer cancel() // Ensure cancellation happens at the end of the function
}
