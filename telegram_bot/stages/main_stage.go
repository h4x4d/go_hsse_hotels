package stages

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type MainStage struct {
	keyboard interface{}
	message  string
}

func (ms *MainStage) ConfigureMessage(message *tgbotapi.MessageConfig) {
	message.Text = ms.message
	message.ReplyMarkup = ms.keyboard
}

func NewMainStage() *MainStage {
	mainStage := new(MainStage)
	mainStage.keyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Отели"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Бронирования"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Выйти"),
		))
	mainStage.message = "Главное меню"
	return mainStage
}
