.PHONY: help init run

help:
	@echo "可用指令："
	@echo "  make help      查看所有可用指令"
	@echo "  make init      安装 Go 依赖和 air 热重载工具"
	@echo "  make run       拉取镜像并用 air 启动项目"

init:
	@echo "安装 Go 依赖..."
	go mod tidy
	@echo "安装 air..."
	go install github.com/cosmtrek/air@latest

run:
	@echo "拉取/启动依赖服务..."
	docker-compose up -d
	@echo "使用 air 启动项目..."
	air