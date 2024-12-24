.PHONY: build run test clean docker-build docker-run

# 构建应用
build:
	go build -o bin/blackappClient ./cmd/client
	go build -o bin/blackappServer ./cmd/server

# 运行应用
run:
	go run ./cmd/server
	go run ./cmd/client

# 运行测试
test:
	go test -v ./...

# 生成 Swagger 文档
swagger:
	swag init -g cmd/server/main.go -o docs

# 清理构建文件
clean:
	rm -rf bin/

# 构建 Docker 镜像
docker-build:
	docker build -t blackappClient .
	docker build -t blackappServer .

# 运行 Docker 容器
docker-run:
	docker run -p 8080:8080 blackappServer
	docker run -p 8081:8081 blackappClient

# 停止 Docker 容器
docker-stop:
	docker-compose down

# 安装依赖
deps:
	go mod download

# 格式化代码
fmt:
	go fmt ./...

# 运行代码检查
lint:
	golangci-lint run 