FROM golang:alpine AS builder

WORKDIR /app
COPY .. .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./command-server/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/command-server/configs ./configs
EXPOSE 8080
CMD ["./main"] 