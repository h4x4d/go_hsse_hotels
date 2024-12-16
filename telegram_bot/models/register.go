package models

type Register struct {
	Email string `json:"email"`
	Login
	Role       string `json:"role"`
	TelegramID int64  `json:"telegram_id"`
}
