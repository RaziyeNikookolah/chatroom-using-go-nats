package message_broker

import (
	"log"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/app"
)

func Run(app app.App) {
	messageBroker := app.MessageBroker()

	log.Println("Broker is started..")
	messageBroker.SetupStream()
	// messageBroker.SetupConsumer("chatroom")
	// go messageBroker.GetAllMessages("chatroom")
}

// go messageBroker.Consume("chatroom", "chatroom-consumer")
// go messageBroker.GetAllMessages("chatroom")
// go messageBroker.SubscribeToChat(context.Background(), "Razye")
// go messageBroker.Publish("chatroom", "Man Razye hastam ..")
// go messageBroker.SubscribeToChat(context.Background(), "Maryam")
// go messageBroker.Publish("chatroom", "Man Maryam hastam ..")
