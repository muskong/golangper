FROM golang:alpine AS builder

WORKDIR /app
# 设置 Go 环境变量
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

COPY . .
RUN cd works/command-client && \
    go mod tidy && \
    cd /app && \
    go build -o main works/command-client/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/works/command-client/configs ./configs
EXPOSE 8081
CMD ["./main"] 