package stages

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"telegram_bot/user_info"
)

type BookingsGetBookingIdStage struct {
	InputStages
}

func NewBookingsGetBookingIdStage() *BookingsGetBookingIdStage {
	bookingsGetBookingIdStage := new(BookingsGetBookingIdStage)

	bookingsGetBookingIdStage.InputStages = *NewInputStages()

	idInput := InputStage{
		Message: "Введите ID бронирования",
		Input: func(userInfo *user_info.UserInfo, telegramId int64, bookingId string) error {
			booking := userInfo.GetUserBooking(telegramId)
			var errId error
			booking.BookingID, errId = strconv.ParseInt(bookingId, 10, 64)
			if errId != nil {
				return errId
			}
			return nil
		},
		Keyboard: tgbotapi.NewRemoveKeyboard(true),
	}

	bookingsGetBookingIdStage.AddStages(idInput)
	return bookingsGetBookingIdStage
}
