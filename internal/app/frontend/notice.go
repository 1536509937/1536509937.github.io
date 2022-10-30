package frontend

import (
	fe "1536509937/ku-bbs/internal/entity/frontend"
	sv "1536509937/ku-bbs/internal/service"
	"1536509937/ku-bbs/internal/service/frontend"
	"github.com/gin-gonic/gin"
)

var Notice = cNotice{}

type cNotice struct{}

func (*cNotice) HomePage(ctx *gin.Context) {
	s := sv.Context(ctx)

	if !s.Check() {
		s.To("/login").WithError("请登录后，再继续操作").Redirect()
		return
	}

	var req fe.GetRemindListReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.To("/").WithError(err).Redirect()
		return
	}

	noticeService := frontend.NoticeService(ctx)

	data, err := noticeService.GetList(&req)
	if err != nil {
		s.To("/").WithError(err.Error()).Redirect()
		return
	}

	remindUnread, _ := frontend.NoticeService(ctx).GetRemindUnread()

	systemUnread, _ := frontend.NoticeService(ctx).GetSystemUnread()

	data["remindUnread"] = remindUnread
	data["systemUnread"] = systemUnread

	noticeService.ReadAll(req.Type)

	s.View("frontend.notice.home", data)
}
