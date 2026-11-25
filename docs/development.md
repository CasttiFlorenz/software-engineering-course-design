# 开发环境

以下命令以 Windows PowerShell 为例。在包含 `compose.yaml`、`go.mod` 和 `web/` 的仓库根目录执行，步骤另有说明时除外。

| 任务 | 运行方式 | 需要安装 |
| --- | --- | --- |
| 首次验收、联调或演示 | A · Compose 启动全部服务 | Git、Docker Desktop（含 Compose） |
| 修改前端或 Go 代码 | B · 本机 Go + Vite | Git、Go 1.25+、Node.js 24（含 npm）；运行数据库时另需 Docker |

## 1. 获取代码

首次克隆时执行：

```powershell
git clone https://github.com/CasttiFlorenz/software-engineering-course-design.git
Set-Location software-engineering-course-design
```

已经克隆的组员进入原仓库目录即可。开始新任务前，先处理手头未提交的改动，再同步 `main`。开工作分支的步骤见 [开发与 Git 协作规范](../CONTRIBUTING.md)。

## 2A. Compose 启动完整环境

1. 启动 Docker Desktop，等 Linux 引擎就绪。在 PowerShell 执行 `docker version`，确认有 **Server** 输出，再执行 `docker compose version`。
2. 在仓库根目录创建本机配置并启动：

   ```powershell
   if (!(Test-Path .env)) { Copy-Item .env.example .env }
   docker compose config --quiet
   docker compose up --build --wait
   docker compose ps
   ```

   首次启动会下载镜像并构建前后端。默认配置用于单机开发，联调环境请在首次初始化数据库前设置独立的 `POSTGRES_PASSWORD`。保留已有的 `.env`，提交时只跟踪 `.env.example`。
3. 打开 [http://127.0.0.1:8080](http://127.0.0.1:8080)，首页应显示「服务在线」。也可以分别检查 API 和数据库：

   ```powershell
   (Invoke-RestMethod http://127.0.0.1:8080/healthz).status
   docker compose exec -T db psql -U yangai -d yangai -Atc "SELECT 1;"
   ```

   预期分别输出 `ok` 和 `1`。如在 `.env` 修改了 `POSTGRES_USER` 或 `POSTGRES_DB`，同步替换查询中的用户名或库名。健康接口验证 Web 到 Go 的代理链路，SQL 验证 PostgreSQL 容器。应用的持久化连接随业务用例接入。
4. 结束时运行 `docker compose down`，数据库卷会保留。修改 Go 或 Vue 后，重新运行启动命令即可构建新镜像。

## 2B. 本机开发 Go 与 Vue

先运行 `go version`、`node --version` 和 `npm --version` 确认工具可用。Compose 的 Web 服务与本机 Go 默认都使用 8080，切换开发方式前先停止占用该端口的环境。

**终端 1 · 从仓库根目录启动 Go**

```powershell
go run ./cmd/server
```

Go 默认监听 `127.0.0.1:8080`。修改 Go 后按 Ctrl+C，再执行启动命令。

**终端 2 · 从仓库根目录启动 Vite**

```powershell
Set-Location web
npm ci
npm run dev
```

打开 Vite 打印的地址，通常为 `http://localhost:5173`。修改 Vue 会自动刷新，Vite 将 `/healthz` 转发到 `127.0.0.1:8080`。依赖锁文件变化时重新运行 `npm ci`。结束时在两个终端分别按 Ctrl+C。

需要单独启动数据库时，从根目录创建 `.env` 后运行 `docker compose up -d --wait db`，再用 2A 的 SQL 命令验证。当前 Compose 数据库没有映射宿主机端口，查询通过 `docker compose exec` 进入容器执行。团队共享环境与数据库备份见 [本机部署与数据库共享](deployment.md)。

## 3. 提交前验证

在仓库根目录逐条运行，前一条失败时先处理错误。

```powershell
go test ./...
go vet ./...
go build ./...
npm --prefix web ci
npm --prefix web run build
```

调整 Dockerfile、反向代理或 Compose 时，再执行 2A 的启动和验证步骤。GitHub Actions 还会检查首页、API 代理及 PostgreSQL 查询。PR 里记录实际运行的检查和结果。

## 常见问题

| 现象 | 处理 |
| --- | --- |
| `docker` 找不到或没有 Server 输出 | 安装完成后重新打开 PowerShell，启动 Docker Desktop，等引擎就绪后再运行 `docker version`。 |
| Compose 报 `POSTGRES_PASSWORD` 缺失 | 确认在仓库根目录执行，且 `.env` 存在、变量有值。 |
| Docker 镜像下载失败 | 确认 Docker Desktop 能访问 Docker Hub。可先运行 `docker compose pull db` 验证拉取，再重试构建。 |
| Compose 显示 8080 已占用 | 停止本机 Go 或旧 Compose 服务。也可以在 `.env` 设 `WEB_PORT=8081`，重新启动并访问 8081。 |
| Vite 首页显示「服务不可用」 | 确认终端 1 的 Go 进程仍在，并直接检查 `http://127.0.0.1:8080/healthz`。 |
| 容器没有进入健康状态 | 运行 `docker compose ps`，再用 `docker compose logs --tail=80 web api db` 查看对应服务。 |
