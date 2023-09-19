package enterprise

import (
	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
)

type QuestionAttachment struct {
	questionID   *vo.UniqueID
	attachmentID *vo.UniqueID
}

func NewQuestionAttachment(attachmentID string, questionID string) *QuestionAttachment {
	questionAttachment := QuestionAttachment{
		attachmentID: vo.NewUniqueID(attachmentID),
		questionID:   vo.NewUniqueID(questionID),
	}

	return &questionAttachment
}

func (ac *QuestionAttachment) GetQuestionID() vo.UniqueID {
	return *ac.questionID
}

func (ac *QuestionAttachment) GetAttachmentID() vo.UniqueID {
	return *ac.attachmentID
}
