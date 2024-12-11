package database_service

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/grpc/client"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
	"go.opentelemetry.io/otel"
	"log"
	"strings"
	"time"
)

func (ds *DatabaseService) CreateBooking(booking *models.Booking) (*int64, error) {
	query := `INSERT INTO bookings`
	// maybe fieldNames can be placed in common place cause other methods also need this info
	var fieldNames []string
	var fields []string
	var values []interface{}

	if booking.DateFrom != nil {
		fieldNames = append(fieldNames, "date_from")
		date, err := time.Parse("02-01-2006", *booking.DateFrom)
		if err != nil {
			return nil, err
		}
		values = append(values, date.Format(time.DateOnly))
	}
	if booking.DateTo != nil {
		fieldNames = append(fieldNames, "date_to")
		date, err := time.Parse("02-01-2006", *booking.DateTo)
		if err != nil {
			return nil, err
		}
		values = append(values, date.Format(time.DateOnly))
	}
	if booking.HotelID != nil {
		fieldNames = append(fieldNames, "hotel_id")
		values = append(values, booking.HotelID)
	}

	if booking.BookingID != 0 {
		fieldNames = append(fieldNames, "booking_id")
		values = append(values, booking.BookingID)
	}

	fieldNames = append(fieldNames, "status")
	values = append(values, booking.Status)
	fieldNames = append(fieldNames, "full_cost")
	values = append(values, booking.FullCost)
	fieldNames = append(fieldNames, "user_id")
	values = append(values, booking.UserID)

	for ind := 0; ind < len(fieldNames); ind++ {
		fields = append(fields, fmt.Sprintf("$%d", ind+1))
	}
	query += fmt.Sprintf(" (%s) VALUES (%s) RETURNING id", strings.Join(fieldNames, ", "),
		strings.Join(fields, ", "))
	log.Println(query)
	errInsert := ds.pool.QueryRow(context.Background(), query, values...).Scan(&booking.BookingID)
	if errInsert != nil {
		return nil, errInsert
	}

	return &booking.BookingID, errInsert
}

func (ds *DatabaseService) Create(ctx context.Context, dateFrom *string, dateTo *string, hotelID *int64, userID string) (*int64, error) {
	// Tracing
	tracer := otel.Tracer("Booking")
	childCtx, span := tracer.Start(ctx, "create booking in database")
	defer span.End()

	hotel, err := client.GetHotelById(childCtx, hotelID)

	if err != nil {
		return nil, err
	}
	dFrom, dateErr1 := time.Parse("02-01-2006", *dateFrom)
	dTo, dateErr2 := time.Parse("02-01-2006", *dateTo)
	if dateErr1 != nil {
		return nil, dateErr1
	}
	if dateErr2 != nil {
		return nil, dateErr2
	}
	cost := hotel.Cost * (int64(dTo.Sub(dFrom).Hours()) / 24)

	booking := &models.Booking{
		DateFrom: dateFrom,
		DateTo:   dateTo,
		HotelID:  hotelID,
		FullCost: cost,
		Status:   "Unpayed",
		UserID:   userID,
	}

	return ds.CreateBooking(booking)
}
