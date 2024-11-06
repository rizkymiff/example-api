FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o kirito ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/kirito .
EXPOSE 8080
CMD ["./kirito"]
