package stages

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"telegram_bot/user_info"
)

type BookingsGetHotelIdStage struct {
	InputStages
}

func NewBookingGetHotelIdStage() *BookingsGetHotelIdStage {
	bookingsGetHotelIdStage := new(BookingsGetHotelIdStage)

	bookingsGetHotelIdStage.InputStages = *NewInputStages()

	idInput := InputStage{
		Message: "Введите ID отеля",
		Input: func(userInfo *user_info.UserInfo, telegramId int64, hotelId string) error {
			booking := userInfo.GetUserBooking(telegramId)
			var errId error
			*booking.HotelID, errId = strconv.ParseInt(hotelId, 10, 64)
			if errId != nil {
				return errId
			}
			return nil
		},
		Keyboard: tgbotapi.NewRemoveKeyboard(true),
	}

	bookingsGetHotelIdStage.AddStages(idInput)
	return bookingsGetHotelIdStage
}
