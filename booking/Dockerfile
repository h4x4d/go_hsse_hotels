FROM golang:1.23-alpine

WORKDIR /booking

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o booking cmd/hotels-booking-server/main.go

EXPOSE $BOOKING_REST_PORT

ENTRYPOINT ["sh", "scripts/wait_and_start.sh"]

