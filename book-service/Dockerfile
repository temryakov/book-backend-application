FROM golang:1.20-alpine

RUN mkdir /book-service

ADD . /book-service

WORKDIR /book-service

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod tidy
