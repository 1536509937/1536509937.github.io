package remind

import (
	"1536509937/ku-bbs/internal/consts"
	"1536509937/ku-bbs/internal/model"
	"log"
)

type FollowObs struct {
	Sender   uint64
	Receiver uint64
}

func (o *FollowObs) Update() {
	r := model.Remind().M.Create(&model.Reminds{
		Sender:        o.Sender,
		Receiver:      o.Receiver,
		SourceId:      0,
		SourceContent: "",
		SourceType:    consts.UserSource,
		SourceUrl:     "",
		Action:        consts.FollowUserRemind,
	})
	if r.Error != nil {
		log.Println(r.Error)
	}
}
