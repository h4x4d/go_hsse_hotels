package stages

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram_bot/user_info"
)

type InputStages struct {
	stages    []InputStage
	currStage int
}

func (is *InputStages) ConfigureMessage(message *tgbotapi.MessageConfig) {
	message.Text = is.GetCurrMessage()
	message.ReplyMarkup = is.GetCurrKeyboard()
}

func (is *InputStages) GetCurrMessage() string {
	return is.stages[is.currStage].Message
}

func (is *InputStages) GetCurrKeyboard() interface{} {
	return is.stages[is.currStage].Keyboard
}

func (is *InputStages) ComputeInput(userInfo *user_info.UserInfo, telegramId int64, input string) error {
	return is.stages[is.currStage].Input(userInfo, telegramId, input)
}

func (is *InputStages) EndStage() {
	is.currStage++
}

func (is *InputStages) StagesFinished() bool {
	return is.currStage >= len(is.stages)
}

func (is *InputStages) AddStages(stages ...InputStage) {
	is.stages = append(is.stages, stages...)
}

func NewInputStages() *InputStages {
	InputStages := new(InputStages)

	InputStages.stages = make([]InputStage, 0)

	return InputStages
}
