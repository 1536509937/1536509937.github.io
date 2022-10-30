package frontend

import (
	"1536509937/ku-bbs/internal/service"
	"github.com/gin-gonic/gin"
)

func RemindService(ctx *gin.Context) *SRemind {
	return &SRemind{ctx: service.Context(ctx)}
}

type SRemind struct {
	ctx *service.BaseContext
}
