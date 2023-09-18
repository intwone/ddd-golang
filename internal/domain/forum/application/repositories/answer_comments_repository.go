package repositories

import "github.com/intwone/ddd-golang/internal/domain/forum/enterprise"

type AnswerCommentsRepositoryInterface interface {
	GetByID(id string) (enterprise.AnswerComment, error)
	GetManyByID(page int64, id string) ([]enterprise.AnswerComment, error)
	Create(answerComment *enterprise.AnswerComment) error
	DeleteByID(id string) error
}
