package frontend

import (
	"1536509937/ku-bbs/internal/service"
	"1536509937/ku-bbs/internal/service/frontend"
	"github.com/gin-gonic/gin"
)

var Checkin = cCheckin{}

type cCheckin struct{}

func (c *cCheckin) StoreSubmit(ctx *gin.Context) {
	s := service.Context(ctx)

	if !s.Check() {
		s.Json(gin.H{"code": 1, "msg": "请登录后在继续操作"})
		return
	}

	if err := frontend.CheckinService(ctx).Store(); err != nil {
		s.Json(gin.H{"code": 1, "msg": err.Error()})
	} else {
		s.Json(gin.H{"code": 0, "msg": "ok"})
	}
}
