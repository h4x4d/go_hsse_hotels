package client

import (
	"context"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/grpc/gen"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/grpc/utils"
	"github.com/h4x4d/go_hsse_hotels/booking/internal/models"
)

func GetHotelById(hotelId *int64) (*models.Hotel, error) {
	conn, err := utils.ConnectToHotel()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := gen.NewHotelClient(conn)
	hotelResp, err := client.GetHotel(context.Background(), &gen.HotelRequest{Id: *hotelId})
	if err != nil {
		return nil, err
	}
	hotel := models.Hotel{
		ID:         hotelResp.Id,
		Name:       &hotelResp.Name,
		City:       &hotelResp.City,
		Address:    &hotelResp.Address,
		Cost:       hotelResp.Cost,
		HotelClass: hotelResp.HotelClass,
		UserID:     hotelResp.UserId,
	}
	return &hotel, err
}
