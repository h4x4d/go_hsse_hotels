FROM golang:1.23-alpine

WORKDIR /notification

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go mod download

RUN go build -o notification cmd/main.go

ENTRYPOINT ["./notification"]


