package comment

import (
	"context"

	"AI-Study-Community/internal/model/do"
	"AI-Study-Community/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	Comment = &cComment{}
)

type cComment struct{}

// Create handles POST /api/comment/create
// It accepts comment creation parameters and returns the created comment ID.
func (c *cComment) Create(ctx context.Context, req *do.CommentCreateReq) (res *do.CommentCreateResp, err error) {
	result, err := service.CommentService.Create(ctx, req)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return result, nil
}

// Reply handles POST /api/comment/reply
// It accepts reply comment parameters and returns the created reply comment ID.
func (c *cComment) Reply(ctx context.Context, req *do.CommentReplyReq) (res *do.CommentReplyResp, err error) {
	result, err := service.CommentService.Reply(ctx, req)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return result, nil
}

// Delete handles POST /api/comment/delete
// It accepts a comment ID and soft-deletes the comment via the service layer.
func (c *cComment) Delete(ctx context.Context, req *do.CommentDeleteReq) (res *do.CommentDeleteResp, err error) {
	result, err := service.CommentService.Delete(ctx, req)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return result, nil
}

// GetDetail handles GET /api/comment/detail
// It accepts a comment ID parameter and returns the comment details.
func (c *cComment) GetDetail(ctx context.Context, req *do.CommentGetOneReq) (res *do.CommentResp, err error) {
	result, err := service.CommentService.GetDetail(ctx, req)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
	return result, nil
}

// GetTree handles GET /api/comment/tree
// It accepts post ID and pagination parameters and returns the comment tree structure.
func (c *cComment) GetTree(ctx context.Context, req *do.CommentGetTreeReq) (res *do.CommentTreeResp, err error) {
	// Validate and set defaults for pagination
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 20
	}
	result, err := service.CommentService.GetTreeByPostId(ctx, req)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return result, nil
}

// RegisterRoute registers the comment controller routes.
func (c *cComment) RegisterRoute(s *ghttp.Server) {
	group := s.Group("/api/comment")
	group.Middleware(ghttp.MiddlewareHandlerResponse)
	group.Bind(
		c.Create,
		c.Reply,
		c.Delete,
		c.GetDetail,
		c.GetTree,
	)
}
