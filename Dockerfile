FROM golang:1.22.5

WORKDIR /app

COPY main.go go.* /app/
