FROM golang:1.23-alpine

WORKDIR /booking
COPY . .

RUN go build -o booking cmd/hotels-booking-server/main.go

EXPOSE 8880

ENTRYPOINT ["./wait_and_start.sh"]

