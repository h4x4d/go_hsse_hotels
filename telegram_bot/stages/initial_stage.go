package stages

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type InitialStage struct {
	keyboard interface{}
	message  string
}

func (is *InitialStage) ConfigureMessage(message *tgbotapi.MessageConfig) {
	message.Text = is.message
	message.ReplyMarkup = is.keyboard
}

func NewInitialStage() *InitialStage {
	initialStage := new(InitialStage)
	initialStage.keyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Войти"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Зарегистрироваться"),
		))
	initialStage.message = "Добро пожаловать в Hotel Booking Service!"
	return initialStage
}
