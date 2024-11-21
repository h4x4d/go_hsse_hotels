FROM golang:1.23-alpine

WORKDIR /hotel
COPY . .

RUN go build -o hotel cmd/main.go

EXPOSE $HOTEL_REST_PORT
EXPOSE $HOTEL_GRPC_PORT

ENTRYPOINT ["sh", "scripts/wait_and_start.sh"]


