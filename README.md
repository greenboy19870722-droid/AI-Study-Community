# AI学习社区

一个基于 Vue3 + GoFrame 的帖子与评论核心功能的全栈社区项目，支持树形层级评论结构。

## 项目简介

**AI学习社区** 是一个面向 AI/机器学习学习者的内容交流平台，核心功能包括：

- **帖子管理**：发布、编辑、删除、查询帖子，支持标签、封面图、状态管理
- **层级评论**：支持一级评论和二级回复，构成树形结构，支持分页查询
- **用户认证**：基于 JWT Token 的登录态管理（开发中）

### 技术栈

| 层级 | 技术 |
|------|------|
| 前端框架 | Vue 3.5 + Vite 8 + Vue Router 5 |
| UI 组件库 | Element Plus 2.13 |
| 状态管理 | Pinia 3.0 |
| HTTP 客户端 | Axios 1.15 |
| 后端框架 | GoFrame 2.10 (Go 1.23) |
| 数据库 | MySQL 8.0 (utf8mb4) |
| ORM | GoFrame Database/ORM 内置驱动 |

### 项目结构

```
AI-Study-Community/
├── database/              # 数据库建表 SQL
│   ├── init.sql           # 完整初始化脚本
│   ├── posts.sql          # 帖子表结构
│   └── comments.sql       # 评论表结构
├── docs/                  # 项目文档
│   └── API.md             # API 接口文档
├── frontend/              # Vue3 前端项目
│   ├── src/
│   │   ├── api/           # API 封装层 (post.js, comment.js, user.js, request.js)
│   │   ├── components/    # 公共组件 (comment/)
│   │   ├── views/         # 页面视图 (post/, Login.vue)
│   │   ├── router/        # 路由配置
│   │   ├── utils/         # 工具函数 (auth.js)
│   │   ├── main.js
│   │   └── App.vue
│   ├── .env               # 前端环境变量 (VITE_API_BASE_URL)
│   └── package.json
├── internal/               # GoFrame 后端源码
│   ├── cmd/               # 入口与路由注册
│   ├── consts/            # 常量定义
│   ├── controller/        # 控制器层 (post/, comment/, user/)
│   ├── dao/               # 数据访问层 (post.go, comment.go, dao.go, user.go)
│   ├── handler/           # HTTP Handler
│   ├── model/             # 模型定义
│   │   ├── entity/        # 数据库实体 (post.go, comment.go, user.go)
│   │   └── do/            # 数据对象 (输入输出结构体)
│   ├── service/           # 业务逻辑层 (post.go, comment.go, user.go)
│   └── packed/            # 资源打包
├── manifest/              # 部署配置
│   ├── config/
│   │   ├── config.yaml    # 主配置
│   │   └── database.yaml  # 数据库配置
│   └── docker/
│       └── Dockerfile     # Docker 镜像构建
├── main.go                # 后端入口文件
└── go.mod                 # Go 依赖
```

---

## 环境配置与运行

### 前置依赖

- **Go 1.21+** (项目使用 Go 1.23)
- **MySQL 8.0+**
- **Node.js 18+** (前端构建)
- **Git**

### 1. 数据库初始化

```sql
-- 登录 MySQL，执行以下命令创建数据库和表
CREATE DATABASE IF NOT EXISTS `ai_study_community` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `ai_study_community`;
-- 然后执行 database/init.sql 中的建表语句
```

或直接使用初始化脚本：

```bash
mysql -u root -p < database/init.sql
```

### 2. 后端配置

编辑 `manifest/config/database.yaml`，修改数据库连接信息：

```yaml
database:
  type: "mysql"
  host: "127.0.0.1"
  port: 3306
  name: "ai_study_community"
  user: "root"
  pass: "your_password"        # 修改为你的 MySQL 密码
  charset: "utf8mb4"
  timezone: "+8:00"
  maxOpenConns: 100
  maxIdleConns: 10
```

### 3. 启动后端

```bash
# 安装 Go 依赖
go mod tidy

# 编译并运行
go run main.go

# 或编译为可执行文件
go build -o main.exe main.go && ./main.exe
```

后端服务默认运行在 `http://localhost:8000`。

### 4. 启动前端

```bash
cd frontend

# 安装依赖
npm install

# 开发模式启动
npm run dev

# 生产构建
npm run build
```

前端开发服务器默认运行在 `http://localhost:5173`。

### 5. Docker 部署（可选）

```bash
# 构建后端 Docker 镜像
cd manifest/docker
docker build -t ai-study-community-backend .

# 运行容器
docker run -d -p 8000:8000 \
  -e DATABASE_HOST=host.docker.internal \
  -e DATABASE_PORT=3306 \
  -e DATABASE_NAME=ai_study_community \
  -e DATABASE_USER=root \
  -e DATABASE_PASS=your_password \
  ai-study-community-backend
```

---

## 核心 API 请求示例

> **基础路径**：`http://localhost:8000`
> **统一响应格式**：`{"code": 0, "message": "success", "data": {...}}`

### 帖子 API

**发布帖子**

```bash
curl -X POST http://localhost:8000/api/post/create \
  -H "Content-Type: application/json" \
  -d '{
    "title": "深度学习入门指南",
    "content": "# 深度学习\n\n深度学习是机器学习的一个分支...",
    "authorId": 1001,
    "tags": "AI,Python,深度学习",
    "coverImage": "https://example.com/cover.jpg"
  }'
```

**查询帖子详情**

```bash
curl "http://localhost:8000/api/post/detail?id=1"
```

**分页查询帖子列表**

```bash
curl "http://localhost:8000/api/post/list?page=1&pageSize=10&keyword=深度学习"
```

**更新帖子**

```bash
curl -X POST http://localhost:8000/api/post/update \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1,
    "title": "深度学习入门指南（修订版）",
    "content": "# 深度学习\n\n修订后的内容...",
    "tags": "AI,Python,机器学习"
  }'
```

**删除帖子（软删除）**

```bash
curl -X POST http://localhost:8000/api/post/delete \
  -H "Content-Type: application/json" \
  -d '{"id": 1}'
```

### 评论 API

**评论帖子（一级评论）**

```bash
curl -X POST http://localhost:8000/api/comment/create \
  -H "Content-Type: application/json" \
  -d '{
    "postId": 1,
    "content": "写得很好，受益匪浅！",
    "authorId": 1002
  }'
```

**回复评论（二级回复）**

```bash
curl -X POST http://localhost:8000/api/comment/reply \
  -H "Content-Type: application/json" \
  -d '{
    "postId": 1,
    "parentId": 101,
    "content": "同意你的观点！",
    "authorId": 1003,
    "replyToUserId": 1002
  }'
```

**查询评论详情**

```bash
curl "http://localhost:8000/api/comment/detail?id=101"
```

**查询帖子评论树（层级结构）**

```bash
curl "http://localhost:8000/api/comment/tree?postId=1&page=1&pageSize=20"
```

**删除评论**

```bash
curl -X POST http://localhost:8000/api/comment/delete \
  -H "Content-Type: application/json" \
  -d '{"id": 101, "authorId": 1002}'
```

---

## 设计文档

### 技术选型

#### 为什么选择 GoFrame？

1. **全栈 GUI 工程化能力**：GoFrame 提供从 API 定义、参数校验、数据库操作到响应处理的全套工程化工具，契合"帖子+评论"这类 CRUD 密集型业务。
2. **内置 ORM**：GoFrame 的 `gdb` 模块对 MySQL 支持成熟，分页、软删除、事务等操作一行代码搞定，开发效率高。
3. **Controller-Service-DAO 分层清晰**：项目结构与 GoFrame 倡导的经典三层架构天然契合，便于多人协作和后续维护。
4. **活跃的社区与文档**：GoFrame 文档详细，对复杂查询（如评论树形结构）的支持有据可查。

#### 数据库表设计

**posts 表**：

- `author_id` 建立索引，支持按作者查帖子的高频场景。
- `status` 字段（0删除/1正常/2置顶/3精华）支持内容运营需求，无需改表结构即可实现精选/置顶。
- `is_deleted` 软删除标记保留数据完整性，`deleted_at` 记录删除时间。
- `view_count`、`like_count`、`comment_count` 反向计数冗余存储，避免联表查询。

**comments 表**：

- `parent_id` 自关联支持无限层级评论，`depth` 字段标记嵌套深度便于前端渲染缩进。
- `fk_comments_post` 外键确保评论所属帖子存在时才能创建，级联删除避免孤儿评论。
- `fk_comments_parent` 外键设置 `ON DELETE SET NULL`，删除父评论时不级联删除子评论（子评论保留但变为"一级评论"），符合多数社区产品的产品逻辑。

#### API 设计

- **RESTful 风格**：`POST /create`、`GET /detail`、`GET /list`、`POST /update`、`POST /delete`，语义清晰。
- **统一响应包装**：`{"code": 0, "message": "success", "data": {...}}`，前端统一拦截器处理错误提示。
- **评论树单独接口**：`GET /comment/tree` 返回 `list`（扁平列表）和 `tree`（嵌套树形）两份数据，前端可按需取用。

---

### 核心实现

#### 评论层级查询的实现

**数据库层面的设计**：

- 每条评论记录 `parent_id` 指向父评论，顶级评论 `parent_id = NULL`。
- `depth` 字段在插入时由服务层计算（父评论 depth + 1），避免递归查询时重复计算。

**服务层 tree 构建逻辑**（`internal/service/comment.go:122`）：

1. **分页获取一级评论**：先按 `parent_id IS NULL` 分页查出顶级评论。
2. **全量获取子评论**：为了在内存中构建树，先用大页容量（10000）查出所有子评论。
3. **内存中组装树**：
   - 第一次遍历：建立 `commentMap[id -> CommentResp]` 和 `childrenMap[parentId -> []CommentResp]`。
   - 第二次遍历：将子评论分配给父评论的 `children` 字段。
   - 最后：从 `childrenMap` 中提取一级评论作为树根返回。

**权衡**：此方案在子评论数量可控（单帖评论 < 10000）时非常高效，避免了数据库递归 CTE 或多次 JOIN。但当帖子评论数极多时，全量加载到内存会产生瓶颈，后文会讨论改进方向。

#### Service-DAO 分层

- **DAO 层**（`internal/dao/post.go`、`comment.go`）：只做数据访问，封装 `Insert`、`Delete`、`GetOne`、`GetList` 等原子操作，不含业务逻辑。
- **Service 层**（`internal/service/post.go`、`comment.go`）：编排业务逻辑，处理参数校验、权限校验（如评论删除需校验 `authorId`）、数据转换。
- **Controller 层**（`internal/controller/post/post.go`）：接收 HTTP 请求参数，调用 Service 层，返回统一格式响应。

---

### 权衡与展望

#### 当前实现的局限性

1. **评论 tree 内存全量加载**：当单帖评论数达到数万条时，`getAllChildrenByPostId` 全量加载到内存会有 OOM 风险。`pageSize: 10000` 是一个经验值折中，不是真正的无限层级支持。
2. **缺少 Redis 缓存**：帖子详情、评论 tree 均无缓存，高并发场景下数据库压力较大。
3. **没有消息队列**：评论异步计数（如 `comment_count++`）目前是同步更新，缺少最终一致性保障。
4. **评论嵌套层级无数据库层面限制**：虽然有 `depth` 字段，但没有在应用层或 DB 触发器限制最大嵌套深度，理论上可以无限嵌套。
5. **用户系统不完整**：`user.go` 的 DAO/Service/Controller 已建立，但 `authorName` 等字段在评论响应中仍为 TODO 状态。

#### 支撑百万级用户的架构调整

如果目标扩展到百万级用户，当前架构需要在以下层面进行调整：

| 方向 | 当前现状 | 扩展方案 |
|------|----------|----------|
| **数据库分片** | 单库单表 | 按 `post_id % N` 或 `user_id % N` 水平分片，或引入 ShardingSphere 中间件 |
| **评论查询** | 全量加载到内存构建树 | 改为递归 CTE 或应用层懒加载（点击展开子评论时才加载），避免全量内存占用 |
| **缓存层** | 无 | 引入 Redis：帖子详情缓存 5 分钟，评论 tree 缓存按 `post_id` 粒度，热点帖子使用 CDN |
| **计数异步化** | 同步更新 `comment_count` | 引入 Kafka/RabbitMQ，评论事件写入队列，消费者异步更新计数字段，保证最终一致性 |
| **读写分离** | 单库读写混合 | 主从复制，读流量分摊到只读从库，评论列表等读密集型操作走从库 |
| **弹性扩展** | 单体部署 | 后端部署为无状态服务，使用 K8s HPA 根据 CPU/内存自动扩缩容；前端静态资源上 CDN |
| **搜索** | 无 | 引入 Elasticsearch，支持帖子标题/内容的全文搜索和标签过滤 |
| **用户系统** | 基本骨架 | 接入 JWT + refresh token，迁移到独立的用户服务（微服务拆分） |

---

## 接口总览

| 方法 | 路径 | 功能 |
|------|------|------|
| POST | /api/post/create | 发布帖子 |
| GET | /api/post/detail | 帖子详情 |
| GET | /api/post/list | 帖子列表（分页） |
| POST | /api/post/update | 更新帖子 |
| POST | /api/post/delete | 删除帖子（软删除） |
| POST | /api/comment/create | 评论帖子 |
| POST | /api/comment/reply | 回复评论 |
| POST | /api/comment/delete | 删除评论（软删除） |
| GET | /api/comment/detail | 评论详情 |
| GET | /api/comment/tree | 帖子评论树（层级结构） |
