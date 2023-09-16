package repositories

import "github.com/intwone/ddd-golang/internal/domain/forum/enterprise"

type AnswerCommentsRepositoryInterface interface {
	Create(answerComment *enterprise.AnswerComment) error
}
