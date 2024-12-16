package stages

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	inner_models "github.com/h4x4d/go_hsse_hotels/pkg/models"
	"telegram_bot/api_service"
	"telegram_bot/models"
	"telegram_bot/user_info"
)

type LoginStage struct {
	InputStages
}

func (rs *LoginStage) Finish(user *models.User, apiService *api_service.Service) (bool, error) {
	resultLogin, errLogin := apiService.Login(user)
	if errLogin != nil || !resultLogin {
		return resultLogin, errLogin
	}

	userToken, errorToken := apiService.GetToken(user.TelegramID)
	if errorToken != nil {
		return false, errorToken
	}

	// setting role and user_id
	var innerUser *inner_models.User
	var errorUser error
	innerUser, errorUser = apiService.CheckToken(userToken.Value)

	if errorUser != nil {
		return false, errorUser
	}
	user.Role = &innerUser.Role
	user.UserID = innerUser.UserID
	user.TelegramID = int64(innerUser.TelegramID)

	return true, nil
}

func NewLoginStage() *LoginStage {
	loginStage := new(LoginStage)

	loginStage.InputStages = *NewInputStages()

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

	loginStage.AddStages(loginInput, passwordInput)
	return loginStage
}
