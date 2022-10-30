package frontend

import "1536509937/ku-bbs/internal/model"

type SystemUserNotice struct {
	model.SystemUserNotices
	Notice model.SystemNotices `gorm:"foreignKey:notice_id"`
}
