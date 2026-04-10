package entity

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/gmodel"
	"github.com/gogf/gf/v2/os/gtime"
)

// Post 帖子表实体
type Post struct {
	gmd.Meta `name:"posts" table:"posts" description:"帖子表"`
	Id          uint64      `json:"id"          dc:"帖子ID"`
	Title       string      `json:"title"       dc:"帖子标题" v:"required|length:1,255"`
	Content     string      `json:"content"     dc:"帖子内容(Markdown/富文本)" v:"required"`
	AuthorId    uint64      `json:"authorId"    dc:"作者用户ID" v:"required"`
	Status      int         `json:"status"      dc:"状态: 0-删除 1-正常 2-置顶 3-精华" v:"in:0,1,2,3"`
	ViewCount   uint        `json:"viewCount"   dc:"浏览量"`
	LikeCount   uint        `json:"likeCount"   dc:"点赞数"`
	CommentCount uint       `json:"commentCount" dc:"评论数"`
	Tags        string      `json:"tags"        dc:"标签,多个逗号分隔"`
	CoverImage  string      `json:"coverImage"  dc:"封面图片URL"`
	IsDeleted   int         `json:"isDeleted"   dc:"软删除: 0-未删除 1-已删除"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"更新时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   dc:"删除时间"`
}

// gmd.Meta implements gmodel.Model for ORM operations.
// It defines metadata for the Post entity.
func (Post) Meta() gdb.ModelMeta {
	return gdb.ModelMeta{
		FieldTimestampsMap: map[string]string{
			"createdAt": "created_at",
			"updatedAt": "updated_at",
		},
		FieldDeleterWithDeletedAt: true,
	}
}
