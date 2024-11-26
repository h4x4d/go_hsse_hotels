package handlers

import (
	"encoding/json"
	"github.com/h4x4d/go_hsse_hotels/notification/internal/models"
	"github.com/h4x4d/go_hsse_hotels/notification/internal/services"
)

func SendNotificationHandler(value []byte) error {
	request := models.Notification{}
	err := json.Unmarshal(value, &request)
	if err != nil {
		return err
	}
	return services.SendNotification(request)
}
