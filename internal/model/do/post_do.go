package do

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PostCreateReq creates post request
type PostCreateReq struct {
	Title      string `json:"title"      dc:"Post title" v:"required|length:1,255"`
	Content    string `json:"content"    dc:"Post content (Markdown/Rich text)" v:"required"`
	AuthorId   uint64 `json:"authorId"   dc:"Author user ID" v:"required"`
	Tags       string `json:"tags"       dc:"Tags, comma separated"`
	CoverImage string `json:"coverImage" dc:"Cover image URL"`
}

// PostDeleteReq delete post request
type PostDeleteReq struct {
	Id uint64 `json:"id" dc:"Post ID" v:"required"`
}

// PostUpdateReq update post request
type PostUpdateReq struct {
	Id          uint64 `json:"id"          dc:"Post ID" v:"required"`
	Title       string `json:"title"       dc:"Post title" v:"required|length:1,255"`
	Content     string `json:"content"     dc:"Post content" v:"required"`
	Tags        string `json:"tags"        dc:"Tags, comma separated"`
	CoverImage  string `json:"coverImage"  dc:"Cover image URL"`
}

// PostGetOneReq get one post request
type PostGetOneReq struct {
	Id uint64 `json:"id" dc:"Post ID" v:"required"`
}

// PostGetListReq get paginated post list request
type PostGetListReq struct {
	Page     int    `json:"page"     dc:"Page number" d:"1"`
	PageSize int    `json:"pageSize" dc:"Page size" d:"10"`
	AuthorId uint64 `json:"authorId" dc:"Author user ID (optional)"`
	Tags     string `json:"tags"     dc:"Tag filter (optional)"`
	Status   *int  `json:"status"   dc:"Status filter (optional)" v:"in:0,1,2,3"`
	Keyword  string `json:"keyword"  dc:"Keyword search (optional)"`
}

// PostGetPageListReq get paginated post list request (alias for Service layer)
type PostGetPageListReq struct {
	Page     uint   `json:"page"     dc:"Page number" d:"1"`
	PageSize uint   `json:"pageSize" dc:"Page size" d:"10"`
	AuthorId uint64 `json:"authorId" dc:"Author user ID (optional)"`
	Tags     string `json:"tags"     dc:"Tag filter (optional)"`
	Status   *int   `json:"status"   dc:"Status filter (optional)" v:"in:0,1,2,3"`
	Keyword  string `json:"keyword"  dc:"Keyword search (optional)"`
}

// PostResp post response
type PostResp struct {
	Id           uint64      `json:"id"            dc:"Post ID"`
	Title        string      `json:"title"         dc:"Post title"`
	Content      string      `json:"content"       dc:"Post content"`
	AuthorId     uint64      `json:"authorId"      dc:"Author user ID"`
	Status       int         `json:"status"        dc:"Status: 0-deleted 1-normal 2-pinned 3-featured"`
	ViewCount    uint        `json:"viewCount"     dc:"View count"`
	LikeCount    uint        `json:"likeCount"     dc:"Like count"`
	CommentCount uint        `json:"commentCount"  dc:"Comment count"`
	Tags         string      `json:"tags"          dc:"Tags"`
	CoverImage   string      `json:"coverImage"    dc:"Cover image URL"`
	IsDeleted    int         `json:"isDeleted"     dc:"Soft delete status"`
	CreatedAt    *gtime.Time `json:"createdAt"     dc:"Created time"`
	UpdatedAt    *gtime.Time `json:"updatedAt"     dc:"Updated time"`
}

// PostListResp post list response
type PostListResp struct {
	Total    uint64     `json:"total"    dc:"Total count"`
	Page     uint       `json:"page"     dc:"Current page"`
	PageSize uint       `json:"pageSize" dc:"Page size"`
	List     []*PostResp `json:"list"    dc:"Post list"`
}

// PostCreateResp create post response
type PostCreateResp struct {
	Id uint64 `json:"id" dc:"New post ID"`
}

// PostDeleteResp delete post response
type PostDeleteResp struct {
	Success bool `json:"success" dc:"Success"`
}

// PostUpdateResp update post response
type PostUpdateResp struct {
	Success bool `json:"success" dc:"Success"`
}
