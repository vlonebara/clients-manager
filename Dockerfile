FROM golang:1.26.1-alpine AS builder

WORKDIR /app

RUN apk add --no-cache build-base

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY internal/ ./internal/
RUN CGO_ENABLED=1 go build -o /app/clients-manager ./cmd/app

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/clients-manager ./clients-manager
RUN mkdir -p /app/db

EXPOSE 8080

CMD ["./clients-manager"]
