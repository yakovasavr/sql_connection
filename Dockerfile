# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /sql_app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd /cmd
COPY ./configs /configs
COPY ./internal /internal

RUN go build -o /test_echo
RUN go build -o /sql_connection /cmd/webserver

EXPOSE 8000
CMD [ "/test_echo" ]