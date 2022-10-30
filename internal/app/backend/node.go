package backend

import (
	be "1536509937/ku-bbs/internal/entity/backend"
	sv "1536509937/ku-bbs/internal/service"
	bs "1536509937/ku-bbs/internal/service/backend"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var Node = cNode{}

type cNode struct{}

// IndexPage 节点管理
func (c *cNode) IndexPage(ctx *gin.Context) {
	s := sv.Context(ctx)

	var req be.GetNodeListReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	if data, err := bs.NodeService(ctx).GetList(&req); err != nil {
		s.Back().WithError(err).Redirect()
	} else {
		s.View("backend.node.index", data)
	}
}

// CreatePage 添加节点
func (c *cNode) CreatePage(ctx *gin.Context) {
	s := sv.Context(ctx)
	s.View("backend.node.create", nil)
}

// CreateSubmit 提交节点
func (c *cNode) CreateSubmit(ctx *gin.Context) {
	s := sv.Context(ctx)

	var req be.CreateNodeReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	if err := g.Validator().Data(req).Run(context.Background()); err != nil {
		s.Back().WithError(err.FirstError()).Redirect()
		return
	}

	if err := bs.NodeService(ctx).Create(&req); err != nil {
		s.Back().WithError(err).Redirect()
	} else {
		s.To("/backend/nodes").WithMsg("发布成功").Redirect()
	}
}

// EditPage 编辑节点
func (c *cNode) EditPage(ctx *gin.Context) {
	s := sv.Context(ctx)
	if node, err := bs.NodeService(ctx).GetDetail(gconv.Uint64(ctx.Param("id"))); err != nil {
		s.To("/backend/nodes").WithError(err).Redirect()
	} else {
		s.View("backend.node.edit", node)
	}
}

// EditSubmit 编辑提交
func (c *cNode) EditSubmit(ctx *gin.Context) {

	s := sv.Context(ctx)

	var req be.CreateNodeReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Back().WithError(err).Redirect()
		return
	}

	if err := g.Validator().Data(req).Run(context.Background()); err != nil {
		s.Back().WithError(err.FirstError()).Redirect()
		return
	}

	if err := bs.NodeService(ctx).Edit(gconv.Uint64(ctx.Param("id")), &req); err != nil {
		s.Back().WithError(err).Redirect()
	} else {
		s.To("/backend/nodes").WithMsg("编辑成功").Redirect()
	}
}

// DeleteSubmit 删除提交
func (c *cNode) DeleteSubmit(ctx *gin.Context) {
	s := sv.Context(ctx)
	t := "/backend/nodes"

	id := gconv.Int64(ctx.Param("id"))
	if id <= 0 {
		s.To(t).WithError("参数错误").Redirect()
		return
	}

	if err := bs.NodeService(ctx).Delete(id); err != nil {
		s.To(t).WithError("删除失败").Redirect()
	} else {
		s.To(t).WithMsg("删除成功").Redirect()
	}
}
