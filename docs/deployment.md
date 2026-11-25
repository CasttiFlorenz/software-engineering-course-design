# 本机运行与团队数据库

## Docker Compose 本机运行

Docker Compose 在一台组员电脑上运行 Web、Go API 和 PostgreSQL。Web 接收浏览器请求，并把 API 请求代理给 Go 服务。当前页面通过健康检查验证 Web 与 API 联通，PostgreSQL 已配置独立数据卷，业务连接和数据迁移随首个持久化用例接入。

安装支持 Compose 的 Docker，复制配置并启动。

```powershell
Copy-Item .env.example .env
docker compose up --build -d
docker compose ps
```

打开 `http://127.0.0.1:8080`。查看日志使用 `docker compose logs -f`，停止使用 `docker compose down`。默认停止后保留数据库卷，`down -v` 会删除数据卷。

## 组员如何使用

| 场景 | 数据库安排 | 组员怎么访问 |
| --- | --- | --- |
| 各自开发 | 每人运行自己的 PostgreSQL | 同步 Git 中的迁移和脱敏演示数据，开发环境互不影响 |
| 同网联调或答辩 | 一台组员电脑运行整套服务并保存数据 | 主机将 `WEB_BIND` 改为 `0.0.0.0`，仅在可信局域网允许 Web 端口入站。组员通过 `http://主机局域网IP:8080` 打开同一个应用 |
| 跨网联调 | 一台组员电脑或团队可用的集成环境保存一份数据 | 通过团队私有网络访问 Web；托管 PostgreSQL 可作为需要持续在线时的选项 |

共享的是同一个应用实例及其背后的数据。Compose 内服务通过 `db:5432` 找到 PostgreSQL，数据库端口保持在容器网络内。需要检查数据时，主机管理员可通过 `docker compose exec db psql -U yangai -d yangai` 进入数据库，用户名与库名以实际配置为准。主机关机、睡眠或断网时，共享环境也会中断。

常用安排为平时各自开发，需要联调或答辩时在一台机器启动共享实例。共享环境设置独立密码，更新前确认迁移和演示数据兼容。需要保留共享数据时用 `pg_dump` 备份到 Git 仓库以外的位置。

参考 [Docker Compose 网络](https://docs.docker.com/compose/how-tos/networking)。
