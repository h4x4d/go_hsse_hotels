package api_service

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"telegram_bot/models"
)

func (s *Service) CreateHotel(hotel *models.Hotel, user *models.User) (bool, error) {
	if hotel == nil {
		return false, errors.New("hotel is nil")
	}

	hotelJson, errorEncode := json.Marshal(hotel)
	if errorEncode != nil {
		return false, errorEncode
	}

	request, errRequest := s.CreateRequest("POST", s.hotelUrl+"hotel/", user)
	if errRequest != nil {
		return false, errRequest
	}
	request.Body = io.NopCloser(bytes.NewBuffer(hotelJson))

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
