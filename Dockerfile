FROM golang:1.23-alpine

# instal bash
RUN apk add --no-cache bash

# instal air
RUN go install github.com/air-verse/air@latest

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go mod tidy