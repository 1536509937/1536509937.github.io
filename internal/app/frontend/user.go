package frontend

import (
	fe "1536509937/ku-bbs/internal/entity/frontend"
	"1536509937/ku-bbs/internal/service"
	"1536509937/ku-bbs/internal/service/frontend"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"
)

var User = cUser{}

type cUser struct{}

func (c *cUser) HomePage(ctx *gin.Context) {
	s := service.Context(ctx)

	var req fe.GetUserHomeReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.To("/").WithError(err).Redirect()
		return
	}
	if data, err := frontend.UserService(ctx).Home(&req); err != nil {
		s.To("/").WithError(err).Redirect()
	} else {
		s.View("frontend.user.home", data)
	}
}

func (c *cUser) EditPage(ctx *gin.Context) {
	s := service.Context(ctx)
	t := ctx.DefaultQuery("tab", "info")

	if !s.Check() {
		s.To("/login").WithError("请登录后，再继续操作").Redirect()
	} else {
		s.View("frontend.user.edit", gin.H{"tab": t})
	}
}

func (c *cUser) EditSubmit(ctx *gin.Context) {
	s := service.Context(ctx)
	t := ctx.DefaultQuery("tab", "info")
	p := ctx.Request.RequestURI

	if t == "info" {
		var req fe.EditUserReq
		if err := ctx.ShouldBind(&req); err != nil {
			s.Back().WithError(err).Redirect()
			return
		}

		if err := g.Validator().Data(req).Run(context.Background()); err != nil {
			s.To(p).WithError(err.FirstError()).Redirect()
			return
		}

		if err := frontend.UserService(ctx).Edit(&req); err != nil {
			s.To(p).WithError(err).Redirect()
		} else {
			s.Back().WithMsg("修改信息成功").Redirect()
		}
	} else if t == "pass" {
		var req fe.EditPasswordReq
		if err := ctx.ShouldBind(&req); err != nil {
			s.Back().WithError(err).Redirect()
			return
		}

		if err := g.Validator().Data(req).Run(context.Background()); err != nil {
			s.To(p).WithError(err.FirstError()).Redirect()
			return
		}

		if err := frontend.UserService(ctx).EditPassword(&req); err != nil {
			s.To(p).WithError(err).Redirect()
		} else {
			s.To("/login").WithMsg("修改密码成功，请重新登录").Redirect()
		}
	} else {
		if err := frontend.UserService(ctx).EditAvatar(ctx); err != nil {
			s.To(p).WithError(err).Redirect()
		} else {
			s.Back().WithMsg("修改头像成功").Redirect()
		}
	}
}
