package database_service

import (
	"context"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
	"github.com/jackc/pgx/v5/pgtype"
)

func (ds *DatabaseService) GetByID(BookingID int64) (*models.Booking, error) {
	bookingRow, errGet := ds.pool.Query(context.Background(),
		"SELECT * FROM bookings WHERE id = $1", BookingID)
	if errGet != nil {
		return nil, errGet
	}
	defer bookingRow.Close()

	if !bookingRow.Next() {
		return nil, nil
	}

	booking := new(models.Booking)
	booking.HotelID = new(int64)

	from := new(pgtype.Date)
	to := new(pgtype.Date)

	// scaning booking object
	errBooking := bookingRow.Scan(&booking.BookingID, from,
		to, booking.HotelID, &booking.FullCost, &booking.Status, &booking.UserID)

	fromStr := from.Time.Format("02-01-2006")
	toStr := to.Time.Format("02-01-2006")
	booking.DateFrom = &fromStr
	booking.DateFrom = &toStr
	return booking, errBooking
}
