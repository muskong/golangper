FROM golang:alpine AS builder

WORKDIR /app
COPY . .
RUN cd works/command-server && \
    go mod tidy && \
    cd /app && \
    CGO_ENABLED=0 GOOS=linux go build -o main works/command-server/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/works/command-server/configs ./configs
EXPOSE 8080
CMD ["./main"] 