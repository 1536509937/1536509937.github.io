package route

import (
	"github.com/gin-gonic/gin"
	srv "1536509937/ku-bbs/internal/service"
	"1536509937/ku-bbs/pkg/config"
)

func isAdmin(ctx *gin.Context) {
	if s := srv.Context(ctx); !s.IsAdmin() {
		s.To("/").WithError("无权限访问").Redirect()
		ctx.Abort()
		return
	} else {
		ctx.Next()
	}
}

func visitor(ctx *gin.Context) {
	if s := srv.Context(ctx); config.Conf.App.VisitMode == "auth" && !s.Check() {
		s.To("/login").WithError("请登录后，再继续操作").Redirect()
		ctx.Abort()
		return
	} else {
		ctx.Next()
	}
}
