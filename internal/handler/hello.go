package handler

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Hello = hHello{}
)

type hHello struct{}

func (h *hHello) Hello(ctx context.Context, req interface{}) (res interface{}, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
