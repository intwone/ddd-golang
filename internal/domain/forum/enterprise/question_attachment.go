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
		questionID:   vo.NewUniqueID(questionID),
		attachmentID: vo.NewUniqueID(attachmentID),
	}

	return &questionAttachment
}

func (ac *QuestionAttachment) GetQuestionID() string {
	return ac.questionID.ToStringUniqueID()
}

func (ac *QuestionAttachment) GetAttachmentID() string {
	return ac.attachmentID.ToStringUniqueID()
}

func (ac *QuestionAttachment) SetAttachmentID(attachmentID vo.UniqueID) {
	ac.attachmentID = &attachmentID
}
