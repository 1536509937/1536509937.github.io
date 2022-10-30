package backend

import (
	srv "1536509937/ku-bbs/internal/service"
	"github.com/gin-gonic/gin"
)

var Home = cHome{}

type cHome struct{}

func (c *cHome) IndexPage(ctx *gin.Context) {
	srv.Context(ctx).View("backend.home.index", gin.H{})
}
