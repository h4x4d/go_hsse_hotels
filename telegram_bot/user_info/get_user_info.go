package user_info

import (
	"telegram_bot/models"
)

func (us *UserInfo) GetUserData(telegramId int64) *models.User {
	user, exists := us.Data[telegramId]
	if !exists {
		user = models.NewUser()
		user.TelegramID = telegramId
		us.Data[telegramId] = user
	}
	return user
}
