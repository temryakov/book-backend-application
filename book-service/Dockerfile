FROM golang:1.20-alpine

RUN mkdir /book-service

ADD . /book-service

WORKDIR /book-service

RUN go build -o main cmd/main.go

CMD ["/book-service/main"]