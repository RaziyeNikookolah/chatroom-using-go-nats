package ports

type IMessageBroker interface {
	Publish(subject, msg string)
	Consume(subject, consumerName string)
	SetupStream()
}
