package post

import (
	"context"

	"AI-Study-Community/internal/model/do"
	"AI-Study-Community/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	Post = &cPost{}
)

type cPost struct{}

// Create handles POST /api/post/create
// It accepts post creation parameters and calls the service layer to create a new post.
func (c *cPost) Create(ctx context.Context, req *do.PostCreateReq) (res *do.PostCreateResp, err error) {
	result, err := service.PostService.Create(ctx, req)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return result, nil
}

// Detail handles GET /api/post/detail
// It accepts a post ID and returns the post details.
func (c *cPost) Detail(ctx context.Context, req *do.PostGetOneReq) (res *do.PostResp, err error) {
	post, err := service.PostService.GetDetail(ctx, req.Id)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if post == nil {
		return nil, nil
	}
	return &do.PostResp{
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
	}, nil
}

// GetPageList handles GET /api/post/list
// It accepts pagination parameters (page, pageSize) and returns a paginated list of posts.
func (c *cPost) GetPageList(ctx context.Context, req *do.PostGetPageListReq) (res *do.PostListResp, err error) {
	// Validate and set defaults for pagination
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}

	// Call service layer to get paginated post list
	result, err := service.PostService.GetPageList(ctx, req)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return result, nil
}

// List is an alias for GetPageList to match the route registration.
func (c *cPost) List(ctx context.Context, req *do.PostGetPageListReq) (res *do.PostListResp, err error) {
	return c.GetPageList(ctx, req)
}

// GetDetail handles GET /api/post/detail
// It accepts a post ID parameter and returns the post details.
func (c *cPost) GetDetail(ctx context.Context, req *do.PostGetOneReq) (res *do.PostResp, err error) {
	// Call service layer to get post detail
	result, err := service.Post.GetDetail(ctx, req.Id)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	// Convert entity to response DTO
	res = &do.PostResp{
		Id:           result.Id,
		Title:        result.Title,
		Content:      result.Content,
		AuthorId:     result.AuthorId,
		Status:       result.Status,
		ViewCount:    result.ViewCount,
		LikeCount:    result.LikeCount,
		CommentCount: result.CommentCount,
		Tags:         result.Tags,
		CoverImage:   result.CoverImage,
		IsDeleted:    result.IsDeleted,
		CreatedAt:    result.CreatedAt,
		UpdatedAt:    result.UpdatedAt,
	}

	return res, nil
}

// Create handles POST /api/post/create
// It accepts post creation parameters and returns the created post ID.
func (c *cPost) Create(ctx context.Context, req *do.PostCreateReq) (res *do.PostCreateResp, err error) {
	// Validate required fields are already handled by gvalid via service layer
	// Call service layer to create the post
	result, err := service.Post.Create(ctx, req)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return result, nil
}

// RegisterRoute registers the post controller routes.
func (c *cPost) RegisterRoute(s *ghttp.Server) {
	group := s.Group("/api/post")
	group.Middleware(ghttp.MiddlewareHandlerResponse)
	group.Bind(
		c.Create,
		c.GetPageList,
		c.GetDetail,
	)
}
