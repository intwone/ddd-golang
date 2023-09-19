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

func (ac *QuestionAttachment) GetQuestionID() string {
	return ac.questionID.ToString()
}

func (ac *QuestionAttachment) GetAttachmentID() string {
	return ac.attachmentID.ToString()
}
