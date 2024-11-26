package database_service

import (
	"context"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
)

func (ds *DatabaseService) GetById(BookingID int64) (*models.Booking, error) {
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
	booking.DateTo = new(string)
	booking.DateFrom = new(string)
	booking.Status = new(string)
	booking.HotelID = new(int64)
	booking.RoomID = new(int64)

	// scaning booking object
	errBooking := bookingRow.Scan(&booking.BookingID, booking.DateFrom,
		booking.DateTo, booking.RoomID, booking.HotelID, &booking.UserID, &booking.FinalCost, booking.Status)
	bookingRow.Close()

	return booking, errBooking
}
