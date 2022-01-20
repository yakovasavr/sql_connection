# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /sql_app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./models models/
COPY ./store store/
COPY ./usecase usecase/
COPY ./webserver webserver/
COPY main.go ./


RUN go build -o /sql_connection ./

EXPOSE 8000
CMD [ "/sql_connection" ]