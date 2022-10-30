package frontend

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"
	fe "1536509937/ku-bbs/internal/entity/frontend"
	"1536509937/ku-bbs/internal/service"
	"1536509937/ku-bbs/internal/service/frontend"
)

var Auth = auth{}

type auth struct{}

func (c *auth) RegisterPage(ctx *gin.Context) {
	service.Context(ctx).View("frontend.auth.register", gin.H{})
}

func (c *auth) RegisterSubmit(ctx *gin.Context) {
	s := service.Context(ctx)
	p := ctx.DefaultQuery("back", "/")

	var req fe.RegisterReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	if err := g.Validator().Data(req).Run(context.Background()); err != nil {
		s.Back().WithError(err.FirstError()).Redirect()
		return
	}

	if err := frontend.UserService(ctx).Register(&req); err != nil {
		s.Back().WithError(err).Redirect()
	} else {
		s.To(p).WithMsg("注册成功， 欢迎来到学习论坛").Redirect()
	}
}

func (c *auth) LoginPage(ctx *gin.Context) {
	p := ctx.DefaultQuery("back", "/")
	s := service.Context(ctx)

	if s.Check() {
		s.To(p).Redirect()
	} else {
		s.View("frontend.auth.login", gin.H{"path": p})
	}
}

func (c *auth) LoginSubmit(ctx *gin.Context) {
	s := service.Context(ctx)
	p := ctx.DefaultQuery("back", "/")

	var req fe.LoginReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	if err := g.Validator().Data(req).Run(context.Background()); err != nil {
		s.Back().WithError(err.FirstError()).Redirect()
		return
	}

	if err := frontend.UserService(ctx).Login(&req); err != nil {
		s.Back().WithError(err).Redirect()
	} else {
		s.To(p).WithMsg("登录成功， 欢迎来到学习论坛").Redirect()
	}
}

func (c *auth) LogoutSubmit(ctx *gin.Context) {
	s := service.Context(ctx)

	frontend.UserService(ctx).Logout()

	s.To("/").WithMsg("退出成功").Redirect()
}
