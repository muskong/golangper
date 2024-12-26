FROM golang:alpine AS builder

WORKDIR /app
COPY . .
RUN ls -la
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./command-client/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/command-client/configs ./configs
EXPOSE 8081
CMD ["./main"] 