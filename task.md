# AI学习社区 - 帖子与评论核心API任务清单（敏捷开发模式）

## 项目概述
- **后端框架**：GoFrame
- **前端框架**：Vue3 + Vite
- **数据库**：MySQL
- **功能**：帖子与评论核心API + 前端页面
- **开发模式**：敏捷Sprint迭代，后端与前端并行开发

---

## Sprint 1: 基础设施搭建

### 1.1 数据库设计

| 任务名称 | 当前状态 | 修改文件 | 完成agentId | 完成时间 | 提交commit前8位 | 备注 |
|---------|---------|---------|------------|---------|----------------|------|
| 1.1.1 设计posts表结构 | ✅ 完成 | database/posts.sql | tuhome-backend-architect | 2026-04-10 20:04 | 0c38cd3 | 包含title, content, author_id, created_at, updated_at等字段 |
| 1.1.2 设计comments表结构 | ✅ 完成 | database/comments.sql | tuhome-backend-architect | 2026-04-10 20:05 | ba167ac | 支持树形层级结构，parent_id自关联 |
| 1.1.3 编写建表SQL脚本 | ✅ 完成 | database/init.sql | tuhome-backend-architect | 2026-04-10 20:07 | 4914263b | 生成完整的init.sql，整合posts和comments表 |
| 1.1.4 执行建表SQL脚本 | ✅ 完成 | - | tuhome-backend-architect | 2026-04-10 20:43 | 4914263b | 执行init.sql创建数据库表 |

### 1.2 后端项目初始化

| 任务名称 | 当前状态 | 修改文件 | 完成agentId | 完成时间 | 提交commit前8位 | 备注 |
|---------|---------|---------|------------|---------|----------------|------|
| 1.2.1 初始化GoFrame项目目录结构 | ✅ 完成 | model/, dao/, service/, controller/, internal/cmd/, internal/consts/, internal/packed/, manifest/config/, manifest/deploy/ | tuhome-backend-architect | 2026-04-10 20:09 | c573b6cd | model/dao/service/controller/internal/manifest目录结构 |
| 1.2.2 创建database.yaml配置文件 | ✅ 完成 | manifest/config/database.yaml | tuhome-backend-architect | 2026-04-10 20:10 | 11eb7e5 | 数据库连接配置 |
| 1.2.3 初始化数据库连接 | ✅ 完成 | internal/dao/dao.go, internal/cmd/cmd.go, manifest/config/config.yaml | tuhome-backend-architect | 2026-04-10 20:11 | 98c71cf7 | g.DB()初始化 |

### 1.3 前端项目初始化

| 任务名称 | 当前状态 | 修改文件 | 完成agentId | 完成时间 | 提交commit前8位 | 备注 |
|---------|---------|---------|------------|---------|----------------|------|
| 1.3.1 初始化Vue3 + Vite项目 | ✅ 完成 | frontend/ | tuhome-frontend-developer | 2026-04-10 20:17 | 2021ebc | 使用create-vite模板 |
| 1.3.2 安装依赖（axios, vue-router, pinia等） | ✅ 完成 | package.json, package-lock.json | tuhome-frontend-developer | 2026-04-10 20:18 | dab4bb9 | 核心依赖包 |
| 1.3.3 配置项目目录结构 | ✅ 完成 | - | tuhome-frontend-developer | 2026-04-10 20:19 | b65919f | views/components/api/store/router目录 |
| 1.3.4 配置API请求基地址 | ✅ 完成 | frontend/.env | tuhome-frontend-developer | 2026-04-10 20:22 | 9f69985 | axios baseURL配置，VITE_API_BASE_URL环境变量指向 http://localhost:8000 |

---

## Sprint 2: 帖子功能开发

### 2.1 后端-Post数据层（Model + Dao）

| 任务名称 | 当前状态 | 修改文件 | 完成agentId | 完成时间 | 提交commit前8位 | 备注 |
|---------|---------|---------|------------|---------|----------------|------|
| 2.1.1 创建Post模型定义 | ✅ 完成 | internal/model/entity/post.go | tuhome-backend-architect | 2026-04-10 20:22 | 20a9a5a | gmodel.Post实体 |
| 2.1.2 创建Post新增/删除/查询输入输出结构体 | ✅ 完成 | internal/model/do/post_do.go | tuhome-backend-architect | 2026-04-10 20:24 | 73bea33 | Do CreateReq/Do DeleteReq/Do GetOneReq等 |
| 2.1.3 创建Post数据访问方法-Insert | ✅ 完成 | internal/dao/post.go | tuhome-backend-architect | 2026-04-10 20:26 | fa0ab9e | dao.Post.Insert() |
| 2.1.4 创建Post数据访问方法-Delete | ✅ 完成 | internal/dao/post.go | tuhome-orchestrator:subagent:4f49609e | 2026-04-10 20:32 | fece73f | dao.Post.Delete() 软删除，返回受影响行数 |
| 2.1.5 创建Post数据访问方法-GetOne | ✅ 完成 | internal/dao/post.go | tuhome-backend-architect | 2026-04-10 20:34 | cd95d87 | dao.Post.GetOne() 支持软删除过滤 |
| 2.1.6 创建Post数据访问方法-GetList | ✅ 完成 | internal/dao/post.go | tuhome-orchestrator:subagent:e04dbd6a | 2026-04-10 20:34 | f6758db | dao.Post.GetList()分页查询，支持authorId/tags/status/keyword过滤，created_at DESC排序 |

### 2.2 后端-Post业务层与接口层

| 任务名称 | 当前状态 | 修改文件 | 完成agentId | 完成时间 | 提交commit前8位 | 备注 |
|---------|---------|---------|------------|---------|----------------|------|
| 2.2.1 实现Post业务方法-Create | ✅ 完成 | internal/service/post.go | tuhome-backend-architect | 2026-04-10 20:32 | 2b8ff65 | service.Post.Create() |
| 2.2.2 实现Post业务方法-Delete | ✅ 完成 | internal/service/post.go | tuhome-orchestrator:subagent:2328be6f | 2026-04-10 20:35 | 012c508 | service.Post.Delete() |
| 2.2.3 实现Post业务方法-GetDetail | ✅ 完成 | internal/service/post.go | tuhome-orchestrator:subagent:d2c94a75 | 2026-04-10 20:34 | dfad6ce | service.Post.GetDetail() |
| 2.2.4 实现Post业务方法-GetPageList | ✅ 完成 | internal/service/post.go | tuhome-backend-architect | 2026-04-10 20:38 | d8fce3e | service.Post.GetPageList() |
| 2.2.5 实现【发布帖子】API接口 | ✅ 完成 | internal/controller/post/post.go, internal/cmd/cmd.go, go.mod, main.go | tuhome-backend-architect | 2026-04-10 20:50 | e4749ae | POST /api/post/create，修复service.PostService调用和main.go |
| 2.2.6 实现【删除帖子】API接口 | ✅ 完成 | internal/controller/post/post.go | tuhome-backend-architect | 2026-04-10 20:47 | a0bab7c | POST /api/post/delete |
| 2.2.7 实现【查询帖子详情】API接口 | ✅ 完成 | internal/controller/post/post.go, internal/cmd/cmd.go | tuhome-orchestrator:subagent:4b5c79cf | 2026-04-10 20:45 | 4b37a67 | GET /api/post/detail |
| 2.2.8 实现【分页查询帖子列表】API接口 | ✅ 完成 | internal/controller/post/post.go | tuhome-orchestrator:subagent:b9212f47 | 2026-04-10 20:45 | 9ac9b46 | GET /api/post/list |
| 2.2.9 注册Post路由组 | ✅ 完成 | internal/cmd/cmd.go | tuhome-orchestrator:subagent:cd571790 | 2026-04-10 20:46 | 406b8a1 | 4个路由全部注册，POST create/delete, GET detail/list |
| 2.2.10 实现【更新帖子】API接口 | ✅ 完成 | internal/controller/post/post.go, internal/service/post.go, internal/dao/post.go | tuhome-backend-architect | 2026-04-10 21:04 | 4a89245 | POST /api/post/update，已完成DAO/Service/Controller三层实现 |

### 2.3 前端-Post API服务层

| 任务名称 | 当前状态 | 修改文件 | 完成agentId | 完成时间 | 提交commit前8位 | 备注 |
|---------|---------|---------|------------|---------|----------------|------|
| 2.3.1 封装Post API调用方法 | ✅ 完成 | frontend/src/api/post.js | tuhome-frontend-developer | 2026-04-10 20:53 | e8ef93a | api/post.js封装CRUD接口 |

### 2.4 前端-Post页面开发

| 任务名称 | 当前状态 | 修改文件 | 完成agentId | 完成时间 | 提交commit前8位 | 备注 |
|---------|---------|---------|------------|---------|----------------|------|
| 2.4.1 帖子列表页（分页+展示） | ✅ 完成 | frontend/src/views/post/PostList.vue | tuhome-frontend-developer | 2026-04-10 20:54 | 034cd08 | views/post/PostList.vue |
| 2.4.2 帖子详情页 | ✅ 完成 | frontend/src/views/post/PostDetail.vue | tuhome-frontend-developer | 2026-04-10 21:00 | 27eff74 | views/post/PostDetail.vue |
| 2.4.3 发布帖子页 | ✅ 完成 | frontend/src/views/post/PostCreate.vue | tuhome-frontend-developer | 2026-04-10 21:05 | ebc21e0 | views/post/PostCreate.vue |
| 2.4.4 编辑帖子页 | ✅ 完成 | frontend/src/views/post/PostEdit.vue, frontend/src/api/post.js | tuhome-frontend-developer | 2026-04-10 21:15 | aee87a8 | views/post/PostEdit.vue |

### 2.5 前端-Post路由配置

| 任务名称 | 当前状态 | 修改文件 | 完成agentId | 完成时间 | 提交commit前8位 | 备注 |
|---------|---------|---------|------------|---------|----------------|------|
| 2.5.1 配置帖子相关路由 | ✅ 完成 | frontend/src/router/index.js, frontend/src/main.js, frontend/src/App.vue | tuhome-frontend-developer | 2026-04-10 21:06 | 7f48779 | router/index.js配置 |

---

## Sprint 3: 评论功能开发

### 3.1 后端-Comment数据层（Model + Dao）

| 任务名称 | 当前状态 | 修改文件 | 完成agentId | 完成时间 | 提交commit前8位 | 备注 |
|---------|---------|---------|------------|---------|----------------|------|
| 3.1.1 创建Comment模型定义 | ✅ 完成 | internal/model/entity/comment.go | tuhome-backend-architect | 2026-04-10 21:06 | a98c3e3 | gmodel.Comment实体 |
| 3.1.2 创建Comment新增/删除/查询输入输出结构体 | ✅ 完成 | internal/model/do/comment_do.go | tuhome-backend-architect | 2026-04-10 21:07 | aa65bad | Do结构体定义 |
| 3.1.3 创建Comment数据访问方法-Insert | ✅ 完成 | internal/dao/comment.go | tuhome-orchestrator:subagent:6748d178 | 2026-04-10 21:16 | 7f16e74 | dao.Comment.Insert() |
| 3.1.4 创建Comment数据访问方法-Delete | ✅ 完成 | internal/dao/comment.go | tuhome-orchestrator:subagent:6748d178 | 2026-04-10 21:16 | 7f16e74 | dao.Comment.Delete() 软删除 |
| 3.1.5 创建Comment数据访问方法-GetOne | ✅ 完成 | internal/dao/comment.go | tuhome-orchestrator:subagent:6748d178 | 2026-04-10 21:16 | 7f16e74 | dao.Comment.GetOne() |
| 3.1.6 创建Comment数据访问方法-GetByPostId | ✅ 完成 | internal/dao/comment.go | tuhome-orchestrator:subagent:6748d178 | 2026-04-10 21:16 | 7f16e74 | dao.Comment.GetByPostId()一级评论 |
| 3.1.7 创建Comment数据访问方法-GetChildren | ✅ 完成 | internal/dao/comment.go | tuhome-orchestrator:subagent:6748d178 | 2026-04-10 21:16 | 7f16e74 | dao.Comment.GetChildren()子评论, GetAllChildrenByParentIds() |

### 3.2 后端-Comment业务层与接口层

| 任务名称 | 当前状态 | 修改文件 | 完成agentId | 完成时间 | 提交commit前8位 | 备注 |
|---------|---------|---------|------------|---------|----------------|------|
| 3.2.1 实现Comment业务方法-Create | ⬜ 待开始 | - | tuhome-backend-architect | - | - | service.Comment.Create() |
| 3.2.2 实现Comment业务方法-Reply | ⬜ 待开始 | - | tuhome-backend-architect | - | - | service.Comment.Reply()回复功能 |
| 3.2.3 实现Comment业务方法-Delete | ⬜ 待开始 | - | tuhome-backend-architect | - | - | service.Comment.Delete() |
| 3.2.4 实现Comment业务方法-GetDetail | ⬜ 待开始 | - | tuhome-backend-architect | - | - | service.Comment.GetDetail() |
| 3.2.5 实现Comment业务方法-GetTreeByPostId | ⬜ 待开始 | - | tuhome-backend-architect | - | - | service.Comment.GetTreeByPostId()层级结构 |
| 3.2.6 实现【评论帖子】API接口 | ⬜ 待开始 | - | tuhome-backend-architect | - | - | POST /api/comment/create |
| 3.2.7 实现【回复评论】API接口 | ⬜ 待开始 | - | tuhome-backend-architect | - | - | POST /api/comment/reply |
| 3.2.8 实现【删除评论】API接口 | ⬜ 待开始 | - | tuhome-backend-architect | - | - | POST /api/comment/delete |
| 3.2.9 实现【查询评论详情】API接口 | ⬜ 待开始 | - | tuhome-backend-architect | - | - | GET /api/comment/detail |
| 3.2.10 实现【分页查询帖子下评论（层级结构）】API接口 | ⬜ 待开始 | - | tuhome-backend-architect | - | - | GET /api/comment/tree |
| 3.2.11 注册Comment路由组 | ⬜ 待开始 | - | tuhome-backend-architect | - | - | 路由分组注册 |

### 3.3 后端-完善main.go入口

| 任务名称 | 当前状态 | 修改文件 | 完成agentId | 完成时间 | 提交commit前8位 | 备注 |
|---------|---------|---------|------------|---------|----------------|------|
| 3.3.1 完善main.go入口文件 | ⬜ 待开始 | - | tuhome-backend-architect | - | - | 整合所有路由和中间件 |

### 3.4 前端-Comment API服务层

| 任务名称 | 当前状态 | 修改文件 | 完成agentId | 完成时间 | 提交commit前8位 | 备注 |
|---------|---------|---------|------------|---------|----------------|------|
| 3.4.1 封装Comment API调用方法 | ⬜ 待开始 | - | tuhome-frontend-developer | - | - | api/comment.js封装CRUD接口 |
| 3.4.2 统一错误处理和响应拦截 | ⬜ 待开始 | - | tuhome-frontend-developer | - | - | axios拦截器统一处理 |

### 3.5 前端-Comment组件开发

| 任务名称 | 当前状态 | 修改文件 | 完成agentId | 完成时间 | 提交commit前8位 | 备注 |
|---------|---------|---------|------------|---------|----------------|------|
| 3.5.1 评论列表组件（层级树形结构） | ⬜ 待开始 | - | tuhome-frontend-developer | - | - | components/comment/CommentTree.vue |
| 3.5.2 一级评论项组件 | ⬜ 待开始 | - | tuhome-frontend-developer | - | - | components/comment/CommentItem.vue |
| 3.5.3 二级回复项组件 | ⬜ 待开始 | - | tuhome-frontend-developer | - | - | components/comment/ReplyItem.vue |
| 3.5.4 评论输入组件（支持回复） | ⬜ 待开始 | - | tuhome-frontend-developer | - | - | components/comment/CommentInput.vue |
| 3.5.5 评论分页组件 | ⬜ 待开始 | - | tuhome-frontend-developer | - | - | components/comment/CommentPagination.vue |

### 3.6 前端-Comment路由配置

| 任务名称 | 当前状态 | 修改文件 | 完成agentId | 完成时间 | 提交commit前8位 | 备注 |
|---------|---------|---------|------------|---------|----------------|------|
| 3.6.1 配置评论相关路由 | ⬜ 待开始 | - | tuhome-frontend-developer | - | - | 评论相关路由配置 |
| 3.6.2 路由守卫和权限控制 | ⬜ 待开始 | - | tuhome-frontend-developer | - | - | router guard配置 |

---

## Sprint 4: 集成测试与文档

### 4.1 前后端联调测试

| 任务名称 | 当前状态 | 修改文件 | 完成agentId | 完成时间 | 提交commit前8位 | 备注 |
|---------|---------|---------|------------|---------|----------------|------|
| 4.1.1 帖子功能前后端联调测试 | ⬜ 待开始 | - | tuhome-api-tester | - | - | Post CRUD全流程测试 |
| 4.1.2 评论功能前后端联调测试 | ⬜ 待开始 | - | tuhome-api-tester | - | - | Comment CRUD全流程测试 |
| 4.1.3 层级评论展示功能测试 | ⬜ 待开始 | - | tuhome-api-tester | - | - | 树形结构展示和分页测试 |

### 4.2 接口文档

| 任务名称 | 当前状态 | 修改文件 | 完成agentId | 完成时间 | 提交commit前8位 | 备注 |
|---------|---------|---------|------------|---------|----------------|------|
| 4.2.1 编写API接口文档 | ⬜ 待开始 | - | tuhome-backend-architect | - | - | 汇总所有API接口说明 |

---

## 📊 Sprint 开发节奏总结

| Sprint | 阶段 | 后端任务 | 前端任务 | 产出 |
|--------|------|---------|---------|------|
| Sprint 1 | 基础设施 | DB设计 + 项目初始化 | 项目初始化 | 基础设施就绪 |
| Sprint 2 | 帖子功能 | Model→Dao→Service→Controller→路由 | API封装 → 页面开发 → 路由 | 帖子功能完整 |
| Sprint 3 | 评论功能 | Model→Dao→Service→Controller→路由 | API封装 → 组件开发 → 路由 | 评论功能完整 |
| Sprint 4 | 集成收尾 | 联调测试 + 文档 | 联调测试 | 交付完成 |

---

**审核后请回复"确认"，我立即进入开发阶段。**
