package user_info

func (us *UserInfo) UserStageExists(telegramId int64) bool {
	_, exists := us.Stage[telegramId]
	return exists
}
