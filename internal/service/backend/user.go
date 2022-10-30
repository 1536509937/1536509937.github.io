package backend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"1536509937/ku-bbs/internal/entity/backend"
	"1536509937/ku-bbs/internal/model"
	"1536509937/ku-bbs/internal/service"
	"1536509937/ku-bbs/pkg/utils/page"
)

func UserService(ctx *gin.Context) *sUser {
	return &sUser{ctx: service.Context(ctx)}
}

type sUser struct {
	ctx *service.BaseContext
}

func (s *sUser) GetList(req *backend.GetUserListReq) (gin.H, error) {
	if req.Page == 0 {
		req.Page = 1
	}

	var (
		users  []*backend.User
		total  int64
		limit  = 20
		offset = (req.Page - 1) * limit
	)

	builder := model.User().M

	if len(req.Keywords) > 0 {
		builder = builder.Where("name like ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}

	if c := builder.Count(&total); c.Error != nil {
		return nil, c.Error
	}

	f := builder.Limit(limit).Offset(offset).Find(&users)

	if f.Error != nil {
		return nil, f.Error
	}

	baseUrl := s.ctx.Ctx.Request.RequestURI

	pagination := page.New(int(total), limit, gconv.Int(req.Page), baseUrl)

	return gin.H{"list": users, "page": pagination, "req": req}, nil
}
