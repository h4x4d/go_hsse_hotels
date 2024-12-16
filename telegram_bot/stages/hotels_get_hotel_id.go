package stages

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"telegram_bot/user_info"
)

type HotelsGetHotelIdStage struct {
	InputStages
}

func NewHotelsGetHotelIdStage() *HotelsGetHotelIdStage {
	hotelsGetHotelIdStage := new(HotelsGetHotelIdStage)

	hotelsGetHotelIdStage.InputStages = *NewInputStages()

	idInput := InputStage{
		Message: "Введите ID отеля",
		Input: func(userInfo *user_info.UserInfo, telegramId int64, hotelId string) error {
			hotel := userInfo.GetUserHotel(telegramId)
			var errId error
			hotel.ID, errId = strconv.ParseInt(hotelId, 10, 64)
			if errId != nil {
				return errId
			}
			return nil
		},
		Keyboard: tgbotapi.NewRemoveKeyboard(true),
	}

	hotelsGetHotelIdStage.AddStages(idInput)
	return hotelsGetHotelIdStage
}
