# MTS HSSE Go project. Booking system
--------
## Scheme of project

![scheme](https://github.com/Sinord/final-project/raw/master/components-diagram.png)

## Stack of technology
  - Docker
  - PostgreSQL
  - Golang

In project REST API is used to ensure user interaction with the system.

## API
  - [Hotel information](/hotel/api/swagger/hotels.yaml)
  - [Booking](/booking/api/swagger/booking.yaml)

## Get started

1) Ensure that Docker is installed
2) Fill .env file like in .env-example
3) Run `docker compose up -d --build` or just `make build` if makefile is installed in your system