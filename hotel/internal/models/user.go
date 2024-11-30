package models

type User struct {
	UserId     string `json:"user_id"`
	Role       string `json:"role"`
	TelegramId int    `json:"telegram_id"`
}
