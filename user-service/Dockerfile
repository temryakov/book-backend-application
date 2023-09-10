FROM golang:1.20-alpine

RUN mkdir /user-service

ADD . /user-service

WORKDIR /user-service

RUN go build -o main cmd/main.go

CMD ["/user-service/main"]