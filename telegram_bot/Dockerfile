FROM golang:1.23-alpine

WORKDIR /pkg
COPY pkg/go.mod .
COPY pkg/go.sum .
RUN go mod download

WORKDIR /telegram_bot
COPY telegram_bot/go.mod .
COPY telegram_bot/go.sum .
RUN go mod download

COPY telegram_bot/ .

COPY ../pkg ../pkg

RUN go build -o telegram_bot cmd/main.go

ENTRYPOINT ["./telegram_bot"]