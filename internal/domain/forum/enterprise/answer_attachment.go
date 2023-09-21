package enterprise

import (
	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
)

type AnswerAttachment struct {
	answerID     *vo.UniqueID
	attachmentID *vo.UniqueID
}

func NewAnswerAttachment(attachmentID string, answerID string) *AnswerAttachment {
	answerAttachment := AnswerAttachment{
		attachmentID: vo.NewUniqueID(attachmentID),
		answerID:     vo.NewUniqueID(answerID),
	}

	return &answerAttachment
}

func (ac *AnswerAttachment) GetAnswerID() string {
	return ac.answerID.ToString()
}

func (ac *AnswerAttachment) GetAttachmentID() string {
	return ac.attachmentID.ToString()
}

func (ac *AnswerAttachment) SetAttachmentID(attachmentID vo.UniqueID) {
	ac.attachmentID = &attachmentID
}
