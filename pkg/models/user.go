package models

type User struct {
	UserID     string `json:"user_id"`
	Role       string `json:"role"`
	TelegramID int    `json:"telegram_id"`
}
