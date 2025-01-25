package ports

import "context"

type IMessageBroker interface {
	Publish(msg string)
	Consume(consumerName string)
	SetupStream()
	SetupConsumer()
	SubscribeToChat(ctx context.Context, username string)
	GetAllMessages() ([]string, error)
}
