package main

import (
	"github.com/h4x4d/go_hsse_hotels/notification/internal/handlers"
	"github.com/h4x4d/go_hsse_hotels/notification/internal/server"
	"log"
	"os"
)

func main() {
	broker := os.Getenv("KAFKA_BROKER")
	topic := os.Getenv("KAFKA_TOPIC")
	groupID := os.Getenv("KAFKA_GROUP_ID")

	notify_handlers := map[string]func([]byte) error{
		"send_notification": handlers.SendNotificationHandler,
	}

	notificationServer := server.NewNotificationServer(&[]string{broker}, &topic, &groupID, notify_handlers)

	if err := notificationServer.Serve(); err != nil {
		log.Fatalln(err)
	}
}
