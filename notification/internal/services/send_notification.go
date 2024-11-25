package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/notification/internal/models"
	"net/http"
	"os"
	"strconv"
)

type SendNotificationRequest struct {
	ChatId string `json:"chat_id"`
	Text   string `json:"text"`
}

func SendNotification(notification models.Notification) error {
	apiKey := os.Getenv("TELEGRAM_API_KEY")
	if apiKey == "" {
		return errors.New("environment variable TELEGRAM_API_KEY is not set")
	}

	telegramAPIURL := fmt.Sprintf("https://api.telegram.org/bot%s", apiKey)

	requestBody, err := json.Marshal(SendNotificationRequest{
		ChatId: strconv.Itoa(notification.TelegramID),
		Text:   fmt.Sprintf("%s\n\n%s", notification.Name, notification.Text),
	})
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	resp, err := http.Post(telegramAPIURL+"/sendMessage", "application/json",
		bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}
	return nil
}
