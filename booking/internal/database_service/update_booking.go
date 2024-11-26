package database_service

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
	"strings"
)

func (ds *DatabaseService) Update(bookingId int64, booking *models.Booking) (*models.Booking, error) {
	query := `UPDATE bookings SET`
	var fieldNames []string
	var values []interface{}

	if booking.DateFrom != nil {
		fieldNames = append(fieldNames, fmt.Sprintf("date_from = $%d", len(values)+1))
		values = append(values, *booking.DateFrom)
	}
	if booking.DateTo != nil {
		fieldNames = append(fieldNames, fmt.Sprintf("date_to = $%d", len(values)+1))
		values = append(values, *booking.DateTo)
	}
	if booking.HotelID != nil {
		fieldNames = append(fieldNames, fmt.Sprintf("hotel_id = $%d", len(values)+1))
		values = append(values, *booking.HotelID)
	}
	if booking.RoomID != nil {
		fieldNames = append(fieldNames, fmt.Sprintf("room_id = $%d", len(values)+1))
		values = append(values, *booking.RoomID)
	}
	if booking.Status != nil {
		fieldNames = append(fieldNames, fmt.Sprintf("status = $%d", len(values)+1))
		values = append(values, *booking.Status)
	}

	if booking.BookingID != 0 {
		fieldNames = append(fieldNames, fmt.Sprintf("booking_id = $%d", len(values)+1))
		values = append(values, booking.BookingID)
	}

	if booking.FinalCost != 0 {
		fieldNames = append(fieldNames, fmt.Sprintf("full_cost = $%d", len(values)+1))
		values = append(values, booking.FinalCost)
	}

	if booking.UserID != 0 {
		fieldNames = append(fieldNames, fmt.Sprintf("user_id = $%d", len(values)+1))
		values = append(values, booking.UserID)
	}

	query += fmt.Sprintf(" %s WHERE %s RETURNING *", strings.Join(fieldNames, ", "),
		fmt.Sprintf("booking_id = $%d", len(values)+1))
	values = append(values, bookingId)
	fmt.Println(query, values)
	errUpdate := ds.pool.QueryRow(context.Background(), query, values...).Scan(&booking.BookingID, booking.DateFrom,
		booking.DateTo, booking.RoomID, booking.HotelID, &booking.UserID, &booking.FinalCost, booking.Status)
	return booking, errUpdate
}
