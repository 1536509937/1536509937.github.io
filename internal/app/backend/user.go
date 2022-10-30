package backend

import (
	"github.com/gin-gonic/gin"

	be "1536509937/ku-bbs/internal/entity/backend"
	sv "1536509937/ku-bbs/internal/service"
	bs "1536509937/ku-bbs/internal/service/backend"
)

var User = cUser{}

type cUser struct{}

func (c *cUser) IndexPage(ctx *gin.Context) {
	s := sv.Context(ctx)

	var req be.GetUserListReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	if data, err := bs.UserService(ctx).GetList(&req); err != nil {
		s.Back().WithError(err).Redirect()
	} else {
		s.View("backend.user.index", data)
	}
}
