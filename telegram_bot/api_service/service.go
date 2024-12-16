package api_service

import (
	"net/http"
	"os"
	"telegram_bot/database_service"

	"github.com/h4x4d/go_hsse_hotels/pkg/client"
)

type Service struct {
	hotelUrl   string
	bookingUrl string
	authUrl    string
	database_service.DatabaseService
	client.Client
	client http.Client
}

func NewService() (*Service, error) {
	service := new(Service)
	service.bookingUrl = "http://" + "booking" + ":" + os.Getenv("BOOKING_REST_PORT") + "/"
	service.hotelUrl = "http://" + "hotel" + ":" + os.Getenv("HOTEL_REST_PORT") + "/"
	service.authUrl = "http://" + "auth" + ":" + os.Getenv("AUTH_REST_PORT") + "/"

	database_pointer, errDatabase := database_service.NewDatabaseService()
	if errDatabase != nil {
		return nil, errDatabase
	}
	service.DatabaseService = *database_pointer
	service.client = http.Client{}

	tokenClient, errorClient := client.NewClient()
	if errorClient != nil {
		return nil, errorClient
	}
	service.Client = *tokenClient

	return service, nil
}
