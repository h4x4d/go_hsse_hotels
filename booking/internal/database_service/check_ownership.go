package database_service

import (
	"github.com/h4x4d/go_hsse_hotels/booking/internal/grpc/client"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
)

func (ds *DatabaseService) CheckOwnership(BookingID int64, user *models.User) (bool, error) {
	if user == nil {
		return false, nil
	}
	if user.Role == "customer" {
		booking, err := ds.GetByID(BookingID)
		if err != nil {
			return false, err
		}
		return booking.UserID == user.UserID, nil
	}
	booking, err := ds.GetByID(BookingID)
	if err != nil {
		return false, err
	}
	hotel, hotelErr := client.GetHotelById(booking.HotelID)
	if hotelErr != nil {
		return false, hotelErr
	}
	return hotel.UserID == user.UserID, nil
}
