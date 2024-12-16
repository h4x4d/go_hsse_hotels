package user_info

import (
	"telegram_bot/api_service"
	"telegram_bot/models"
)

type Stage interface{}

type UserInfo struct {
	Stage   map[int64]Stage
	Data    map[int64]*models.User
	Hotel   map[int64]*models.Hotel
	Booking map[int64]*models.Booking
	Service map[int64]api_service.Service
}

func NewUserInfo() *UserInfo {
	userInfo := new(UserInfo)
	userInfo.Data = make(map[int64]*models.User)
	userInfo.Stage = make(map[int64]Stage)
	userInfo.Hotel = make(map[int64]*models.Hotel)
	userInfo.Booking = make(map[int64]*models.Booking)
	return userInfo
}
