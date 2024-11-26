FROM golang:1.23-alpine

WORKDIR /auth

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o auth cmd/main.go

EXPOSE $AUTH_REST_PORT

ENTRYPOINT ["./auth", "--port=8800", "--host=0.0.0.0"]

