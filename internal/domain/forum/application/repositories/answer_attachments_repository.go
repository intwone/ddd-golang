package repositories

import "github.com/intwone/ddd-golang/internal/domain/forum/enterprise"

type AnswerAttachmentsRepositoryInterface interface {
	GetManyByAnswerID(answerID string) ([]enterprise.AnswerAttachment, error)
}
