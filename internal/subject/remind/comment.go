package remind

import (
	"errors"
	"fmt"
	"1536509937/ku-bbs/internal/consts"
	"1536509937/ku-bbs/internal/model"
	"gorm.io/gorm"
	"log"
)

type CommentObs struct {
	Sender    uint64
	TopicID   uint64
	CommentId uint64
}

func (o *CommentObs) Update() {
	var topic model.Topics
	r := model.Topic().M.Where("id", o.TopicID).First(&topic)
	if r.Error != nil && !errors.Is(r.Error, gorm.ErrRecordNotFound) {
		log.Println(r.Error)
		return
	}

	if o.Sender == topic.UserId {
		return
	}

	sourceUrl := fmt.Sprintf("/topics/%d?j=comment%d", o.TopicID, o.CommentId)

	r = model.Remind().M.Create(&model.Reminds{
		Sender:        o.Sender,
		Receiver:      topic.UserId,
		SourceId:      topic.ID,
		SourceContent: topic.Title,
		SourceType:    model.Topic().Table,
		SourceUrl:     sourceUrl,
		Action:        consts.CommentTopicRemind,
	})

	if r.Error != nil {
		log.Println(r.Error)
	}
}
