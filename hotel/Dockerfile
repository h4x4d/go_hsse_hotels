FROM golang:1.23-alpine

WORKDIR /hotel
COPY . .

RUN go build -o hotel cmd/hotels-hotel-server/main.go

EXPOSE 8888

ENTRYPOINT ["sh", "wait_and_start.sh"]


