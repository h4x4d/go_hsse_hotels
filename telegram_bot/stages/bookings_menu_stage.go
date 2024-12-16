package stages

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type BookingsMenuStage struct {
	keyboard interface{}
	message  string
}

func (bs *BookingsMenuStage) ConfigureMessage(message *tgbotapi.MessageConfig) {
	message.Text = bs.message
	message.ReplyMarkup = bs.keyboard
}

func NewBookingsMenuStage(role string) *BookingsMenuStage {
	bookingsMenuStage := new(BookingsMenuStage)
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Получить по ID"),
		))
	switch role {
	case "customer":
		keyboard.Keyboard = append(keyboard.Keyboard,
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Забронировать отель"),
			))
	case "hotelier":
		keyboard.Keyboard = append(keyboard.Keyboard,
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Получить по ID отеля"),
			))
	}

	keyboard.Keyboard = append(keyboard.Keyboard,
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Вернутся обратно"),
		))

	bookingsMenuStage.keyboard = keyboard

	bookingsMenuStage.message = "Меню бронирования"
	return bookingsMenuStage
}
