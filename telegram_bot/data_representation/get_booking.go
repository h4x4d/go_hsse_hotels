package data_representation

import (
	"strconv"
	"telegram_bot/models"
)

func GetBooking(booking *models.Booking) string {
	var result string
	result += "Информация о бронировании\n"
	result += "ID бронирования: " + "\"" + strconv.FormatInt(booking.BookingID, 10) + "\"" + ";\n"
	result += "Период брони: " + "\"" + *booking.DateFrom + " - " + *booking.DateTo + "\"" + ";\n"
	result += "Статус бронирования " + "\"" + booking.Status + "\"" + ";\n"
	result += "Создан пользователем с ID " + "\"" + booking.UserID[:12] + "..." + "\"" + ";\n"
	result += "Относится к отелю с ID " + "\"" + strconv.FormatInt(*booking.HotelID, 10) + "\"" + ";\n"

	return result
}
