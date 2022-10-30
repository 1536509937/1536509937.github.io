package backend

import "1536509937/ku-bbs/internal/model"

type GetUserListReq struct {
	Page     int    `form:"page"`
	Keywords string `form:"keywords"`
}

type User struct {
	model.Users
}
