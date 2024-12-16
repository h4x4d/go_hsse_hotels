package user_info

import (
	"telegram_bot/models"
)

func (us *UserInfo) GetUserHotel(telegramId int64) *models.Hotel {
	hotel, exists := us.Hotel[telegramId]
	if !exists {
		hotel = models.NewHotel()
		us.Hotel[telegramId] = hotel
	}
	return hotel
}
