FROM golang:alpine AS builder

WORKDIR /app
COPY . .
RUN cd works/command-client
RUN pwd
RUN ls -la
RUN go mod tidy
RUN cd /app
RUN CGO_ENABLED=0 GOOS=linux go build -o main works/command-client/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/command-client/configs ./configs
EXPOSE 8081
CMD ["./main"] 