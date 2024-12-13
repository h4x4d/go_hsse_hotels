package models

type Notification struct {
	Name       string `json:"name"`
	Text       string `json:"text"`
	TelegramID int    `json:"telegram_id"`
}
