package user_info

func (us *UserInfo) GetUserStage(telegramId int64) Stage {
	stage, exists := us.Stage[telegramId]
	if !exists {
		us.Stage[telegramId] = new(Stage)
		stage = us.Stage[telegramId]
	}
	return stage
}
