package main

import (
	_ "AI-Study-Community/internal/packed"

	"AI-Study-Community/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var (
		ctx = gctx.New()
	)
	cmd.Main.Run(ctx)
}
