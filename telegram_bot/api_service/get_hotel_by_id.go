package api_service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"telegram_bot/models"
)

func (s *Service) GetHotelByID(hotelID int64, user *models.User) (*models.Hotel, error) {
	path := s.hotelUrl + "hotel/" + strconv.Itoa(int(hotelID))
	request, errRequest := s.CreateRequest("GET", path, user)
	if errRequest != nil {
		return nil, errRequest
	}

	responseHotel, errHotel := s.client.Do(request)
	if errHotel != nil {
		return nil, errHotel
	}
	defer responseHotel.Body.Close()

	if responseHotel.StatusCode != http.StatusOK {
		return nil, errors.New("get hotels by id failed")
	}

	hotelJSON, errJSON := ioutil.ReadAll(responseHotel.Body)
	if errJSON != nil {
		return nil, errJSON
	}

	hotel := new(models.Hotel)

	errDecode := json.Unmarshal(hotelJSON, &hotel)
	if errDecode != nil {
		return nil, errDecode
	}
	return hotel, nil
}
