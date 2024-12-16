package stages

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type HotelsMenuStage struct {
	keyboard interface{}
	message  string
}

func (hs *HotelsMenuStage) ConfigureMessage(message *tgbotapi.MessageConfig) {
	message.Text = hs.message
	message.ReplyMarkup = hs.keyboard
}

func NewHotelsMenuStage(role string) *HotelsMenuStage {
	hotelsMenuStage := new(HotelsMenuStage)
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Получить все"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Получить по ID"),
		))

	if role == "hotelier" {
		keyboard.Keyboard = append(keyboard.Keyboard,
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Создать отель"),
			))
	}

	keyboard.Keyboard = append(keyboard.Keyboard,
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Вернутся обратно"),
		))

	hotelsMenuStage.keyboard = keyboard

	hotelsMenuStage.message = "Меню отелей"
	return hotelsMenuStage
}
