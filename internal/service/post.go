package service

import (
	"context"
	"errors"

	"AI-Study-Community/internal/dao"
	"AI-Study-Community/internal/model/do"
	"AI-Study-Community/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// Post is the service layer for posts
type Post struct {
}

// PostService is the singleton instance of Post service
var PostService = &Post{}

// ErrPostNotOwned is returned when a user tries to modify a post they don't own
var ErrPostNotOwned = errors.New("post not owned by user")

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

// GetDetail retrieves a single post by ID.
// It calls the DAO layer to fetch the post details.
// Returns the post entity or nil if not found.
func (s *Post) GetDetail(ctx context.Context, id uint64) (*entity.Post, error) {
	return dao.PostDao.GetOne(ctx, id)
}

// Update updates an existing post by ID.
// It validates the input and calls the DAO layer to update the post.
// Returns true if the update was successful.
func (s *Post) Update(ctx context.Context, req *do.PostUpdateReq) (bool, error) {
	rowsAffected, err := dao.PostDao.Update(ctx, req)
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}

// Delete soft-deletes a post by ID.
// It calls the DAO layer to perform the deletion.
// Returns true if the deletion was successful.
func (s *Post) Delete(ctx context.Context, id uint64,currentUserId uint64) (bool, error) {


	post, err := dao.PostDao.GetOne(ctx, id)
	if err != nil {
		return false, err
	}

	if post.AuthorId != currentUserId {
		g.Log().Errorf(ctx, "Delete post permission denied: post.AuthorId=%d, currentUserId=%d", post.AuthorId, currentUserId)
		return false, ErrPostNotOwned
	}	

	rowsAffected, err := dao.PostDao.Delete(ctx, id)
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}

// GetPageList retrieves a paginated list of posts with optional filters.
// It calls the DAO layer to fetch the post list and total count.
// Returns the paginated post list response.
func (s *Post) GetPageList(ctx context.Context, req *do.PostGetPageListReq) (*do.PostListResp, error) {
	// Convert to DAO request format
	daoReq := &do.PostGetListReq{
		Page:     int(req.Page),
		PageSize: int(req.PageSize),
		AuthorId: req.AuthorId,
		Tags:     req.Tags,
		Status:   req.Status,
		Keyword:  req.Keyword,
	}

	// Call DAO layer to get paginated list
	list, total, err := dao.PostDao.GetList(ctx, daoReq)
	if err != nil {
		return nil, err
	}

	// Convert entity posts to response DTOs
	postList := make([]*do.PostResp, 0, len(list))
	for _, post := range list {
		postList = append(postList, &do.PostResp{
			Id:           post.Id,
			Title:        post.Title,
			Content:      post.Content,
			AuthorId:     post.AuthorId,
			Status:       post.Status,
			ViewCount:    post.ViewCount,
			LikeCount:    post.LikeCount,
			CommentCount: post.CommentCount,
			Tags:         post.Tags,
			CoverImage:   post.CoverImage,
			IsDeleted:    post.IsDeleted,
			CreatedAt:    post.CreatedAt,
			UpdatedAt:    post.UpdatedAt,
		})
	}

	return &do.PostListResp{
		Total:    uint64(total),
		Page:     req.Page,
		PageSize: req.PageSize,
		List:     postList,
	}, nil
}


