FROM golang:1.20-alpine

RUN mkdir /review-service

ADD . /review-service

WORKDIR /review-service

RUN go install github.com/cosmtrek/air@latest

RUN go mod tidy
