package remind

import (
	"1536509937/ku-bbs/internal/consts"
	"1536509937/ku-bbs/internal/model"
	"fmt"
	"log"
)

type ReplyObs struct {
	Sender    uint64
	Receiver  uint64
	TopicID   uint64
	CommentId uint64
}

func (o *ReplyObs) Update() {
	var topic *model.Topics
	r := model.Topic().M.Where("id", o.TopicID).Find(&topic)
	if r.Error != nil || topic.ID <= 0 {
		log.Println(r.Error)
		return
	}

	if o.Sender == o.Receiver {
		return
	}

	r = model.Remind().M.Create(&model.Reminds{
		Sender:        o.Sender,
		Receiver:      o.Receiver,
		SourceId:      topic.ID,
		SourceContent: topic.Title,
		SourceType:    consts.TopicSource,
		SourceUrl:     fmt.Sprintf("/topics/%d?j=comment%d", o.TopicID, o.CommentId),
		Action:        consts.ReplyCommentRemind,
	})

	if r.Error != nil {
		log.Println(r.Error)
	}
}
