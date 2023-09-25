package factories

import (
	"strconv"

	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

func AnswerAttachmentsFactory(count int, answerID string) *[]enterprise.AnswerAttachment {
	answerAttachments := make([]enterprise.AnswerAttachment, count)

	for i := 0; i < count; i++ {
		attachmentID := strconv.Itoa(i + 1)
		attachment := enterprise.NewAnswerAttachment(attachmentID, answerID)
		answerAttachments[i] = *attachment
	}

	return &answerAttachments
}
