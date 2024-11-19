package handlers

import (
	"context"
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/utils"
	"github.com/jackc/pgx/v5"
	"net/http"
	"os"
)

func CreateHotelHandler(params hotel.CreateHotelParams, principal interface{}) (responder middleware.Responder) {
	defer utils.CatchPanic(&responder)

	// connecting to database hotel
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "db", os.Getenv("POSTGRES_PORT"), "hotel")
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, err.Error())
	}
	defer conn.Close(context.Background())

	// executing insertion of hotel
	_, errInsertHotel := conn.Exec(context.Background(),
		"INSERT INTO hotels (id, name, city, address, hotel_class) VALUES ($1, $2, $3, $4, $5)",
		params.Object.ID, params.Object.Name, params.Object.City, params.Object.Address, params.Object.HotelClass)
	if errInsertHotel != nil {
		return middleware.Error(http.StatusInternalServerError, errInsertHotel.Error())
	}

	result := new(hotel.CreateHotelOK)
	return result
}
