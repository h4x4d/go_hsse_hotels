package api_service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"telegram_bot/models"
)

func (s *Service) GetBookings(hotelID int64, user *models.User) ([]models.Booking, error) {
	path := s.bookingUrl + "booking/"

	urlObject, errUrl := url.Parse(path)
	if errUrl != nil {
		return nil, errUrl
	}

	params := url.Values{}
	params.Add("hotel_id", strconv.FormatInt(hotelID, 10))

	urlObject.RawQuery = params.Encode()

	request, errRequest := s.CreateRequest("GET", urlObject.String(), user)
	if errRequest != nil {
		return nil, errRequest
	}

	responseBookings, errBookings := s.client.Do(request)
	if errBookings != nil {
		return nil, errBookings
	}
	defer responseBookings.Body.Close()

	if responseBookings.StatusCode != http.StatusOK {
		return nil, errors.New("get bookings by hotel id failed")
	}

	bookingsJSON, errJSON := ioutil.ReadAll(responseBookings.Body)
	if errJSON != nil {
		return nil, errJSON
	}

	var bookings []models.Booking

	errDecode := json.Unmarshal(bookingsJSON, &bookings)
	if errDecode != nil {
		return nil, errDecode
	}
	return bookings, nil
}
