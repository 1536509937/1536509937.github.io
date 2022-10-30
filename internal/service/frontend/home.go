package frontend

import (
	"github.com/gin-gonic/gin"
	"1536509937/ku-bbs/internal/service"
)

func HomeService(ctx *gin.Context) *SHome {
	return &SHome{ctx: service.Context(ctx)}
}

type SHome struct {
	ctx *service.BaseContext
}
