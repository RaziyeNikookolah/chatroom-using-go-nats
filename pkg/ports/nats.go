package ports

import "context"

type IMessageBroker interface {
	Publish(subject, msg string)
	Consume(subject, consumerName string)
	SetupStream()
	SetupConsumer(subject string)
	SubscribeToChat(ctx context.Context, username string)
	GetAllMessages(subject string)
}
