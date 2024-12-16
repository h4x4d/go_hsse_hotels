package stages

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"telegram_bot/api_service"
	"telegram_bot/user_info"
)

type CreateBookingStage struct {
	InputStages
}

func (cbs *CreateBookingStage) Finish(userInfo *user_info.UserInfo, telegramId int64, apiService *api_service.Service) (bool, error) {
	user := userInfo.GetUserData(telegramId)
	booking := userInfo.GetUserBooking(telegramId)
	booking.UserID = user.UserID
	return apiService.CreateBooking(booking, user)
}

func NewCreateBookingStage() *CreateBookingStage {
	createBookingStage := new(CreateBookingStage)

	createBookingStage.InputStages = *NewInputStages()

	dateFromInput := InputStage{
		Message: "Введите дату начала (дд:мм:гггг)",
		Input: func(userInfo *user_info.UserInfo, telegramId int64, dateFrom string) error {
			booking := userInfo.GetUserBooking(telegramId)
			*booking.DateFrom = dateFrom
			return nil
		},
		Keyboard: tgbotapi.NewRemoveKeyboard(true),
	}

	dateToInput := InputStage{
		Message: "Введите дату конца (дд-мм-гггг)",
		Input: func(userInfo *user_info.UserInfo, telegramId int64, dateTo string) error {
			booking := userInfo.GetUserBooking(telegramId)
			*booking.DateTo = dateTo
			return nil
		},
		Keyboard: tgbotapi.NewRemoveKeyboard(true),
	}

	//fullCostInput := InputStage{
	//	Message: "Введите полную стоимость",
	//	Input: func(userInfo *user_info.UserInfo, telegramId int64, fullCost string) error {
	//		booking := userInfo.GetUserBooking(telegramId)
	//		var errParse error
	//		booking.FullCost, errParse = strconv.ParseInt(fullCost, 10, 64)
	//		if errParse != nil {
	//			return errParse
	//		}
	//		return nil
	//	},
	//}

	hotelIdInput := InputStage{
		Message: "Введите ID отеля",
		Input: func(userInfo *user_info.UserInfo, telegramId int64, hotelId string) error {
			booking := userInfo.GetUserBooking(telegramId)
			var errParse error
			*booking.HotelID, errParse = strconv.ParseInt(hotelId, 10, 64)
			log.Println(*booking.HotelID)
			if errParse != nil {
				return errParse
			}
			return nil
		},
		Keyboard: tgbotapi.NewRemoveKeyboard(true),
	}

	createBookingStage.AddStages(dateFromInput, dateToInput, hotelIdInput)
	return createBookingStage
}
