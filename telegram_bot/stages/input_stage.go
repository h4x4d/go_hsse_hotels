package stages

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram_bot/user_info"
)

type InputStage struct {
	Message  string
	Input    func(*user_info.UserInfo, int64, string) error
	Keyboard interface{}
}

func NewInputStage() *InputStage {
	inputStage := new(InputStage)
	inputStage.Keyboard = tgbotapi.NewRemoveKeyboard(true)
	return inputStage
}
