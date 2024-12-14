package handlers

import (
	"encoding/json"
	"github.com/h4x4d/go_hsse_hotels/notification/internal/models"
	"github.com/h4x4d/go_hsse_hotels/notification/internal/services"
	"log/slog"
	"net/http"
)

func SendNotificationHandler(value []byte) error {
	request := models.Notification{}
	err := json.Unmarshal(value, &request)
	if err != nil {
		return err
	}
	err = services.SendNotification(request)
	if err != nil {
		// Logging
		slog.Error(
			"failed send notification",
			slog.Group("notification-properties",
				slog.String("name", request.Name),
				slog.String("text", request.Text),
				slog.Int("telegram-id", request.TelegramID),
			),
			slog.Int("status_code", http.StatusInternalServerError),
			slog.String("error", err.Error()),
		)
	}
	return err
}
