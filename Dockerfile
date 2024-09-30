FROM docker.io/library/golang:1.23.1-bookworm

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o /golang-rest-api

CMD [ "/golang-rest-api" ]
