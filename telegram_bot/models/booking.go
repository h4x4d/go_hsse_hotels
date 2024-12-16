package models

type Booking struct {
	// booking id
	BookingID int64 `json:"booking_id,omitempty"`

	// date from
	// Example: 31-12-2024
	// Required: true
	// Pattern: ^\d{2}-\d{2}-\d{4}$
	DateFrom *string `json:"date_from"`

	// date to
	// Example: 31-12-2025
	// Required: true
	// Pattern: ^\d{2}-\d{2}-\d{4}$
	DateTo *string `json:"date_to"`

	// full cost
	FullCost int64 `json:"full_cost,omitempty"`

	// hotel id
	// Required: true
	HotelID *int64 `json:"hotel_id"`

	// status of booking
	// Enum: ["Waiting","Confirmed","Canceled"]
	Status string `json:"status,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`
}

func NewBooking() *Booking {
	booking := new(Booking)
	booking.DateTo = new(string)
	booking.DateFrom = new(string)
	booking.HotelID = new(int64)
	return booking
}
