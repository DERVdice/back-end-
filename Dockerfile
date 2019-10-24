FROM golang:latest

RUN mkdir -p /go/src/back-end-test
WORKDIR /go/src/back-end-test
COPY . /go/src/back-end-test

EXPOSE 8080

