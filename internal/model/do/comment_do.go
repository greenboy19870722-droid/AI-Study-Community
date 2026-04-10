package do

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CommentCreateReq creates comment request
type CommentCreateReq struct {
	PostId        uint64  `json:"postId" v:"required"`
	ParentId      *uint64 `json:"parentId"`
	Content       string  `json:"content" v:"required"`
	AuthorId      uint64  `json:"authorId" v:"required"`
	ReplyToUserId *uint64 `json:"replyToUserId"`
}

// CommentDeleteReq delete comment request
type CommentDeleteReq struct {
	Id uint64 `json:"id" v:"required"`
}

// CommentUpdateReq update comment request
type CommentUpdateReq struct {
	Id      uint64 `json:"id" v:"required"`
	Content string `json:"content" v:"required"`
}

// CommentGetOneReq get one comment request
type CommentGetOneReq struct {
	Id uint64 `json:"id" v:"required"`
}

// CommentGetByPostIdReq get comments by post ID request
type CommentGetByPostIdReq struct {
	PostId   uint64 `json:"postId" v:"required"`
	Page     int    `json:"page" d:"1"`
	PageSize int    `json:"pageSize" d:"10"`
}

// CommentGetChildrenReq get child comments by parent ID request
type CommentGetChildrenReq struct {
	ParentId uint64 `json:"parentId" v:"required"`
	Page     int    `json:"page" d:"1"`
	PageSize int    `json:"pageSize" d:"10"`
}

// CommentResp comment response
type CommentResp struct {
	Id            uint64        `json:"id"`
	PostId        uint64        `json:"postId"`
	UserId        uint64        `json:"userId"       dc:"User ID"`
	Content       string        `json:"content"`
	ParentId      *uint64       `json:"parentId"`
	AuthorId      uint64        `json:"authorId"`
	AuthorName    string        `json:"authorName"`
	ReplyToUserId *uint64       `json:"replyToUserId"`
	ReplyToAuthor string        `json:"replyToAuthor"`
	Depth         uint          `json:"depth"`
	LikeCount     uint          `json:"likeCount"`
	Status        int           `json:"status"`
	CreatedAt     *gtime.Time   `json:"createdAt"`
	UpdatedAt     *gtime.Time   `json:"updatedAt"`
	Children      []*CommentResp `json:"children"`
}

// CommentListResp comment list response
type CommentListResp struct {
	Total    int            `json:"total"`
	Page     int            `json:"page"`
	PageSize int            `json:"pageSize"`
	List     []*CommentResp `json:"list"`
}

// CommentReplyReq reply comment request
type CommentReplyReq struct {
	PostId        uint64  `json:"postId" v:"required"`
	ParentId      uint64  `json:"parentId" v:"required"`
	Content       string  `json:"content" v:"required"`
	AuthorId      uint64  `json:"authorId" v:"required"`
	ReplyToUserId *uint64 `json:"replyToUserId"`
}

// CommentGetTreeReq get comment tree request
type CommentGetTreeReq struct {
	PostId   uint64 `json:"postId" v:"required"`
	Page     int    `json:"page" d:"1"`
	PageSize int    `json:"pageSize" d:"20"`
}

// CommentTreeResp comment tree response
type CommentTreeResp struct {
	List []*CommentResp `json:"list" dc:"Comment list (flat)"`
	Tree []*CommentResp `json:"tree" dc:"Comment tree (hierarchical)"`
}

// CommentCreateResp create comment response
type CommentCreateResp struct {
	Id uint64 `json:"id"`
}

// CommentDeleteResp delete comment response
type CommentDeleteResp struct {
	Success bool `json:"success"`
}
