package data_representation

import (
	"strconv"
	"telegram_bot/models"
)

func GetHotel(hotel *models.Hotel) string {
	var result string
	result += "Информация об отеле;\n"
	result += "ID отеля: " + "\"" + strconv.FormatInt(hotel.ID, 10) + "\"" + ";\n"
	result += "Название отеля: " + "\"" + *hotel.Name + "\"" + "\n"
	result += "Адрес отеля: " + "\"" + *hotel.City + ", " + *hotel.Address + "\"" + ";\n"
	result += "Стоимость отеля: " + strconv.FormatInt(hotel.Cost, 10) + "₽" + ";\n"
	result += "Класс отеля: " + "\"" + strconv.FormatInt(hotel.HotelClass, 10) + "\"" + ";\n"
	result += "Создан пользователем с ID " + "\"" + hotel.UserID[:12] + "..." + "\"" + ";\n"

	return result
}
