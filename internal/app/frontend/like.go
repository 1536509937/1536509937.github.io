package frontend

import (
	fe "1536509937/ku-bbs/internal/entity/frontend"
	"1536509937/ku-bbs/internal/service"
	"1536509937/ku-bbs/internal/service/frontend"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"
)

var Like = cLike{}

type cLike struct{}

func (c *cLike) LikeSubmit(ctx *gin.Context) {
	s := service.Context(ctx)

	if !s.Check() {
		s.Json(gin.H{"code": 1, "msg": "请登录后在继续操作"})
		return
	}

	var req fe.LikeReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Json(gin.H{"code": 1, "msg": err.Error()})
		return
	}

	if err := g.Validator().Data(req).Run(context.Background()); err != nil {
		s.Json(gin.H{"code": 1, "msg": err.FirstError()})
		return
	}

	if err := frontend.LikeService(ctx).Like(&req); err != nil {
		s.Json(gin.H{"code": 1, "msg": err.Error()})
	} else {
		s.Json(gin.H{"code": 0, "msg": "ok"})
	}
}
