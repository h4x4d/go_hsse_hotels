package api_service

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"telegram_bot/models"
)

func (s *Service) CreateBooking(booking *models.Booking, user *models.User) (bool, error) {
	if booking == nil {
		return false, errors.New("booking is nil")
	}

	bookingJson, errorEncode := json.Marshal(booking)
	if errorEncode != nil {
		return false, errorEncode
	}

	request, errRequest := s.CreateRequest("POST", s.bookingUrl+"booking/", user)
	if errRequest != nil {
		return false, errRequest
	}

	request.Body = io.NopCloser(bytes.NewBuffer(bookingJson))

	responseCreate, errCreate := s.client.Do(request)
	if errCreate != nil {
		return false, errCreate
	}
	defer responseCreate.Body.Close()

	if responseCreate.StatusCode != http.StatusOK {
		return false, nil
	}
	return true, nil
}
