package cmd

import (
	"context"

	"AI-Study-Community/internal/controller/post"
	"AI-Study-Community/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// Initialize database connection
			if err = dao.Init(ctx); err != nil {
				g.Log().Fatal(ctx, err)
				return err
			}
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/api", func(group *ghttp.RouterGroup) {
					group.Group("/post", func(group *ghttp.RouterGroup) {
						group.POST("/create", post.Create)
						group.GET("/list", post.GetPageList)
					})
				})
			})
			s.Run()
			return nil
		},
	}
)
