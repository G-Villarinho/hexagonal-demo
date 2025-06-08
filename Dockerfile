FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/http

FROM alpine:latest

WORKDIR /app

RUN apk update && apk upgrade --no-cache

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]