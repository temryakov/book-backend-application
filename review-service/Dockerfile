FROM golang:1.20-alpine

RUN mkdir /review-service

ADD . /review-service

WORKDIR /review-service

RUN go build -o main cmd/main.go

CMD ["/review-service/main"]