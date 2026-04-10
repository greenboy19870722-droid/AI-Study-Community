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
	result, err := service.Post.GetPageList(ctx, req)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return result, nil
}

// Delete handles POST /api/post/delete
// It accepts a post ID and calls the service layer to soft-delete the post.
func (c *cPost) Delete(ctx context.Context, req *do.PostDeleteReq) (res *do.PostDeleteResp, err error) {
	success, err := service.Post.Delete(ctx, req.Id)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	return &do.PostDeleteResp{
		Success: success,
	}, nil
}

// RegisterRoute registers the post controller routes.
func (c *cPost) RegisterRoute(s *ghttp.Server) {
	group := s.Group("/api/post")
	group.Middleware(ghttp.MiddlewareHandlerResponse)
	group.Bind(
		ghttp.HandlerFunc(c.GetPageList),
		ghttp.HandlerFunc(c.Delete),
	)
}
