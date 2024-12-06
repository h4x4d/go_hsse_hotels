package database_service

import (
	"context"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
	"strings"
)

// GetAll actually there is no need in this method because api does not contains such
func (ds *DatabaseService) GetAll(HotelID *int64) ([]*models.Booking, error) {
	query := "SELECT booking_id FROM bookings"
	var conditions []string
	var values []interface{}

	if HotelID != nil {
		conditions = append(conditions, fmt.Sprintf("hotel_id=$%d", len(values)+1))
		values = append(values, *HotelID)
	}
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	bookingIdRow, errGetId := ds.pool.Query(context.Background(), query)
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
