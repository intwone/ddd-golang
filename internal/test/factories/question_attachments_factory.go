package factories

import (
	"strconv"

	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
)

func QuestionAttachmentsFactory(count int, questionID string) []enterprise.QuestionAttachment {
	questionAttachments := make([]enterprise.QuestionAttachment, count)

	for i := 0; i < count; i++ {
		attachmentID := strconv.Itoa(i + 1)
		attachment := enterprise.NewQuestionAttachment(attachmentID, questionID)
		questionAttachments[i] = *attachment
	}

	return questionAttachments
}
