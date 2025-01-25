package nats

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/client/domain"
	"github.com/nats-io/nats.go/jetstream"
)

func (n *NATS) ConsumeUnreadMessages(consumerName string) ([]string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var messages []string

	// Try to get the existing consumer or create it if it doesn't exist
	consumer, err := n.js.Consumer(ctx, "chatroom", consumerName)
	if err != nil {
		// Create a new consumer if not found
		consumer, err = n.js.CreateOrUpdateConsumer(ctx, "chatroom", jetstream.ConsumerConfig{
			Name:          consumerName,
			Durable:       consumerName, // Durable to track progress
			AckPolicy:     jetstream.AckExplicitPolicy,
			DeliverPolicy: jetstream.DeliverNewPolicy, // Only unread messages
			AckWait:       30 * time.Second,
			MaxDeliver:    -1,
			ReplayPolicy:  jetstream.ReplayInstantPolicy,
		})
		if err != nil {
			log.Fatalf("Failed to create consumer %s: %v", consumerName, err)
			return messages, err
		}
	}

	// log.Printf("Listening for unread messages on subject: %s", "chatroom")

	// Pull unread messages from the stream
	cctx, err := consumer.Consume(func(msg jetstream.Msg) {
		// log.Printf("Received message: %s", string(msg.Subject()))
		messages = append(messages, string(msg.Data()))
		msg.Ack()
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cctx.Stop()

	return messages, nil
}
func (n *NATS) GetAllMessages(username string) ([]domain.Message, error) {
	ctx := context.Background()

	// Retrieve the stream
	stream, err := n.js.Stream(ctx, "chatroom")
	if err != nil {
		log.Fatalf("Failed to retrieve stream: %v", err)
	}

	// Create a unique consumer name
	name := username + time.Now().Format("150405")
	_, err = stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Name:          name,
		Durable:       "",
		AckPolicy:     jetstream.AckExplicitPolicy,
		DeliverPolicy: jetstream.DeliverAllPolicy, // Receive all messages from the start
		AckWait:       30 * time.Second,
		MaxDeliver:    -1,
		ReplayPolicy:  jetstream.ReplayInstantPolicy,
	})
	if err != nil {
		log.Fatalf("Failed to create or get consumer: %v", err)
	}

	// Get consumer
	consumer, err := n.js.Consumer(ctx, "chatroom", name)
	if err != nil {
		log.Fatalf("Failed to get consumer: %v", err)
	}

	// Fetch messages
	msgs, err := consumer.Fetch(1000, jetstream.FetchMaxWait(10*time.Second))
	if err != nil {
		log.Fatalf("Failed to fetch messages: %v", err)
	}

	var messages []domain.Message
	for msg := range msgs.Messages() {
		var receivedMsg domain.Message
		err := json.Unmarshal(msg.Data(), &receivedMsg)
		if err != nil {
			log.Printf("Error decoding message: %v", err)
			continue
		}

		// Convert timestamp to desired format
		parsedTime, err := time.Parse(time.RFC3339, receivedMsg.Timestamp)
		if err == nil {
			receivedMsg.Timestamp = parsedTime.Format("02 January 15:04")
		}

		messages = append(messages, receivedMsg)
	}

	// Delete consumer after fetching messages
	defer func() {
		err = stream.DeleteConsumer(ctx, name)
		if err != nil {
			log.Fatalf("Failed to delete consumer: %v", err)
		}
	}()

	return messages, nil
}
func (n *NATS) GetActiveUsers() ([]string, error) {
	ctx := context.Background()
	var activeConsumers []string
	// Consumer ro misazim ke az aval hame payam ha ro daryaft kone
	stream, err := n.js.Stream(ctx, "chatroom")
	if err != nil {
		log.Fatalf("Failed to retrieve stream: %v", err)
	}

	consumerList := stream.ConsumerNames(ctx)

	for c := range consumerList.Name() {
		consumer, err := stream.Consumer(ctx, c)
		if err != nil {
			log.Printf("Failed to retrieve consumer info for %s: %v", c, err)
			continue
		}

		// Check if the consumer has active interest
		info, err := consumer.Info(ctx)
		if err != nil {
			log.Printf("Failed to get consumer info for %s: %v", c, err)
			continue
		}

		if info.PushBound {
			activeConsumers = append(activeConsumers, c)
		}
	}

	return activeConsumers, nil
}
