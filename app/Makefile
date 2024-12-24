.PHONY: build run test clean docker-build docker-run

# 构建应用
build:
	go build -o bin/blackClient ./cmd/client
	go build -o bin/blackServer ./cmd/server

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
	docker build -t blackClient .
	docker build -t blackServer .

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

# systemctl 生成 blackServer.service
systemd:
	@echo "[Unit]" > blackServer.service
	@echo "Description=Black Server Service" >> blackServer.service
	@echo "After=network.target postgresql.service redis.service" >> blackServer.service
	@echo "" >> blackServer.service
	@echo "[Service]" >> blackServer.service
	@echo "Type=simple" >> blackServer.service
	@echo "User=www-data" >> blackServer.service
	@echo "Group=www-data" >> blackServer.service
	@echo "WorkingDirectory=/sites/blacklist/app" >> blackServer.service
	@echo "ExecStart=/sites/blacklist/app/bin/blackServer" >> blackServer.service
	@echo "Restart=always" >> blackServer.service
	@echo "RestartSec=3" >> blackServer.service
	@echo "" >> blackServer.service
	@echo "[Install]" >> blackServer.service
	@echo "WantedBy=multi-user.target" >> blackServer.service
	@sudo mv blackServer.service /etc/systemd/system/
	@sudo systemctl daemon-reload
	@echo "Black Server service has been installed. Use 'sudo systemctl start blackServer' to start it." 

	@echo "[Unit]" > blackClient.service
	@echo "Description=Black Client Service" >> blackClient.service
	@echo "After=network.target postgresql.service redis.service" >> blackClient.service
	@echo "" >> blackClient.service
	@echo "[Service]" >> blackClient.service
	@echo "Type=simple" >> blackClient.service
	@echo "User=www-data" >> blackClient.service
	@echo "Group=www-data" >> blackClient.service
	@echo "WorkingDirectory=/sites/blacklist/app" >> blackClient.service
	@echo "ExecStart=/sites/blacklist/app/bin/blackClient" >> blackClient.service
	@echo "Restart=always" >> blackClient.service
	@echo "RestartSec=3" >> blackClient.service
	@echo "" >> blackClient.service
	@echo "[Install]" >> blackClient.service
	@echo "WantedBy=multi-user.target" >> blackClient.service
	@sudo mv blackClient.service /etc/systemd/system/
	@sudo systemctl daemon-reload
	@echo "Black Client service has been installed. Use 'sudo systemctl start blackClient' to start it." 