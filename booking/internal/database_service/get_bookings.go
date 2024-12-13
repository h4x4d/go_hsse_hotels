package database_service

import (
	"context"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
)

// GetAll actually there is no need in this method because api does not contains such
func (ds *DatabaseService) GetAll(HotelID *int64) ([]*models.Booking, error) {
	bookingIdRow, errGetId := ds.pool.Query(context.Background(),
		"SELECT id FROM bookings WHERE hotel_id=$1", *HotelID)
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
		booking, errGetBooking := ds.GetByID(bookingId)
		if errGetBooking != nil {
			return nil, errGetBooking
		}
		bookings = append(bookings, booking)
	}
	return bookings, nil
}
