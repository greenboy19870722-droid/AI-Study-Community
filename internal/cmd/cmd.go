package cmd

import (
	"context"

	"AI-Study-Community/internal/controller/comment"
	"AI-Study-Community/internal/controller/post"
	"AI-Study-Community/internal/dao"

	"github.com/gogf/gf/v2/frame/g"
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
			post.Post.RegisterRoute(s)
			comment.Comment.RegisterRoute(s)
			s.Run()
			return nil
		},
	}
)
