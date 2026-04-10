package do

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CommentCreateReq creates comment request (first-level comment on post)
type CommentCreateReq struct {
	PostId   uint64 `json:"postId"   dc:"Post ID" v:"required"`
	Content  string `json:"content"  dc:"Comment content" v:"required|length:1,5000"`
	AuthorId uint64 `json:"authorId" dc:"Author user ID" v:"required"`
}

// CommentReplyReq reply to existing comment request
type CommentReplyReq struct {
	PostId        uint64  `json:"postId"        dc:"Post ID" v:"required"`
	ParentId      uint64  `json:"parentId"      dc:"Parent comment ID" v:"required"`
	Content       string  `json:"content"       dc:"Reply content" v:"required|length:1,5000"`
	AuthorId      uint64  `json:"authorId"      dc:"Author user ID" v:"required"`
	ReplyToUserId *uint64 `json:"replyToUserId" dc:"User ID being replied to"`
}

// CommentDeleteReq delete comment request
type CommentDeleteReq struct {
	Id uint64 `json:"id" dc:"Comment ID" v:"required"`
}

// CommentGetOneReq get one comment request
type CommentGetOneReq struct {
	Id uint64 `json:"id" dc:"Comment ID" v:"required"`
}

// CommentGetListReq get paginated comment list request
type CommentGetListReq struct {
	Page     uint   `json:"page"     dc:"Page number" d:"1"`
	PageSize uint   `json:"pageSize" dc:"Page size" d:"20"`
	PostId   uint64 `json:"postId"   dc:"Post ID filter (optional)"`
}

// CommentGetByPostIdReq get comments by post id request
type CommentGetByPostIdReq struct {
	Page     uint   `json:"page"     dc:"Page number" d:"1"`
	PageSize uint   `json:"pageSize" dc:"Page size" d:"20"`
	PostId   uint64 `json:"postId"   dc:"Post ID" v:"required"`
}

// CommentGetTreeReq get comment tree by post id request
type CommentGetTreeReq struct {
	Page     uint   `json:"page"     dc:"Page number" d:"1"`
	PageSize uint   `json:"pageSize" dc:"Page size" d:"10"`
	PostId   uint64 `json:"postId"   dc:"Post ID" v:"required"`
}

// CommentResp comment response
type CommentResp struct {
	Id            uint64      `json:"id"            dc:"Comment ID"`
	PostId        uint64      `json:"postId"        dc:"Post ID"`
	ParentId      *uint64     `json:"parentId"      dc:"Parent comment ID"`
	Content       string      `json:"content"       dc:"Comment content"`
	AuthorId      uint64      `json:"authorId"      dc:"Author user ID"`
	ReplyToUserId *uint64     `json:"replyToUserId" dc:"Replied-to user ID"`
	Depth         int         `json:"depth"         dc:"Comment depth (0=first level)"`
	LikeCount     uint        `json:"likeCount"     dc:"Like count"`
	Status        int         `json:"status"        dc:"Status: 0-deleted 1-normal"`
	IsDeleted     int         `json:"isDeleted"     dc:"Soft delete: 0-not deleted 1-deleted"`
	CreatedAt     *gtime.Time `json:"createdAt"    dc:"Created time"`
	UpdatedAt     *gtime.Time `json:"updatedAt"    dc:"Updated time"`
}

// CommentTreeResp comment tree node response (includes children)
type CommentTreeResp struct {
	Id            uint64           `json:"id"            dc:"Comment ID"`
	PostId        uint64           `json:"postId"        dc:"Post ID"`
	ParentId      *uint64          `json:"parentId"     dc:"Parent comment ID"`
	Content       string           `json:"content"      dc:"Comment content"`
	AuthorId      uint64          `json:"authorId"     dc:"Author user ID"`
	ReplyToUserId *uint64          `json:"replyToUserId" dc:"Replied-to user ID"`
	Depth         int              `json:"depth"         dc:"Comment depth (0=first level)"`
	LikeCount     uint             `json:"likeCount"    dc:"Like count"`
	Status        int              `json:"status"       dc:"Status: 0-deleted 1-normal"`
	IsDeleted     int              `json:"isDeleted"    dc:"Soft delete: 0-not deleted 1-deleted"`
	CreatedAt     *gtime.Time      `json:"createdAt"   dc:"Created time"`
	UpdatedAt     *gtime.Time      `json:"updatedAt"   dc:"Updated time"`
	Children      []*CommentTreeResp `json:"children"  dc:"Child comments (nested)"`

}

// CommentListResp comment list response
type CommentListResp struct {
	Total    uint64         `json:"total"    dc:"Total count"`
	Page     uint           `json:"page"     dc:"Current page"`
	PageSize uint           `json:"pageSize" dc:"Page size"`
	List     []*CommentResp `json:"list"    dc:"Comment list"`
}

// CommentTreeListResp comment tree list response (flat list of top-level comments with nested children)
type CommentTreeListResp struct {
	Total    uint64           `json:"total"    dc:"Total count (top-level only)"`
	Page     uint             `json:"page"     dc:"Current page"`
	PageSize uint             `json:"pageSize" dc:"Page size"`
	List     []*CommentTreeResp `json:"list" dc:"Comment tree list (top-level with nested children)"`
}

// CommentCreateResp create comment response
type CommentCreateResp struct {
	Id uint64 `json:"id" dc:"New comment ID"`
}

// CommentReplyResp reply comment response
type CommentReplyResp struct {
	Id uint64 `json:"id" dc:"New reply comment ID"`
}

// CommentDeleteResp delete comment response
type CommentDeleteResp struct {
	Success bool `json:"success" dc:"Success"`
}
