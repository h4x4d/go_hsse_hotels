package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/h4x4d/go_hsse_hotels/hotel/internal/models"
	hotel2 "github.com/h4x4d/go_hsse_hotels/hotel/internal/restapi/operations/hotel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"testing"
)

// WARNING: ALL TESTS REQUIRE RUNNING DOCKER COMPOSE PACKAGE (ONLY ON TEST DB, IT WILL BE DESTROYED)

func setup() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestMain(m *testing.M) {
	setup()

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("HOTEL_DB_NAME"))
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	_, delErr := pool.Query(context.Background(), "delete from hotels;")
	if delErr != nil {
		log.Fatal(err)
	}
	code := m.Run()

	_, delErr = pool.Query(context.Background(), "delete from hotels;")
	if delErr != nil {
		log.Fatal(err)
	}
	os.Exit(code)
}

type RegisterBody struct {
	Email      string `json:"email"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	TelegramID int64  `json:"telegram_id"`
}

func GetToken(t *testing.T, role string) string {
	authUrl := "http://127.0.0.1:" + os.Getenv("AUTH_REST_PORT")

	body := RegisterBody{
		Email:      fmt.Sprintf("testEmail%s@yandex.ru", role),
		Login:      fmt.Sprintf("testLogin%s", role),
		Password:   "testPassword",
		Role:       role,
		TelegramID: 123456,
	}
	data, marshalErr := json.Marshal(body)
	assert.Nil(t, marshalErr, "Error marshalling body")

	request, err := http.NewRequest("POST", authUrl+"/register", bytes.NewBuffer(data))
	assert.Nil(t, err, "Error creating request")
	request.Header.Set("Content-Type", "application/json")

	resp, httpErr := http.DefaultClient.Do(request)
	assert.Nil(t, httpErr, "Error on response")

	if resp.StatusCode == 409 {
		request, err := http.NewRequest("POST", authUrl+"/login", bytes.NewBuffer(data))
		assert.Nil(t, err, "Error creating request")
		request.Header.Set("Content-Type", "application/json")

		resp, httpErr = http.DefaultClient.Do(request)
		assert.Nil(t, httpErr, "Error on response")
	}
	assert.Equal(t, 200, resp.StatusCode, "Error registering")

	respData := struct {
		Token string `json:"token"`
	}{}
	decodeErr := json.NewDecoder(resp.Body).Decode(&respData)
	assert.Nil(t, decodeErr, "Error decoding id")
	return respData.Token
}

func GetHotels(t *testing.T, hotelUrl string) *[]models.Hotel {
	request, err := http.NewRequest("GET", hotelUrl+"/hotel", nil)
	assert.Nil(t, err, "Error creating request")
	request.Header.Set("Content-Type", "application/json")

	resp, httpErr := http.DefaultClient.Do(request)
	assert.Nil(t, httpErr, "Error on response")

	assert.Equal(t, 200, resp.StatusCode, "Error getting hotels")

	var oldHotels []models.Hotel
	decodeErr := json.NewDecoder(resp.Body).Decode(&oldHotels)
	assert.Nil(t, decodeErr, "Error decoding hotels")

	return &oldHotels
}

func GetHotelById(t *testing.T, hotelUrl string, hotelID int64) *models.Hotel {
	request, err := http.NewRequest("GET", hotelUrl+"/hotel/"+strconv.Itoa(int(hotelID)), nil)
	assert.Nil(t, err, "Error creating request")
	request.Header.Set("Content-Type", "application/json")

	resp, httpErr := http.DefaultClient.Do(request)
	assert.Nil(t, httpErr, "Error on response")

	assert.Equal(t, 200, resp.StatusCode, "Error getting hotel")

	hotel := models.Hotel{}
	decodeErr := json.NewDecoder(resp.Body).Decode(&hotel)
	assert.Nil(t, decodeErr, "Error decoding hotels")

	return &hotel
}

func CreateHotel(t *testing.T, hotel models.Hotel, hotelUrl string, token string) int64 {
	data, marshalErr := json.Marshal(hotel)
	assert.Nil(t, marshalErr, "Error marshalling hotel")

	request, err := http.NewRequest("POST", hotelUrl+"/hotel", bytes.NewReader(data))
	assert.Nil(t, err, "Error creating request")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("api_key", token)

	resp, httpErr := http.DefaultClient.Do(request)
	assert.Nil(t, httpErr, "Error on response")

	assert.Equal(t, 200, resp.StatusCode, "Error creating hotel")

	respData := hotel2.CreateHotelOKBody{}
	decodeErr := json.NewDecoder(resp.Body).Decode(&respData)
	assert.Nil(t, decodeErr, "Error decoding id")

	return respData.ID
}

func UpdateHotel(t *testing.T, hotelId int64, hotel models.Hotel, hotelUrl string, token string) models.Hotel {
	data, marshalErr := json.Marshal(hotel)
	assert.Nil(t, marshalErr, "Error marshalling hotel")

	request, err := http.NewRequest("PUT", hotelUrl+"/hotel/"+strconv.Itoa(int(hotelId)), bytes.NewReader(data))
	assert.Nil(t, err, "Error creating request")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("api_key", token)

	resp, httpErr := http.DefaultClient.Do(request)
	assert.Nil(t, httpErr, "Error on response")

	assert.Equal(t, 200, resp.StatusCode, "Error creating hotel")

	respData := models.Hotel{}
	decodeErr := json.NewDecoder(resp.Body).Decode(&respData)
	assert.Nil(t, decodeErr, "Error decoding id")

	return respData
}

func CmpHotels(hotel1 models.Hotel, hotel2 models.Hotel) bool {
	return !(*hotel1.Name != *hotel2.Name ||
		*hotel1.City != *hotel2.City ||
		*hotel1.Address != *hotel2.Address ||
		hotel1.HotelClass != hotel2.HotelClass ||
		hotel1.Cost != hotel2.Cost)
}

func TestHotelInteraction(t *testing.T) {
	hotelUrl := "http://127.0.0.1:" + os.Getenv("HOTEL_REST_PORT")
	token := GetToken(t, "hotelier")

	name := "Name" + strconv.Itoa(rand.Int())
	city := "City" + strconv.Itoa(rand.Int())
	address := "City, Street, 2" + strconv.Itoa(rand.Int())
	hotel := models.Hotel{
		Name:       &name,
		City:       &city,
		Cost:       1000,
		HotelClass: 3,
		Address:    &address,
	}
	CreateHotel(t, hotel, hotelUrl, token)

	newHotels := GetHotels(t, hotelUrl)

	assert.Equal(t, len(*newHotels), 1)

	assert.Equal(t, CmpHotels(hotel, (*newHotels)[0]), true)
	id := (*newHotels)[0].ID

	byId := GetHotelById(t, hotelUrl, id)
	assert.Equal(t, CmpHotels(hotel, *byId), true)

	name = "Name" + strconv.Itoa(rand.Int())
	city = "City" + strconv.Itoa(rand.Int())
	address = "City, Street, 2" + strconv.Itoa(rand.Int())
	updatedHotel := models.Hotel{
		Name:       &name,
		City:       &city,
		Cost:       1000,
		HotelClass: 3,
		Address:    &address,
	}

	got := UpdateHotel(t, id, updatedHotel, hotelUrl, token)
	assert.Equal(t, CmpHotels(updatedHotel, got), true)

	byId = GetHotelById(t, hotelUrl, id)
	assert.Equal(t, CmpHotels(updatedHotel, *byId), true)
}
