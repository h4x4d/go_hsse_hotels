package api_service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"telegram_bot/models"
)

func (s *Service) GetHotels(user *models.User) ([]models.Hotel, error) {
	request, errRequest := s.CreateRequest("GET", s.hotelUrl+"hotel/", user)
	if errRequest != nil {
		return nil, errRequest
	}

	responseHotels, errHotels := s.client.Do(request)
	if errHotels != nil {
		return nil, errHotels
	}
	defer responseHotels.Body.Close()

	if responseHotels.StatusCode != http.StatusOK {
		return nil, errors.New(responseHotels.Status)
	}

	hotelsJSON, errJSON := ioutil.ReadAll(responseHotels.Body)
	if errJSON != nil {
		return nil, errJSON
	}

	var hotels []models.Hotel

	errDecode := json.Unmarshal(hotelsJSON, &hotels)
	if errDecode != nil {
		return nil, errDecode
	}
	return hotels, nil
}
