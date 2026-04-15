package user

import (
	"context"

	"AI-Study-Community/internal/model/do"
	"AI-Study-Community/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	User = &cUser{}
)

type cUser struct{}

func (c *cUser) Create(ctx context.Context, req *do.UserReq) (res *do.UserResp, err error) {


	result, err := service.UserService.Create(ctx, req)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

    r := ghttp.RequestFromCtx(ctx)
    session := r.Session
	session.Set("userId",result.Id)


	return result, nil
}

func (c *cUser) RegisterRoute(s *ghttp.Server) {
	group := s.Group("/api/user")
	group.Middleware(ghttp.MiddlewareHandlerResponse)
	group.POST("/login", c.Create)

}
