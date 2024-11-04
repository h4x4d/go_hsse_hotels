FROM golang:1.23-alpine

WORKDIR /hotel
COPY . .

RUN go build -o hotel cmd/hotels-hotel-server/main.go

EXPOSE 8888

CMD ["./hotel", "--port=8888", "--host=0.0.0.0"]

