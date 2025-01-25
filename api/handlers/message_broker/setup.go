package message_broker

import (
	"log"

	"github.com/RaziyeNikookolah/chatroom-using-go-nats/app"
)

func Run(app app.App) {
	messageBroker := app.MessageBroker()

	log.Println("Broker is started..")
	messageBroker.SetupStream()

}
