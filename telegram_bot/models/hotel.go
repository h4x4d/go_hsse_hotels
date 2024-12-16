package models

type Hotel struct {

	// address
	// Example: Red Square №1
	// Required: true
	Address *string `json:"address"`

	// city
	// Example: Moscow
	// Required: true
	City *string `json:"city"`

	// cost
	Cost int64 `json:"cost,omitempty"`

	// number of stars of hotel
	// Enum: [0,1,2,3,4,5]
	HotelClass int64 `json:"hotel_class,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	// Example: Radisson
	// Required: true
	Name *string `json:"name"`

	// user id
	UserID string `json:"user_id,omitempty"`
}

func NewHotel() *Hotel {
	hotel := new(Hotel)
	hotel.Address = new(string)
	hotel.City = new(string)
	hotel.Name = new(string)
	return hotel
}
