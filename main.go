package main

import (
	_ "AI-Study-Community/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"AI-Study-Community/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var (
		ctx = gctx.New()
	)
	cmd.Main.Run(ctx)
}
