package database_service

import (
	"context"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
)

func (ds *DatabaseService) GetAll() ([]*models.Booking, error) {
	bookingIdRow, errGetId := ds.pool.Query(context.Background(), "SELECT booking_id FROM bookings")
	if errGetId != nil {
		return nil, errGetId
	}
	defer bookingIdRow.Close()

	bookings := make([]*models.Booking, 0)

	for bookingIdRow.Next() {
		var bookingId int64
		errScanId := bookingIdRow.Scan(&bookingId)
		if errScanId != nil {
			return nil, errScanId
		}
		booking, errGetBooking := ds.GetById(bookingId)
		if errGetBooking != nil {
			return nil, errGetBooking
		}
		bookings = append(bookings, booking)
	}
	return bookings, nil
}
