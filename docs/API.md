# AI学习社区 API 接口文档

> 本文档记录 AI 学习社区后端所有 RESTful API 接口。  
> **基础路径**：`http://localhost:8000`  
> **响应格式**：统一 JSON，通过 `code`、`message`、`data` 字段包装。

---

## 通用响应格式

### 成功响应

```json
{
  "code": 0,
  "message": "success",
  "data": { ... }
}
```

### 错误响应

```json
{
  "code": 1,
  "message": "错误描述",
  "data": null
}
```

---

## 公共错误码

| code | 说明 |
|------|------|
| 0 | 成功 |
| 1 | 请求参数错误 / 业务错误 |
| 500 | 服务器内部错误 |

---

## Post API

### 1. POST /api/post/create — 发布帖子

创建一篇新帖子。

**请求参数（JSON Body）**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| title | string | 是 | 帖子标题，长度 1-255 字符 |
| content | string | 是 | 帖子内容（支持 Markdown/富文本） |
| authorId | uint64 | 是 | 作者用户 ID |
| tags | string | 否 | 标签，逗号分隔，如 `AI,Python` |
| coverImage | string | 否 | 封面图片 URL |

**请求示例**

```json
{
  "title": "深度学习入门指南",
  "content": "# 深度学习\n\n深度学习是机器学习的一个分支...",
  "authorId": 1001,
  "tags": "AI,Python,深度学习",
  "coverImage": "https://example.com/cover.jpg"
}
```

**响应参数**

| 参数 | 类型 | 说明 |
|------|------|------|
| id | uint64 | 新建帖子的 ID |

**响应示例**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 42
  }
}
```

---

### 2. GET /api/post/detail — 帖子详情

根据帖子 ID 获取帖子详细信息。

**请求参数（Query）**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id | uint64 | 是 | 帖子 ID |

**请求示例**

```
GET /api/post/detail?id=42
```

**响应参数**

| 参数 | 类型 | 说明 |
|------|------|------|
| id | uint64 | 帖子 ID |
| title | string | 帖子标题 |
| content | string | 帖子内容 |
| authorId | uint64 | 作者用户 ID |
| status | int | 状态：0-已删除 1-正常 2-置顶 3-精选 |
| viewCount | uint | 浏览次数 |
| likeCount | uint | 点赞次数 |
| commentCount | uint | 评论次数 |
| tags | string | 标签 |
| coverImage | string | 封面图片 URL |
| isDeleted | int | 软删除标记 |
| createdAt | string | 创建时间 |
| updatedAt | string | 更新时间 |

**响应示例**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 42,
    "title": "深度学习入门指南",
    "content": "# 深度学习\n\n深度学习是机器学习的一个分支...",
    "authorId": 1001,
    "status": 1,
    "viewCount": 156,
    "likeCount": 23,
    "commentCount": 8,
    "tags": "AI,Python,深度学习",
    "coverImage": "https://example.com/cover.jpg",
    "isDeleted": 0,
    "createdAt": "2026-04-10 12:00:00",
    "updatedAt": "2026-04-10 12:00:00"
  }
}
```

---

### 3. GET /api/post/list — 帖子列表

分页查询帖子列表，支持多种过滤条件。

**请求参数（Query）**

| 参数 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| page | int | 否 | 1 | 页码 |
| pageSize | int | 否 | 10 | 每页数量 |
| authorId | uint64 | 否 | - | 按作者 ID 过滤 |
| tags | string | 否 | - | 按标签过滤 |
| status | int | 否 | - | 按状态过滤（0-3） |
| keyword | string | 否 | - | 关键词搜索 |

**请求示例**

```
GET /api/post/list?page=1&pageSize=10&keyword=深度学习
```

**响应参数**

| 参数 | 类型 | 说明 |
|------|------|------|
| total | uint64 | 总记录数 |
| page | uint | 当前页码 |
| pageSize | uint | 每页数量 |
| list | array | 帖子列表，每项同帖子详情字段 |

**响应示例**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total": 100,
    "page": 1,
    "pageSize": 10,
    "list": [
      {
        "id": 42,
        "title": "深度学习入门指南",
        "content": "...",
        "authorId": 1001,
        "status": 1,
        "viewCount": 156,
        "likeCount": 23,
        "commentCount": 8,
        "tags": "AI,Python",
        "coverImage": "",
        "isDeleted": 0,
        "createdAt": "2026-04-10 12:00:00",
        "updatedAt": "2026-04-10 12:00:00"
      }
    ]
  }
}
```

---

### 4. POST /api/post/update — 更新帖子

更新指定帖子的标题、内容、标签等信息。

**请求参数（JSON Body）**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id | uint64 | 是 | 帖子 ID |
| title | string | 是 | 帖子标题，长度 1-255 字符 |
| content | string | 是 | 帖子内容 |
| tags | string | 否 | 标签，逗号分隔 |
| coverImage | string | 否 | 封面图片 URL |

**请求示例**

```json
{
  "id": 42,
  "title": "深度学习入门指南（修订版）",
  "content": "# 深度学习\n\n修订后的内容...",
  "tags": "AI,Python,机器学习",
  "coverImage": "https://example.com/cover-v2.jpg"
}
```

**响应参数**

| 参数 | 类型 | 说明 |
|------|------|------|
| success | bool | 是否成功 |

**响应示例**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "success": true
  }
}
```

---

### 5. POST /api/post/delete — 删除帖子

软删除指定帖子。

**请求参数（JSON Body）**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id | uint64 | 是 | 帖子 ID |

**请求示例**

```json
{
  "id": 42
}
```

**响应参数**

| 参数 | 类型 | 说明 |
|------|------|------|
| success | bool | 是否成功 |

**响应示例**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "success": true
  }
}
```

---

## Comment API

### 6. POST /api/comment/create — 评论帖子

在指定帖子下发表一级评论。

**请求参数（JSON Body）**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| postId | uint64 | 是 | 被评论的帖子 ID |
| content | string | 是 | 评论内容 |
| authorId | uint64 | 是 | 评论作者 ID |
| parentId | uint64 | 否 | 父评论 ID（一级评论不传） |
| replyToUserId | uint64 | 否 | 回复目标用户 ID |

**请求示例**

```json
{
  "postId": 42,
  "content": "写得很好，受益匪浅！",
  "authorId": 1002,
  "replyToUserId": null,
  "parentId": null
}
```

**响应参数**

| 参数 | 类型 | 说明 |
|------|------|------|
| id | uint64 | 新建评论的 ID |

**响应示例**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 101
  }
}
```

---

### 7. POST /api/comment/reply — 回复评论

在指定父评论下发表二级回复。

**请求参数（JSON Body）**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| postId | uint64 | 是 | 所属帖子 ID |
| parentId | uint64 | 是 | 父评论 ID |
| content | string | 是 | 回复内容 |
| authorId | uint64 | 是 | 回复作者 ID |
| replyToUserId | uint64 | 否 | 被回复的用户 ID |

**请求示例**

```json
{
  "postId": 42,
  "parentId": 101,
  "content": "同意你的观点！",
  "authorId": 1003,
  "replyToUserId": 1002
}
```

**响应参数**

| 参数 | 类型 | 说明 |
|------|------|------|
| id | uint64 | 新建回复评论的 ID |

**响应示例**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 102
  }
}
```

---

### 8. POST /api/comment/delete — 删除评论

软删除指定评论（需校验作者权限）。

**请求参数（JSON Body）**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id | uint64 | 是 | 评论 ID |
| authorId | uint64 | 是 | 操作用户 ID（需与评论作者一致） |

**请求示例**

```json
{
  "id": 102,
  "authorId": 1003
}
```

**响应参数**

| 参数 | 类型 | 说明 |
|------|------|------|
| success | bool | 是否成功 |

**响应示例**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "success": true
  }
}
```

---

### 9. GET /api/comment/detail — 评论详情

根据评论 ID 获取单条评论的详细信息。

**请求参数（Query）**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id | uint64 | 是 | 评论 ID |

**请求示例**

```
GET /api/comment/detail?id=101
```

**响应参数**

| 参数 | 类型 | 说明 |
|------|------|------|
| id | uint64 | 评论 ID |
| postId | uint64 | 所属帖子 ID |
| userId | uint64 | 用户 ID |
| content | string | 评论内容 |
| parentId | uint64/null | 父评论 ID |
| authorId | uint64 | 评论作者 ID |
| authorName | string | 评论作者名称 |
| replyToUserId | uint64/null | 被回复用户 ID |
| replyToAuthor | string | 被回复用户名称 |
| depth | uint | 嵌套层级深度 |
| likeCount | uint | 点赞数 |
| status | int | 状态 |
| createdAt | string | 创建时间 |
| updatedAt | string | 更新时间 |
| children | array | 子评论列表（递归结构） |

**响应示例**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 101,
    "postId": 42,
    "userId": 1002,
    "content": "写得很好，受益匪浅！",
    "parentId": null,
    "authorId": 1002,
    "authorName": "张三",
    "replyToUserId": null,
    "replyToAuthor": "",
    "depth": 0,
    "likeCount": 5,
    "status": 1,
    "createdAt": "2026-04-10 14:00:00",
    "updatedAt": "2026-04-10 14:00:00",
    "children": [
      {
        "id": 102,
        "postId": 42,
        "userId": 1003,
        "content": "同意你的观点！",
        "parentId": 101,
        "authorId": 1003,
        "authorName": "李四",
        "replyToUserId": 1002,
        "replyToAuthor": "张三",
        "depth": 1,
        "likeCount": 2,
        "status": 1,
        "createdAt": "2026-04-10 14:30:00",
        "updatedAt": "2026-04-10 14:30:00",
        "children": []
      }
    ]
  }
}
```

---

### 10. GET /api/comment/tree — 评论树

获取指定帖子的完整评论树结构，支持分页。

**请求参数（Query）**

| 参数 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| postId | uint64 | 是 | - | 帖子 ID |
| page | int | 否 | 1 | 页码（一级评论分页） |
| pageSize | int | 否 | 20 | 每页数量 |

**请求示例**

```
GET /api/comment/tree?postId=42&page=1&pageSize=20
```

**响应参数**

| 参数 | 类型 | 说明 |
|------|------|------|
| total | int | 一级评论总数 |
| list | array | 扁平化评论列表 |
| tree | array | 层级树形结构列表 |

**响应示例**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total": 10,
    "list": [
      {
        "id": 101,
        "postId": 42,
        "userId": 1002,
        "content": "写得很好！",
        "parentId": null,
        "authorId": 1002,
        "authorName": "张三",
        "replyToUserId": null,
        "replyToAuthor": "",
        "depth": 0,
        "likeCount": 5,
        "status": 1,
        "createdAt": "2026-04-10 14:00:00",
        "updatedAt": "2026-04-10 14:00:00",
        "children": []
      },
      {
        "id": 102,
        "postId": 42,
        "userId": 1003,
        "content": "同意！",
        "parentId": 101,
        "authorId": 1003,
        "authorName": "李四",
        "replyToUserId": 1002,
        "replyToAuthor": "张三",
        "depth": 1,
        "likeCount": 2,
        "status": 1,
        "createdAt": "2026-04-10 14:30:00",
        "updatedAt": "2026-04-10 14:30:00",
        "children": []
      }
    ],
    "tree": [
      {
        "id": 101,
        "postId": 42,
        "userId": 1002,
        "content": "写得很好！",
        "parentId": null,
        "authorId": 1002,
        "authorName": "张三",
        "replyToUserId": null,
        "replyToAuthor": "",
        "depth": 0,
        "likeCount": 5,
        "status": 1,
        "createdAt": "2026-04-10 14:00:00",
        "updatedAt": "2026-04-10 14:00:00",
        "children": [
          {
            "id": 102,
            "postId": 42,
            "userId": 1003,
            "content": "同意！",
            "parentId": 101,
            "authorId": 1003,
            "authorName": "李四",
            "replyToUserId": 1002,
            "replyToAuthor": "张三",
            "depth": 1,
            "likeCount": 2,
            "status": 1,
            "createdAt": "2026-04-10 14:30:00",
            "updatedAt": "2026-04-10 14:30:00",
            "children": []
          }
        ]
      }
    ]
  }
}
```

---

## 接口总览

| # | 方法 | 路径 | 功能 |
|---|------|------|------|
| 1 | POST | /api/post/create | 发布帖子 |
| 2 | GET | /api/post/detail | 帖子详情 |
| 3 | GET | /api/post/list | 帖子列表 |
| 4 | POST | /api/post/update | 更新帖子 |
| 5 | POST | /api/post/delete | 删除帖子 |
| 6 | POST | /api/comment/create | 评论帖子 |
| 7 | POST | /api/comment/reply | 回复评论 |
| 8 | POST | /api/comment/delete | 删除评论 |
| 9 | GET | /api/comment/detail | 评论详情 |
| 10 | GET | /api/comment/tree | 评论树 |
