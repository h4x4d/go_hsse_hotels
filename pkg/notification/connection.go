package notification

import (
	"context"
	"github.com/segmentio/kafka-go"
	"os"
)

type KafkaConnection struct {
	Writer *kafka.Conn
}

func (kc KafkaConnection) Close() error {
	err := kc.Writer.Close()
	return err
}

func NewKafkaConnection(broker string, topic string) (*KafkaConnection, error) {
	writer, err := kafka.DialLeader(context.Background(), "tcp", broker, topic, 0)
	if err != nil {
		return nil, err
	}
	return &KafkaConnection{Writer: writer}, nil
}

func NewEnvKafkaConnection() (*KafkaConnection, error) {
	return NewKafkaConnection(
		os.Getenv("KAFKA_BROKER"),
		os.Getenv("KAFKA_TOPIC"),
	)
}
