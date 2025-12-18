FROM golang:1.25.4 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o service ./cmd/app

FROM golang:1.25.4
WORKDIR /app
COPY --from=builder /app/service .

EXPOSE 3000
CMD ["./service"]
