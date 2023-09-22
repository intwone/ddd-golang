package enterprise

import "github.com/intwone/ddd-golang/internal/domain/shared"

type QuestionAttachmentsList struct {
	*shared.WatchedList
}

func NewQuestionAttachmentsList(initialItems []interface{}) *QuestionAttachmentsList {
	return &QuestionAttachmentsList{
		WatchedList: shared.NewWatchedList(initialItems),
	}
}
