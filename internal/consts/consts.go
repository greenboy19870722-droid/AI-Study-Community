package consts

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func GetCurrentUserId(r *ghttp.Request) uint64 {
	session := r.Session
	userIdVal, err := session.Get("userId")
	if err != nil {
		return 0
	}
	return userIdVal.Uint64()
}
