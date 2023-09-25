package repositories

import "github.com/intwone/ddd-golang/internal/domain/forum/enterprise"

type QuestionAttachmentsRepositoryInterface interface {
	GetManyByQuestionID(questionID string) (*[]enterprise.QuestionAttachment, error)
}
