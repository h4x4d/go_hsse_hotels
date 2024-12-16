package user_info

func (us *UserInfo) SetUserStage(telegramId int64, stage Stage) {
	us.Stage[telegramId] = stage
}
