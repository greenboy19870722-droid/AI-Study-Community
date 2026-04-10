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
// It accepts post creation parameters and returns the created post ID.
func (c *cPost) Create(ctx context.Context, req *do.PostCreateReq) (res *do.PostCreateResp, err error) {
	result, err := service.PostService.Create(ctx, req)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return result, nil
}

// GetDetail handles GET /api/post/detail
// It accepts a post ID parameter and returns the post details.
func (c *cPost) GetDetail(ctx context.Context, req *do.PostGetOneReq) (res *do.PostResp, err error) {
	result, err := service.PostService.GetDetail(ctx, req.Id)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
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

// List handles GET /api/post/list
// It accepts pagination parameters (page, pageSize) and returns a paginated list of posts.
func (c *cPost) List(ctx context.Context, req *do.PostGetPageListReq) (res *do.PostListResp, err error) {
	// Validate and set defaults for pagination
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	result, err := service.PostService.GetPageList(ctx, req)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return result, nil
}

// Update handles POST /api/post/update
// It accepts post update parameters and updates the post via the service layer.
func (c *cPost) Update(ctx context.Context, req *do.PostUpdateReq) (res *do.PostUpdateResp, err error) {
	success, err := service.PostService.Update(ctx, req)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return &do.PostUpdateResp{Success: success}, nil
}

// Delete handles POST /api/post/delete
// It accepts a post ID and soft-deletes the post via the service layer.
func (c *cPost) Delete(ctx context.Context, req *do.PostDeleteReq) (res *do.PostDeleteResp, err error) {
	success, err := service.PostService.Delete(ctx, req.Id)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return &do.PostDeleteResp{Success: success}, nil
}

// RegisterRoute registers the post controller routes.
func (c *cPost) RegisterRoute(s *ghttp.Server) {
	group := s.Group("/api/post")
	group.Middleware(ghttp.MiddlewareHandlerResponse)
	group.POST("/create", c.Create)
	group.POST("/update", c.Update)
	group.POST("/delete", c.Delete)
	group.GET("/list", c.List)
	group.GET("/detail", c.GetDetail)
}
