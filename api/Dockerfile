# Dev stage, hot reloading with air
FROM golang:1.23-alpine AS dev

# Install git & curl
RUN apk add --no-cache git curl

# Install air for hot reloading
# RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Set up working directory
WORKDIR /app

# Copy go.mod and go.sum for dependency installation
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@latest

# Use a volume for hot reloading to avoid rebuilding the container
# VOLUME ["/app"]

# Default command for running the service with air
CMD ["air"]

# Build stage
FROM golang:1.23-alpine AS build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./cmd

# ev. add production image here
