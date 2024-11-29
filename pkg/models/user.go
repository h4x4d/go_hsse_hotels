package models

type User struct {
	UserId     string `json:"userId"`
	Role       string `json:"role"`
	TelegramId int    `json:"telegram_id"`
}
