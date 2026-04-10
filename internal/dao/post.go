package dao

import (
	"context"

	"AI-Study-Community/internal/model/do"
	"AI-Study-Community/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// Post is the DAO for posts table
type Post struct {
	table string
}

// PostDao is the singleton instance of Post DAO
var PostDao = &Post{
	table: "posts",
}

// Insert inserts a new post record into the database.
// It creates a post with status=1 (normal) and returns the inserted record's ID.
func (p *Post) Insert(ctx context.Context, req *do.PostCreateReq) (uint64, error) {
	post := entity.Post{
		Title:      req.Title,
		Content:    req.Content,
		AuthorId:   req.AuthorId,
		Status:      1, // Normal status
		Tags:       req.Tags,
		CoverImage: req.CoverImage,
		IsDeleted:  0, // Not deleted
	}

	result, err := p.db().Model(p.table).Data(post).Insert()
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

// db returns the underlying database model for further operations
func (p *Post) db() *gdb.Model {
	return Database.Model(p.table)
}
