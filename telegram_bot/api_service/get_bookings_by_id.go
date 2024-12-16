package api_service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"telegram_bot/models"
)

func (s *Service) GetBookingByID(bookingID int64, user *models.User) (*models.Booking, error) {
	path := s.bookingUrl + "booking/" + strconv.FormatInt(bookingID, 10)
	request, errRequest := s.CreateRequest("GET", path, user)
	if errRequest != nil {
		return nil, errRequest
	}

	responseBooking, errBooking := s.client.Do(request)
	if errBooking != nil {
		return nil, errBooking
	}
	defer responseBooking.Body.Close()

	if responseBooking.StatusCode != http.StatusOK {
		return nil, errors.New("get booking by id failed")
	}

	bookingJSON, errJSON := ioutil.ReadAll(responseBooking.Body)
	if errJSON != nil {
		return nil, errJSON
	}

	booking := new(models.Booking)

	errDecode := json.Unmarshal(bookingJSON, &booking)
	if errDecode != nil {
		return nil, errDecode
	}
	return booking, nil
}
