<h1 align="center"> Snippetbox </h1>

<p align="center">
  <img src="docs/img/logo.png" alt="Snippetbox Logo" width="40" />
</p>

> 本项目源于 Alex Edwards 的《Let's Go》一书，作者的博客与开源仓库：
> - 博客：https://www.alexedwards.net/blog
> - GitHub：https://github.com/alexedwards/

---

## 📋 项目简介

Snippetbox 是一个简洁高效的 Go 语言 Web 应用，允许用户粘贴和分享文本片段，类似于 Pastebin 或 GitHub Gist。支持用户注册、登录、代码片段增删查、密码修改等功能，适合学习、二次开发和实际部署。

---

## 🚀 项目特性

- 用户注册与登录
- 代码片段创建、浏览、管理
- 支持密码修改
- CSRF 防护与表单校验
- 响应式前端页面
- 支持 Docker 部署
- 单元测试与 CI/CD 支持

---

## 🗂️ 目录结构

```text
├── cmd/                # 应用入口与路由
│   └── web/
├── internal/           # 业务逻辑与数据模型
│   ├── models/
│   └── validator/
├── ui/                 # 前端模板与静态资源
│   ├── html/
│   └── static/
├── tls/                # TLS 证书
├── tmp/                # 临时文件
├── docs/               # 项目文档与图片
├── init.sql            # 数据库初始化脚本
├── docker-compose.yaml # Docker 编排
├── go.mod / go.sum     # Go 依赖管理
└── README.md           # 项目说明
```

---

## ⚡ 快速开始

### 1. 克隆项目
```bash
git clone https://github.com/你的用户名/snippet.git
cd snippet
```

### 2. 配置数据库

初始化数据库结构：
```bash
mysql -u root -p < init.sql
```

### 3. 运行项目（推荐方式：Linux 环境下使用 Makefile）

1. **初始化依赖**
   ```bash
   make init
   ```
   该命令会自动安装 Go 依赖并安装 air 热重载工具。

2. **启动项目**
   ```bash
   make run
   ```
   该命令会拉取/启动依赖服务（如数据库），并使用 air 启动项目，支持热重载。

> ⚠️ 本项目开发和测试环境为 Linux，其他操作系统请自行适配。

#### 其他方式（不推荐）

- 直接运行 Go 程序
  ```bash
  go run ./cmd/web
  ```
- 直接使用 Docker Compose
  ```bash
  docker-compose up --build
  ```

### 4. 环境变量说明

本项目依赖一系列环境变量（如数据库连接信息、服务端口等），请在启动前**自行在 Linux 环境下设置**。常见环境变量如下：

- `SERVER_PORT`：服务监听端口
- `MYSQL_USER`：数据库用户名
- `MYSQL_PASSWORD`：数据库密码
- `MYSQL_HOST`：数据库主机
- `MYSQL_DATABASE`：数据库名

你可以在本地创建 `.env` 文件，并参考 `cmd/web/main.go` 中的变量读取方式进行配置。例如：

```env
SERVER_PORT=4000
MYSQL_USER=root
MYSQL_PASSWORD=yourpassword
MYSQL_HOST=127.0.0.1:3306
MYSQL_DATABASE=snippetbox
```

**注意：请勿将敏感信息提交到版本库。**

### 4. 访问

浏览器访问： [http://localhost:4000](http://localhost:4000)

---

## 📝 贡献指南

欢迎提交 Issue 和 PR！请遵循 [Conventional Commits](https://www.conventionalcommits.org/zh-hans/v1.0.0/) 规范。

1. Fork 本仓库
2. 新建分支进行开发
3. 提交 PR 并描述变更内容

详细开发文档见 [docs/开发指南.md](docs/开发指南.md)

---

## 📚 更多文档

- [docs/部署说明.md](docs/部署说明.md)
- [docs/FAQ.md](docs/FAQ.md)

---


### TLS Generate
```shell
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```