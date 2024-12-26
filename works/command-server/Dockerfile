FROM golang:alpine AS builder

WORKDIR /app
# 设置 Go 环境变量
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

RUN go env
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env

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