FROM golang:alpine AS builder

WORKDIR /app
# 设置 Go 环境变量
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

COPY . .
RUN cd works/command-client && \
    go env && \
    go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn && \
    go env && \
    go mod tidy && \
    cd /app && \
    CGO_ENABLED=0 GOOS=linux go build -o main works/command-client/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/works/command-client/configs ./configs
EXPOSE 8081
CMD ["./main"] 