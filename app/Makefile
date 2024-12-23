.PHONY: build run test clean docker-build docker-run

# 构建应用
build:
	go build -o bin/blackapp ./cmd/server

# 运行应用
run:
	go run ./cmd/server

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
	docker build -t blackapp .

# 运行 Docker 容器
docker-run:
	docker-compose up

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