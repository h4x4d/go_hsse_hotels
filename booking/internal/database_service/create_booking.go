package database_service

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
	"strings"
)

func (ds *DatabaseService) Create(booking *models.Booking) (*int64, error) {
	query := `INSERT INTO bookings`
	var fieldNames []string
	var fields []string
	var values []interface{}

	if booking.DateFrom != nil {
		fieldNames = append(fieldNames, "date_from")
		values = append(values, booking.DateFrom)
	}
	if booking.DateTo != nil {
		fieldNames = append(fieldNames, "date_to")
		values = append(values, booking.DateTo)
	}
	if booking.HotelID != nil {
		fieldNames = append(fieldNames, "hotel_id")
		values = append(values, booking.HotelID)
	}
	if booking.RoomID != nil {
		fieldNames = append(fieldNames, "room_id")
		values = append(values, booking.RoomID)
	}
	if booking.Status != nil {
		fieldNames = append(fieldNames, "status")
		values = append(values, booking.Status)
	}

	if booking.BookingID != 0 {
		fieldNames = append(fieldNames, "booking_id")
		values = append(values, booking.BookingID)
	}

	fieldNames = append(fieldNames, "full_cost")
	values = append(values, booking.FinalCost)
	fieldNames = append(fieldNames, "user_id")
	values = append(values, booking.UserID)

	for ind := 0; ind < len(fieldNames); ind++ {
		fields = append(fields, fmt.Sprintf("$%d", ind+1))
	}
	query += fmt.Sprintf(" (%s) VALUES (%s) RETURNING id", strings.Join(fieldNames, ", "),
		strings.Join(fields, ", "))
	errInsert := ds.pool.QueryRow(context.Background(), query, values...).Scan(&booking.BookingID)
	if errInsert != nil {
		return nil, errInsert
	}

	return &booking.BookingID, errInsert
}
