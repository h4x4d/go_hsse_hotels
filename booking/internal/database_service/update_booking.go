package database_service

import (
	"context"
	"errors"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/grpc/client"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
	"github.com/jackc/pgx/v5/pgtype"
	"slices"
	"strings"
	"time"
)

func (ds *DatabaseService) Update(bookingId int64, booking *models.Booking) (*models.Booking, error) {
	query := `UPDATE bookings SET`
	var settings []string
	var values []interface{}

	if booking.DateFrom != nil {
		settings = append(settings, fmt.Sprintf("date_from = $%d", len(values)+1))
		date, err := time.Parse("02-01-2006", *booking.DateFrom)
		if err != nil {
			return nil, err
		}
		values = append(values, date.Format(time.DateOnly))
	}

	if booking.DateTo != nil {
		settings = append(settings, fmt.Sprintf("date_to = $%d", len(values)+1))
		date, err := time.Parse("02-01-2006", *booking.DateTo)
		if err != nil {
			return nil, err
		}
		values = append(values, date.Format(time.DateOnly))
	}

	if booking.HotelID != nil {
		settings = append(settings, fmt.Sprintf("hotel_id = $%d", len(values)+1))
		values = append(values, *booking.HotelID)
	}
	if booking.FullCost == 0 {
		hotel, err := client.GetHotelById(booking.HotelID)
		if err != nil {
			return nil, err
		}

		dFrom, dateErr1 := time.Parse("02-01-2006", *booking.DateFrom)
		dTo, dateErr2 := time.Parse("02-01-2006", *booking.DateTo)
		if dateErr1 != nil {
			return nil, dateErr1
		}
		if dateErr2 != nil {
			return nil, dateErr2
		}
		booking.FullCost = hotel.Cost * (int64(dTo.Sub(dFrom).Hours()) / 24)
	}

	settings = append(settings, fmt.Sprintf("full_cost = $%d", len(values)+1))
	values = append(values, booking.FullCost)

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
		fmt.Sprintf("id = $%d", len(values)+1))
	values = append(values, bookingId)

	from := new(pgtype.Date)
	to := new(pgtype.Date)

	errUpdate := ds.pool.QueryRow(context.Background(), query, values...).Scan(&booking.BookingID, from,
		to, booking.HotelID, &booking.FullCost, &booking.Status, &booking.UserID)

	fromStr := from.Time.Format("02-01-2006")
	toStr := to.Time.Format("02-01-2006")
	booking.DateFrom = &fromStr
	booking.DateTo = &toStr
	return booking, errUpdate
}
