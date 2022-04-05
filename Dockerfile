# syntax=docker/dockerfile:1

FROM golang:1.16-alpine
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY body.json ./

RUN go mod download

COPY *.go ./

RUN go build -o /docker-api-app

CMD [ "/docker-api-app" ]