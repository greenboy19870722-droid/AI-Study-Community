package main

import (
	_ "AI-Study-Community/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"AI-Study-Community/internal/cmd"
)

func main() {
	var (
		ctx = gctx.New()
	)
	cmd.Main.Run(ctx)
}
