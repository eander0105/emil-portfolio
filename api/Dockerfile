FROM golang:1.22-alpine as BASE

# Install necessary tools and dependencies
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh curl
# RUN apk add --no-cache git curl bash

# Install air for hot reloading
# RUN curl -fLo /binRUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# /air https://raw.githubusercontent.com/cosmtrek/air/master/bin/linux/air && chmod +x /bin/air

# Set up working directory
WORKDIR /app

# Copy go.mod and go.sum for dependency installation
COPY go.mod go.sum ./

# Download dependencies
# RUN go mod download

# Copy the rest of the application code
COPY . .

# Use a volume for hot reloading to avoid rebuilding the container
VOLUME ["/app"]

# Default command for running the service with air
CMD ["air"]
