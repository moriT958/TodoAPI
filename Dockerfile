FROM golang:1.23.3 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o main /app/.

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/main .
USER 1001
CMD ["./main"]
