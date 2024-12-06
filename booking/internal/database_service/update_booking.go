package database_service

import (
	"context"
	"errors"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
	"slices"
	"strings"
)

func (ds *DatabaseService) Update(bookingId int64, booking *models.Booking) (*models.Booking, error) {
	query := `UPDATE bookings SET`
	var settings []string
	var values []interface{}

	if booking.DateFrom != nil {
		settings = append(settings, fmt.Sprintf("date_from = $%d", len(values)+1))
		values = append(values, *booking.DateFrom)
	}

	if booking.DateTo != nil {
		settings = append(settings, fmt.Sprintf("date_to = $%d", len(values)+1))
		values = append(values, *booking.DateTo)
	}

	if booking.HotelID != nil {
		settings = append(settings, fmt.Sprintf("hotel_id = $%d", len(values)+1))
		values = append(values, *booking.HotelID)
	}

	if booking.FullCost != 0 {
		settings = append(settings, fmt.Sprintf("full_cost = $%d", len(values)+1))
		values = append(values, booking.FullCost)
	}

	statuses := []string{"Waiting", "Payed", "Confirmed", "Finished"}
	if slices.Contains(statuses, booking.Status) {
		settings = append(settings, fmt.Sprintf("status = $%d", len(values)+1))
		values = append(values, booking.Status)
	} else if booking.Status != "" {
		return nil, errors.New("wrong status of the given booking")
	}

	if booking.UserID != "" {
		settings = append(settings, fmt.Sprintf("user_id = $%d", len(values)+1))
		values = append(values, booking.UserID)
	}

	query += fmt.Sprintf(" %s WHERE %s RETURNING *", strings.Join(settings, ", "),
		fmt.Sprintf("booking_id = $%d", len(values)+1))
	values = append(values, bookingId)
	errUpdate := ds.pool.QueryRow(context.Background(), query, values...).Scan(&booking.BookingID, booking.DateFrom,
		booking.DateTo, booking.HotelID, &booking.UserID, &booking.FullCost, booking.Status)
	return booking, errUpdate
}
