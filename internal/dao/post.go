package dao

import (
	"context"

	"AI-Study-Community/internal/model/do"
	"AI-Study-Community/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
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

// GetOne queries a single post record by ID.
// It returns the post entity or nil if not found.
func (p *Post) GetOne(ctx context.Context, id uint64) (*entity.Post, error) {
	var post *entity.Post
	err := p.db().Where("id", id).Scan(&post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// Delete soft-deletes a post by setting IsDeleted=1 and DeletedAt.
// It uses soft-delete pattern based on the IsDeleted field.
// Returns the number of affected rows.
func (p *Post) Delete(ctx context.Context, id uint64) (int64, error) {
	result, err := p.db().Model(p.table).
		Where("id", id).
		Data(g.Map{
			"is_deleted": 1,
			"deleted_at": gtime.Now(),
		}).
		Update()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// db returns the underlying database model for further operations
func (p *Post) db() *gdb.Model {
	return Database.Model(p.table)
}
