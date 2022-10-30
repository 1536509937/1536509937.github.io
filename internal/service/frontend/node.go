package frontend

import (
	"github.com/gin-gonic/gin"
	"1536509937/ku-bbs/internal/consts"
	"1536509937/ku-bbs/internal/model"
	"1536509937/ku-bbs/internal/service"
)

func NodeService(ctx *gin.Context) *SNode {
	return &SNode{ctx: service.Context(ctx)}
}

type SNode struct {
	ctx *service.BaseContext
}

func (s *SNode) GetEnableNodes() ([]*model.Nodes, error) {
	var nodes []*model.Nodes
	r := model.Node().M.Where("state", consts.EnableState).Order("sort DESC").Find(&nodes)
	if r.Error != nil {
		return nil, r.Error
	} else {
		return nodes, nil
	}
}
