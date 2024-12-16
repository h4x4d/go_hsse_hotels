package stages

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram_bot/api_service"
	"telegram_bot/models"
	"telegram_bot/user_info"
)

type RegisterStage struct {
	InputStages
}

func (rs *RegisterStage) Finish(user *models.User, apiService *api_service.Service) (bool, error) {
	token, errRegister := apiService.Register(user)
	if errRegister != nil {
		return false, errRegister
	}
	return token != nil, nil
}

func NewRegisterStage() *RegisterStage {
	registerStage := new(RegisterStage)

	registerStage.InputStages = *NewInputStages()

	mailInput := InputStage{
		Message: "Введите почту",
		Input: func(userInfo *user_info.UserInfo, telegramId int64, mail string) error {
			user := userInfo.GetUserData(telegramId)
			*user.Mail = mail
			return nil
		},
		Keyboard: tgbotapi.NewRemoveKeyboard(true),
	}

	roleInput := InputStage{
		Message: "Выберите роль([hotelier, customer])",
		Input: func(userInfo *user_info.UserInfo, telegramId int64, role string) error {
			user := userInfo.GetUserData(telegramId)
			*user.Role = role
			return nil
		},
		Keyboard: tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("hotelier"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("customer"),
			)),
	}

	loginInput := InputStage{
		Message: "Введите логин",
		Input: func(userInfo *user_info.UserInfo, telegramId int64, login string) error {
			user := userInfo.GetUserData(telegramId)
			*user.Login = login
			return nil
		},
		Keyboard: tgbotapi.NewRemoveKeyboard(true),
	}

	passwordInput := InputStage{
		Message: "Введите пароль",
		Input: func(userInfo *user_info.UserInfo, telegramId int64, password string) error {
			user := userInfo.GetUserData(telegramId)
			*user.Password = password
			return nil
		},
		Keyboard: tgbotapi.NewRemoveKeyboard(true),
	}

	registerStage.AddStages(mailInput, roleInput, loginInput, passwordInput)
	return registerStage
}
