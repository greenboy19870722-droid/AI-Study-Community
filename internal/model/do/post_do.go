package do

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PostCreateReq 创建帖子请求
type PostCreateReq struct {
	Title      string `json:"title"      dc:"帖子标题" v:"required|length:1,255"`
	Content    string `json:"content"    dc:"帖子内容(Markdown/富文本)" v:"required"`
	AuthorId   uint64 `json:"authorId"   dc:"作者用户ID" v:"required"`
	Tags       string `json:"tags"        dc:"标签,多个逗号分隔"`
	CoverImage string `json:"coverImage"  dc:"封面图片URL"`
}

// PostDeleteReq 删除帖子请求
type PostDeleteReq struct {
	Id uint64 `json:"id" dc:"帖子ID" v:"required"`
}

// PostUpdateReq 更新帖子请求
type PostUpdateReq struct {
	Id          uint64 `json:"id"          dc:"帖子ID" v:"required"`
	Title       string `json:"title"       dc:"帖子标题" v:"required|length:1,255"`
	Content     string `json:"content"     dc:"帖子内容(Markdown/富文本)" v:"required"`
	Tags        string `json:"tags"        dc:"标签,多个逗号分隔"`
	CoverImage  string `json:"coverImage"  dc:"封面图片URL"`
}

// PostGetOneReq 查询单个帖子请求
type PostGetOneReq struct {
	Id uint64 `json:"id" dc:"帖子ID" v:"required"`
}

// PostGetListReq 分页查询帖子列表请求
type PostGetListReq struct {
	Page       int    `json:"page"        dc:"页码" d:"1"`
	PageSize   int    `json:"pageSize"    dc:"每页数量" d:"10"`
	AuthorId   uint64 `json:"authorId"    dc:"作者用户ID(可选)"`
	Tags       string `json:"tags"        dc:"标签筛选(可选)"`
	Status     *int   `json:"status"      dc:"状态筛选(可选)" v:"in:0,1,2,3"`
	Keyword    string `json:"keyword"     dc:"关键词搜索(可选)"`
}

// PostGetPageListReq 分页查询帖子列表请求(别名,用于Service层)
type PostGetPageListReq struct {
	Page       uint   `json:"page"        dc:"页码" d:"1"`
	PageSize   uint   `json:"pageSize"    dc:"每页数量" d:"10"`
	AuthorId   uint64 `json:"authorId"    dc:"作者用户ID(可选)"`
	Tags       string `json:"tags"        dc:"标签筛选(可选)"`
	Status     *int   `json:"status"      dc:"状态筛选(可选)" v:"in:0,1,2,3"`
	Keyword    string `json:"keyword"     dc:"关键词搜索(可选)"`
}

// PostResp 帖子响应
type PostResp struct {
	Id           uint64      `json:"id"           dc:"帖子ID"`
	Title        string      `json:"title"        dc:"帖子标题"`
	Content      string      `json:"content"      dc:"帖子内容"
	AuthorId     uint64      `json:"authorId"     dc:"作者用户ID"
	Status       int         `json:"status"       dc:"状态: 0-删除 1-正常 2-置顶 3-精华"
	ViewCount    uint        `json:"viewCount"    dc:"浏览量"`
	LikeCount    uint        `json:"likeCount"    dc:"点赞数"
	CommentCount uint        `json:"commentCount" dc:"评论数"
	Tags         string      `json:"tags"         dc:"标签"
	CoverImage   string      `json:"coverImage"   dc:"封面图片URL"
	IsDeleted    int         `json:"isDeleted"    dc:"软删除状态"
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    dc:"更新时间"`
}

// PostListResp 帖子列表响应
type PostListResp struct {
	Total    uint64     `json:"total"    dc:"总数"`
	Page     uint       `json:"page"     dc:"当前页"`
	PageSize uint       `json:"pageSize" dc:"每页数量"`
	List     []*PostResp `json:"list"    dc:"帖子列表"`
}

// PostCreateResp 创建帖子响应
type PostCreateResp struct {
	Id uint64 `json:"id" dc:"新创建的帖子ID"`
}

// PostDeleteResp 删除帖子响应
type PostDeleteResp struct {
	Success bool `json:"success" dc:"是否成功"`
}

// PostUpdateResp 更新帖子响应
type PostUpdateResp struct {
	Success bool `json:"success" dc:"是否成功"`
}
