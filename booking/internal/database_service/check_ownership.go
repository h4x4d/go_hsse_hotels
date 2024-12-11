package database_service

import (
	"context"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/grpc/client"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
	"go.opentelemetry.io/otel"
)

func (ds *DatabaseService) CheckOwnership(ctx context.Context, BookingID int64, user *models.User) (bool, error) {

	if user == nil {
		return false, nil
	}
	if user.Role == "customer" {
		booking, err := ds.GetByID(BookingID)
		if err != nil {
			return false, err
		}
		if booking == nil {
			return false, nil
		}
		return booking.UserID == user.UserID, nil
	}
	booking, err := ds.GetByID(BookingID)
	if err != nil {
		return false, err
	}

	tracer := otel.Tracer("Booking")
	ctx, span := tracer.Start(ctx, "check ownership db")
	defer span.End()

	hotel, hotelErr := client.GetHotelById(ctx, booking.HotelID)
	if hotelErr != nil {
		return false, hotelErr
	}
	return hotel.UserID == user.UserID, nil
}
