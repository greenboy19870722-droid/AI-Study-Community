package cmd

import (
	"context"

	"AI-Study-Community/internal/controller/comment"
	"AI-Study-Community/internal/controller/post"
	"AI-Study-Community/internal/controller/user"
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
				g.Log().Fatal(ctx, "Database initialization failed:", err)
				return err
			}

			// Get server instance and configure
			s := g.Server()

			// Configure server settings
			s.SetPort(8000)
			s.SetServerRoot("")
			s.SetOpenApiPath("/api.json")
			s.SetSwaggerPath("/swagger")

			// s.SetSessionCookieOutput(true)
			// s.SetSessionMaxAge(time.Hour * 24)
			// s.SetSessionStorage(gsession.NewStorageMemory())


			// Add global middleware
			// CORS middleware for cross-origin API access
			s.Use(ghttp.MiddlewareCORS)
			// Handler response middleware for unified response format
			s.Use(ghttp.MiddlewareHandlerResponse)

			user.User.RegisterRoute(s)
			
			// Register Post routes (/api/post/*)
			post.Post.RegisterRoute(s)

			// Register Comment routes (/api/comment/*)
			comment.Comment.RegisterRoute(s)

			// Run the server
			s.Run()

			return nil
		},
	}
)
