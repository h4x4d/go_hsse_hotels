package stages

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"telegram_bot/api_service"
	"telegram_bot/user_info"
)

type CreateHotelStage struct {
	InputStages
}

func (cbs *CreateHotelStage) Finish(userInfo *user_info.UserInfo, telegramId int64, apiService *api_service.Service) (bool, error) {
	user := userInfo.GetUserData(telegramId)
	hotel := userInfo.GetUserHotel(telegramId)
	hotel.UserID = user.UserID
	return apiService.CreateHotel(hotel, user)
}

func NewCreateHotelStage() *CreateHotelStage {
	createHotelStage := new(CreateHotelStage)

	createHotelStage.InputStages = *NewInputStages()

	nameInput := InputStage{
		Message: "Введите название",
		Input: func(userInfo *user_info.UserInfo, telegramId int64, name string) error {
			hotel := userInfo.GetUserHotel(telegramId)
			*hotel.Name = name
			return nil
		},
		Keyboard: tgbotapi.NewRemoveKeyboard(true),
	}

	addressInput := InputStage{
		Message: "Введите адрес",
		Input: func(userInfo *user_info.UserInfo, telegramId int64, address string) error {
			hotel := userInfo.GetUserHotel(telegramId)
			*hotel.Address = address
			return nil
		},
		Keyboard: tgbotapi.NewRemoveKeyboard(true),
	}

	cityInput := InputStage{
		Message: "Введите город",
		Input: func(userInfo *user_info.UserInfo, telegramId int64, city string) error {
			hotel := userInfo.GetUserHotel(telegramId)
			*hotel.City = city
			return nil
		},
		Keyboard: tgbotapi.NewRemoveKeyboard(true),
	}

	costInput := InputStage{
		Message: "Введите стоимость за день (в рублях)",
		Input: func(userInfo *user_info.UserInfo, telegramId int64, cost string) error {
			hotel := userInfo.GetUserHotel(telegramId)
			var errorParse error
			hotel.Cost, errorParse = strconv.ParseInt(cost, 10, 64)
			if errorParse != nil {
				return errorParse
			}
			return nil
		},
		Keyboard: tgbotapi.NewRemoveKeyboard(true),
	}

	classInput := InputStage{
		Message: "Введите класс отеля([0, 1, 2, 3, 4, 5])",
		Input: func(userInfo *user_info.UserInfo, telegramId int64, class string) error {
			hotel := userInfo.GetUserHotel(telegramId)
			var errorParse error
			hotel.HotelClass, errorParse = strconv.ParseInt(class, 10, 64)
			if errorParse != nil {
				return errorParse
			}
			return nil
		},
		Keyboard: tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("0"),
				tgbotapi.NewKeyboardButton("1"),
				tgbotapi.NewKeyboardButton("2"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("3"),
				tgbotapi.NewKeyboardButton("4"),
				tgbotapi.NewKeyboardButton("5"),
			)),
	}

	createHotelStage.AddStages(nameInput, addressInput, cityInput, costInput, classInput)
	return createHotelStage
}
