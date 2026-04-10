package dao

import (
	"context"

	"AI-Study-Community/internal/model/do"
	"AI-Study-Community/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
)

// Comment is the DAO for comments table
type Comment struct {
	table string
}

// CommentDao is the singleton instance of Comment DAO
var CommentDao = &Comment{
	table: "comments",
}

// Insert inserts a new comment record into the database.
// It creates a comment with status=1 (normal) and depth=0 for top-level comments.
// Returns the inserted record's ID.
func (c *Comment) Insert(ctx context.Context, req *do.CommentCreateReq) (uint64, error) {
	// Determine depth: if ParentId is nil, depth is 0 (top-level comment)
	depth := uint(0)
	if req.ParentId != nil {
		// Get parent comment to determine depth
		parent, err := c.GetOne(ctx, *req.ParentId)
		if err != nil {
			return 0, err
		}
		if parent != nil {
			depth = parent.Depth + 1
		}
	}

	comment := entity.Comment{
		PostId:       req.PostId,
		ParentId:     req.ParentId,
		Content:      req.Content,
		AuthorId:     req.AuthorId,
		ReplyToUserId: req.ReplyToUserId,
		Depth:        depth,
		LikeCount:    0,
		Status:       1, // Normal status
		IsDeleted:    0, // Not deleted
	}

	result, err := c.db().Data(comment).Insert()
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

// GetOne queries a single comment record by ID.
// It returns the comment entity or nil if not found.
// Only returns non-deleted comments (is_deleted=0).
func (c *Comment) GetOne(ctx context.Context, id uint64) (*entity.Comment, error) {
	var comment entity.Comment
	record, err := c.db().Where("id", id).Where("is_deleted", 0).One()
	if err != nil {
		return nil, err
	}
	if record.IsEmpty() {
		return nil, nil
	}
	if err := record.Struct(&comment); err != nil {
		return nil, err
	}
	return &comment, nil
}

// Delete soft-deletes a comment by setting IsDeleted=1 and DeletedAt.
// It uses soft-delete pattern based on the IsDeleted field.
// Returns the number of affected rows.
func (c *Comment) Delete(ctx context.Context, id uint64) (int64, error) {
	result, err := c.db().
		Where("id", id).
		Data(map[string]interface{}{
			"is_deleted": 1,
			"deleted_at": gtime.Now(),
		}).
		Update()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// GetByPostId retrieves paginated top-level comments (parent_id IS NULL) for a post.
// Returns the comment list, total count, and any error encountered.
func (c *Comment) GetByPostId(ctx context.Context, req *do.CommentGetByPostIdReq) ([]*entity.Comment, int, error) {
	m := c.db()

	// Apply soft-delete filter (only show non-deleted comments)
	m = m.Where("is_deleted", 0)

	// Filter by post ID
	m = m.Where("post_id", req.PostId)

	// Only get top-level comments (parent_id IS NULL)
	m = m.Where("parent_id", nil)

	// Count total records
	total, err := m.Count()
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
	result, err := m.Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		All()
	if err != nil {
		return nil, 0, err
	}

	// Convert Result to []*entity.Comment
	comments := make([]*entity.Comment, 0, len(result))
	for _, record := range result {
		var comment entity.Comment
		if err := record.Struct(&comment); err != nil {
			// Skip records that fail to convert
			continue
		}
		comments = append(comments, &comment)
	}

	return comments, total, nil
}

// GetChildren retrieves child comments by parent ID.
// Returns the comment list, total count, and any error encountered.
func (c *Comment) GetChildren(ctx context.Context, req *do.CommentGetChildrenReq) ([]*entity.Comment, int, error) {
	m := c.db()

	// Apply soft-delete filter (only show non-deleted comments)
	m = m.Where("is_deleted", 0)

	// Filter by parent ID
	m = m.Where("parent_id", req.ParentId)

	// Count total records
	total, err := m.Count()
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

	// Fetch paginated results ordered by created_at ASC (chronological order)
	result, err := m.Order("created_at ASC").
		Offset(offset).
		Limit(pageSize).
		All()
	if err != nil {
		return nil, 0, err
	}

	// Convert Result to []*entity.Comment
	comments := make([]*entity.Comment, 0, len(result))
	for _, record := range result {
		var comment entity.Comment
		if err := record.Struct(&comment); err != nil {
			// Skip records that fail to convert
			continue
		}
		comments = append(comments, &comment)
	}

	return comments, total, nil
}

// db returns the underlying database model for further operations
func (c *Comment) db() *gdb.Model {
	return Database.Model(c.table)
}
