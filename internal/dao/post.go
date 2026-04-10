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

// GetList retrieves a paginated list of posts with optional filters.
// It supports filtering by authorId, tags, status, and keyword search.
// Returns the post list, total count, and any error encountered.
func (p *Post) GetList(ctx context.Context, req *do.PostGetListReq) ([]*entity.Post, int, error) {
	m := p.db()

	// Apply soft-delete filter (only show non-deleted posts)
	m = m.Where("is_deleted", 0)

	// Apply filters
	if req.AuthorId > 0 {
		m = m.Where("author_id", req.AuthorId)
	}
	if req.Tags != "" {
		m = m.Where("tags", req.Tags)
	}
	if req.Status != nil {
		m = m.Where("status", *req.Status)
	}
	if req.Keyword != "" {
		m = m.WhereLike("title", "%"+req.Keyword+"%").
			OrWhereLike("content", "%"+req.Keyword+"%")
	}

	// Count total records
	total, err := m.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Calculate offset with defaults
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	// Fetch paginated results ordered by created_at DESC
	list, err := m.Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Select(ctx)
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

// db returns the underlying database model for further operations
func (p *Post) db() *gdb.Model {
	return Database.Model(p.table)
}
