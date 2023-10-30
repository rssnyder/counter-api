FROM golang:1.21-alpine
LABEL org.opencontainers.image.source https://github.com/rssnyder/counter-api

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /counter-api

ENTRYPOINT /counter-api