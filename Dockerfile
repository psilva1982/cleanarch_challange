FROM golang:1.23.2-alpine AS builder

RUN apk add --no-cache git && \
    go install github.com/google/wire/cmd/wire@latest

WORKDIR /app

COPY . .

WORKDIR /app/cmd/ordersystem

RUN go mod download
RUN go build -o /app/ordersystem main.go wire_gen.go

FROM alpine:latest

COPY --from=builder /app/ordersystem /usr/local/bin/ordersystem
COPY --from=builder /app/cmd/ordersystem/.env /usr/local/bin/.env

WORKDIR /usr/local/bin

EXPOSE 8080

CMD ["ordersystem"]