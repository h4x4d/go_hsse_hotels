FROM golang:1.23-alpine

WORKDIR /hotel

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o hotel cmd/main.go

EXPOSE $HOTEL_REST_PORT
EXPOSE $HOTEL_GRPC_PORT

ENTRYPOINT ["./hotel"]


