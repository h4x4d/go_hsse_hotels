package server

import (
	"errors"
	"github.com/segmentio/kafka-go"
)

func GetFormat(headers []kafka.Header) string {
	for _, header := range headers {
		if header.Key == "format" {
			return string(header.Value)
		}
	}
	return ""
}

func (server *NotificationKafkaServer) ProceedMessage(key []byte, value []byte, headers []kafka.Header) error {
	if GetFormat(headers) == "json" {
		return server.handlers[string(key)](value)
	}
	return errors.New("invalid format")
}
