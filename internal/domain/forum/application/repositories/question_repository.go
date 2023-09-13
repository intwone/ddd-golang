package repositories

import "github.com/intwone/ddd-golang/internal/domain/forum/enterprise"

type QuestionRepositoryInterface interface {
	Create(question *enterprise.Question)
	GetBySlug(slug string) (enterprise.Question, error)
}
