FROM golang:1.23-alpine

WORKDIR /pkg
COPY pkg/go.mod .
COPY pkg/go.sum .
RUN go mod download

WORKDIR /hotel
COPY hotel/go.mod .
COPY hotel/go.sum .
RUN go mod download

COPY hotel/ .

COPY ../pkg ../pkg

WORKDIR /hotel
RUN go build -o hotel cmd/main.go

EXPOSE $HOTEL_REST_PORT
EXPOSE $HOTEL_GRPC_PORT

ENTRYPOINT ["./hotel"]


