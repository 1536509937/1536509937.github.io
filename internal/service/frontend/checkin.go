package frontend

import (
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"1536509937/ku-bbs/internal/consts"
	"1536509937/ku-bbs/internal/model"
	"1536509937/ku-bbs/internal/service"
	"1536509937/ku-bbs/pkg/db"
	time2 "1536509937/ku-bbs/pkg/utils/time"
	"gorm.io/gorm"
)

func CheckinService(ctx *gin.Context) *SCheckin {
	return &SCheckin{ctx: service.Context(ctx)}
}

type SCheckin struct {
	ctx *service.BaseContext
}

func (s *SCheckin) Store() error {
	uid := s.ctx.Auth().ID

	var checkin model.Checkins
	f := model.Checkin().M.Where("user_id", uid).Find(&checkin)
	if f.Error != nil {
		return f.Error
	}

	if checkin.ID > 0 && checkin.LastTime.Format("2006-01-02") >= time.Now().Format("2006-01-02") {
		return errors.New("请勿重复签到")
	}

	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if checkin.ID > 0 {
			preDate := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
			data := map[string]interface{}{
				"cumulative_days": gorm.Expr("cumulative_days + 1"),
				"last_time":       time.Now(),
			}
			if checkin.LastTime.Format("2006-01-02") == preDate {
				data["continuity_days"] = gorm.Expr("continuity_days + 1")
			} else {
				data["continuity_days"] = 1
			}
			u := tx.Model(&model.Checkins{}).Where("id", checkin.ID).Where("last_time", checkin.LastTime).Updates(data)
			if u.Error != nil || u.RowsAffected <= 0 {
				log.Println(u.Error)
				return errors.New("签到失败")
			}
		} else {
			c := tx.Model(&model.Checkins{}).Create(&model.Checkins{
				UserId:         uid,
				CumulativeDays: 1,
				ContinuityDays: 1,
				LastTime:       time.Now(),
			})
			if c.Error != nil || c.RowsAffected <= 0 {
				log.Println(c.Error)
				return errors.New("签到失败")
			}
		}

		c := tx.Model(&model.IntegralLogs{}).Create(&model.IntegralLogs{
			UserId:  uid,
			Rewards: consts.CHECKINReward,
			Mode:    consts.CheckinMode,
		})
		if c.Error != nil || c.RowsAffected <= 0 {
			log.Println(c.Error)
			return errors.New("签到失败")
		}

		u := tx.Model(&model.Users{}).Where("id", uid).Updates(map[string]interface{}{
			"integral": gorm.Expr("integral + ?", consts.CHECKINReward),
		})
		if u.Error != nil || u.RowsAffected <= 0 {
			log.Println(u.Error)
			return errors.New("签到失败")
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *SCheckin) IsCheckin() (bool, error) {
	if !s.ctx.Check() {
		return false, nil
	}

	date := time2.ToDateString(time.Now())

	startAt := date + " 00:00:00"
	endedAt := date + " 23:59:59"

	var checkin *model.Checkins
	f := model.Checkin().M.
		Where("last_time >= ?", startAt).
		Where("last_time <= ?", endedAt).
		Where("user_id", s.ctx.Auth().ID).
		Find(&checkin)
	if f.Error != nil {
		return false, f.Error
	}

	return checkin.ID > 0, nil
}
