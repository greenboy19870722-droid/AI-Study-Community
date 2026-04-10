package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Comment 评论表实体
type Comment struct {
	Id             uint64      `json:"id"              dc:"评论ID"`
	PostId         uint64      `json:"postId"          dc:"所属帖子ID" v:"required"`
	ParentId       *uint64     `json:"parentId"        dc:"父评论ID(NULL为一级评论)"`
	Content        string      `json:"content"         dc:"评论内容" v:"required"`
	AuthorId       uint64      `json:"authorId"        dc:"评论者用户ID" v:"required"`
	ReplyToUserId  *uint64     `json:"replyToUserId"   dc:"被回复的用户ID(NULL则是一级评论)"`
	Depth          int         `json:"depth"           dc:"评论层级深度(0=一级)" v:"min:0"`
	LikeCount      uint        `json:"likeCount"       dc:"点赞数"`
	Status         int         `json:"status"          dc:"状态: 0-删除 1-正常" v:"in:0,1"`
	IsDeleted      int         `json:"isDeleted"       dc:"软删除: 0-未删除 1-已删除"`
	CreatedAt      *gtime.Time `json:"createdAt"      dc:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      dc:"更新时间"`
	DeletedAt      *gtime.Time `json:"deletedAt"      dc:"删除时间"`
}
