package user_info

import (
	"telegram_bot/models"
)

func (us *UserInfo) GetUserBooking(telegramId int64) *models.Booking {
	booking, exists := us.Booking[telegramId]
	if !exists {
		booking = models.NewBooking()
		us.Booking[telegramId] = booking
	}
	return booking
}
