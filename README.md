# 研改 - 课程作业互评与修订平台


研改将作业答疑、匿名互评和作品修订放在同一流程中。学生可以针对含糊的评语追问，把反馈关联到修订位置；教师可以对照初稿、反馈与修订稿。首版覆盖一个班的一次作业。目前可启动 Web、API 和数据库环境，业务功能按下列增量交付。

## 快速启动

需要 Docker Desktop（含 Compose）。启动 Docker Desktop 后，在仓库根目录运行：

```powershell
if (!(Test-Path .env)) { Copy-Item .env.example .env }
docker compose up --build --wait
```

打开 [http://127.0.0.1:8080](http://127.0.0.1:8080)。Compose 会启动 Vue、Go API 和 PostgreSQL。

停止服务运行 `docker compose down`，数据库数据会保留。

## 技术选型

| 部分 | 技术 | 用途 |
| --- | --- | --- |
| 后端 | Go 1.25+、`net/http` | 单体 JSON API |
| 前端 | Vue 3、TypeScript、Vite、Node.js 24 | 作业、互评与版本对照界面 |
| 数据库 | PostgreSQL 17 | 保存作品版本、评阅任务和修订关联 |
| 部署 | Docker Compose、Nginx | 本机启动和局域网联调，统一 Web 与 API 入口 |
| 工程协作 | GitHub Actions、Issues、PR | 验证构建、记录需求与评审 |
| 外部集成 | 第三方邮件服务，待接入 | 评阅和补派提醒 |

[选型依据、部署结构与接入边界](docs/design/technology-stack.md)

## 项目功能

| 功能 | 首版设计 |
| --- | --- |
| 作业发布与提交 | 教师设置要求、评价维度和截止时间。学生提交作品并查看提交状态，截止时封存供互评使用的版本。 |
| 作业问答 | 学生在作业页提问，教师或助教可将共性问题标为统一答复，供班内同学查阅。 |
| 匿名互评 | 系统按有效提交分配任务，排除自评和重复分配。学生查看待评作品、完成状态及期限，并按评价维度给出反馈。 |
| 评语讨论 | 作者在具体评语下追问，对应评阅人匿名回复。相关学生和授权教学人员可访问讨论。 |
| 修订与对照 | 作者记录对每条建议的处理方式，并关联修改位置。教师可以对照初稿、反馈、追问和修订稿。 |
| 漏评与提醒 | 助教查看逾期任务并安排补评，系统向相关学生发送邮件待办提醒。 |

## 小组协作

| 岗位 | 职责 | 负责人 |
| --- | --- | --- |
| 项目经理 | 计划、跟踪、变更及风险管理 | 待组内确认 |
| 需求分析师 | 调研、需求规约与场景建模 | 待组内确认 |
| 架构／设计工程师 | 概要设计与详细设计 | 待组内确认 |
| 开发工程师（至少两人） | 编码、配置管理和单元测试 | 待组内确认 |
| 测试／质量工程师 | 测试设计、组织评审与缺陷跟踪 | 待组内确认 |
| AI 使用负责人（可兼任） | 记录 AI 使用、人工复核及最终报告 | 待组内确认 |

**开发与 Git 提交规范**集中在根目录的 [CONTRIBUTING.md](CONTRIBUTING.md)，涵盖分支命名、`type(scope): summary` 提交格式、Issue、PR、评审与提交前验证。

## 仓库目录

```text
.
├── .github/workflows/        # Go、前端和 Compose CI
├── cmd/server/              # Go 服务启动与运行配置
├── internal/httpapi/        # HTTP 路由与处理
├── web/                     # Vue 前端、Vite 配置和 npm 锁文件
│   └── src/                 # 页面组件和应用入口
├── deploy/                  # 前端 Dockerfile、Nginx 配置
├── migrations/              # PostgreSQL 迁移约定
├── docs/
│   ├── README.md            # 文档索引
│   ├── work-plans/          # 课程周计划、调研与 SRS 分工
│   ├── research-reports/    # 调研报告
│   ├── user-stories/        # 用户故事与验收标准
│   ├── requirements/        # SRS、用例与追溯矩阵
│   ├── models/              # 分析模型源文件
│   ├── design/              # 技术选型、架构和接口设计
│   ├── testing/             # 测试计划和验证记录
│   ├── delivery/            # 迭代回顾、答辩与交付资料
│   ├── development.md       # 开发环境与验证命令
│   └── deployment.md        # 本机部署、共享数据库与备份
├── .env.example             # Compose 本地配置示例
├── compose.yaml             # Web、API 与数据库服务编排
├── Dockerfile               # Go API 镜像
├── go.mod                   # Go 模块与版本
└── CONTRIBUTING.md           # 开发与 Git 协作规范
```
