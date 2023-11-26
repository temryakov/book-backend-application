FROM golang:1.20-alpine

RUN mkdir /user-service

ADD . /user-service

WORKDIR /user-service

RUN go install github.com/cosmtrek/air@latest

RUN go mod tidy
