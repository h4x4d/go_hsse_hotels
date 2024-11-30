FROM golang:1.23-alpine

WORKDIR /booking

COPY booking/go.mod .
COPY booking/go.sum .
RUN go mod download

COPY booking/ .

RUN go build -o booking cmd/hotels-booking-server/main.go

EXPOSE $BOOKING_REST_PORT

ENTRYPOINT ["./booking", "--port=8880", "--host=0.0.0.0"]

