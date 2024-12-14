FROM golang:1.23-alpine

WORKDIR /pkg
COPY pkg/go.mod .
COPY pkg/go.sum .
RUN go mod download

WORKDIR /booking

COPY booking/go.mod .
COPY booking/go.sum .
RUN go mod download

COPY booking/ .
COPY ../pkg ../pkg

RUN go build -o booking cmd/hotels-booking-server/main.go

EXPOSE $BOOKING_REST_PORT

ENTRYPOINT ["./booking", "--port=8880", "--host=0.0.0.0"]

