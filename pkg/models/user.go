package models

type User struct {
	UserID     string `json:"user_id"`
	Role       string `json:"role"`
	TelegramId int    `json:"telegram_id"`
}
