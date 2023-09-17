package repositories

import "github.com/intwone/ddd-golang/internal/domain/forum/enterprise"

type QuestionCommentsRepositoryInterface interface {
	GetByID(id string) (enterprise.QuestionComment, error)
	Create(questionComment *enterprise.QuestionComment) error
	DeleteByID(id string) error
}
