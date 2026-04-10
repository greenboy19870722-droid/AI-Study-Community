package service

import (
	"context"

	"AI-Study-Community/internal/dao"
	"AI-Study-Community/internal/model/do"
)

// Post is the service layer for posts
type Post struct {
}

// PostService is the singleton instance of Post service
var PostService = &Post{}

// Create creates a new post.
// It validates the input and calls the DAO layer to insert the post.
// Returns the created post ID.
func (s *Post) Create(ctx context.Context, req *do.PostCreateReq) (*do.PostCreateResp, error) {
	// Call DAO layer to insert the post
	id, err := dao.PostDao.Insert(ctx, req)
	if err != nil {
		return nil, err
	}

	return &do.PostCreateResp{
		Id: id,
	}, nil
}
