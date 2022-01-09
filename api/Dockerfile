FROM golang:latest

WORKDIR /usr/api

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build

EXPOSE 9090

RUN go run .