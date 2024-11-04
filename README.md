# MTS HSSE Go project. Booking system
--------
* Scheme of project
[https://github.com/Sinord/final-project/raw/master/components-diagram.png]

* Stack of technology
 - Docker
 - PostgreSQL
 - Golang
In project REST API is used to ensure user interaction with the system.

* API
 - Hotel information: ```/hotel/docks/swagger/hotels.yaml```
 - Booking: ```/booking/docks/swagger/booking.yaml```

* Get started

1) *start docker*
2) ```go run /booking/cmd/hotels-booking-server/main.go```
3) ```go run /hotel/cmd/hotels-hotel-server/main.go```