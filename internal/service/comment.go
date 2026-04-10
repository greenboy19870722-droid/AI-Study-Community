package service

import (
	"context"

	"AI-Study-Community/internal/dao"
	"AI-Study-Community/internal/model/do"
)

// Comment is the service layer for comments
type Comment struct {
}

// CommentService is the singleton instance of Comment service
var CommentService = &Comment{}

// Create creates a new comment.
// It validates the input and calls the DAO layer to insert the comment.
// Returns the created comment ID.
func (s *Comment) Create(ctx context.Context, req *do.CommentCreateReq) (*do.CommentCreateResp, error) {
	// Call DAO layer to insert the comment
	id, err := dao.CommentDao.Insert(ctx, req)
	if err != nil {
		return nil, err
	}

	return &do.CommentCreateResp{
		Id: id,
	}, nil
}
