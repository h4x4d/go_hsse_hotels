package server

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
	"os/signal"
	"syscall"
)

type KafkaServer interface {
	Serve() error
	ProceedMessage(key []byte, value []byte, headers []kafka.Header) error
}

type NotificationKafkaServer struct {
	Brokers *[]string
	Topic   *string
	GroupID *string

	handlers map[string]func([]byte) error
	reader   *kafka.Reader

	KafkaServer
}

func NewNotificationServer(brokers *[]string, topic *string, groupID *string,
	handlers map[string]func([]byte) error) *NotificationKafkaServer {
	return &NotificationKafkaServer{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: *brokers,
			Topic:   *topic,
			GroupID: *groupID,
		}),
		handlers: handlers,
	}
}

func (server *NotificationKafkaServer) Serve() error {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-stopChan
		fmt.Println("Shutdown signal received. Exiting gracefully...")
		cancel()
	}()

	reader := server.reader
	defer reader.Close()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Consumer stopped")
			return nil
		default:
			message, err := reader.ReadMessage(context.Background())
			if err != nil {
				return err
			}
			proceedErr := server.ProceedMessage(message.Key, message.Value, message.Headers)
			if proceedErr != nil {
				return proceedErr
			}
		}
	}
}
