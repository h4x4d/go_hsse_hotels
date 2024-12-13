package notification

import (
	"encoding/json"
	"github.com/h4x4d/go_hsse_hotels/pkg/models"
	"github.com/segmentio/kafka-go"
)

func (kc KafkaConnection) SendNotification(notification models.Notification) error {
	notify, marshalErr := json.Marshal(notification)
	if marshalErr != nil {
		return marshalErr
	}
	_, err := kc.Writer.WriteMessages(
		kafka.Message{
			Key:   []byte("send_notification"),
			Value: notify,
			Headers: []kafka.Header{
				{
					Key:   "format",
					Value: []byte("json"),
				},
			},
		})
	if err != nil {
		return err
	}
	return nil
}
