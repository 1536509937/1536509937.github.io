package frontend

import "1536509937/ku-bbs/internal/model"

type Follow struct {
	model.Follows
	Follower *model.Users `gorm:"foreignKey:user_id"`
	Fans     *model.Users `gorm:"foreignKey:target_id"`
}
