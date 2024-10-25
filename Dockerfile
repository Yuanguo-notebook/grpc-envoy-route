
FROM golang:1.21
RUN apt-get update && apt-get install -y dnsutils

WORKDIR /app



COPY src ./src
COPY go.mod go.sum ./
RUN go mod download
