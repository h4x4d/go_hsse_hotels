package handlers

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	models2 "github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/utils"
	"google.golang.org/protobuf/proto"
	"log/slog"
	"net/http"
)

func (handler *Handler) GetHotelsHandler(params hotel.GetHotelsParams) (responder middleware.Responder) {
	// Tracing
	_, span := handler.tracer.Start(context.Background(), "get hotels")
	defer span.End()

	defer utils.CatchPanic(&responder)

	hotelClassLog := int64(0)
	if params.HotelClass != nil {
		hotelClassLog = *params.HotelClass
	}
	hotelNameLog := proto.String("empty")
	if params.HotelClass != nil {
		hotelNameLog = params.Name
	}
	hotelCityLog := proto.String("empty")
	if params.HotelClass != nil {
		hotelCityLog = params.City
	}

	payload, err := handler.Database.GetAll(params.City, params.HotelClass, params.Name)
	if err != nil {
		// Logging
		slog.Error(
			"failed get hotels",
			slog.Group("hotel-filter",
				slog.Int64("hotel-class", hotelClassLog),
				slog.String("name", *hotelNameLog),
				slog.String("city", *hotelCityLog),
			),
			slog.Int("status_code", http.StatusInternalServerError),
			slog.String("error", err.Error()),
		)
		return utils.HandleInternalError(err)
	}

	if len(payload) == 0 {
		// Logging
		slog.Info(
			"failed get hotels",
			slog.Group("hotel-filter",
				slog.Int64("hotel-class", hotelClassLog),
				slog.String("name", *hotelNameLog),
				slog.String("city", *hotelCityLog),
			),
			slog.Int("status_code", http.StatusNotFound),
			slog.String("error", "Suitable hotels not found"),
		)

		notFound := int64(http.StatusNotFound)
		return new(hotel.GetHotelsNotFound).WithPayload(&models2.Error{
			ErrorMessage:    "Suitable hotels not found",
			ErrorStatusCode: &notFound,
		})
	}
	// Logging
	slog.Info(
		"get hotels",
		slog.Group("hotel-filter",
			slog.Int64("hotel-class", hotelClassLog),
			slog.String("name", *hotelNameLog),
			slog.String("city", *hotelCityLog),
		),
		slog.Int("status_code", hotel.GetHotelsOKCode),
	)

	result := new(hotel.GetHotelsOK)
	result = result.WithPayload(payload)
	return result
}
