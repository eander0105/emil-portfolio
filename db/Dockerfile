FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@latest

EXPOSE 8090

CMD ["air", "-c", ".air.toml"]
