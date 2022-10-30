package frontend

import (
	fe "1536509937/ku-bbs/internal/entity/frontend"
	sv "1536509937/ku-bbs/internal/service"
	"1536509937/ku-bbs/internal/service/frontend"
	"github.com/gin-gonic/gin"
)

var Home = home{}

type home struct{}

func (c *home) HomePage(ctx *gin.Context) {
	s := sv.Context(ctx)

	var req fe.GetTopicListReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	data, _ := frontend.TopicService(ctx).GetList(&req)
	nodes, _ := frontend.NodeService(ctx).GetEnableNodes()
	checked, _ := frontend.CheckinService(ctx).IsCheckin()

	data["nodes"] = nodes
	data["checked"] = checked

	s.View("frontend.home.index", data)
}
