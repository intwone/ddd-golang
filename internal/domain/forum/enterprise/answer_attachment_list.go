package enterprise

import "github.com/intwone/ddd-golang/internal/shared"

type AnswerAttachmentsList struct {
	*shared.WatchedList
}

func NewAnswerAttachmentsList(initialItems []interface{}) *AnswerAttachmentsList {
	return &AnswerAttachmentsList{
		WatchedList: shared.NewWatchedList(initialItems),
	}
}
