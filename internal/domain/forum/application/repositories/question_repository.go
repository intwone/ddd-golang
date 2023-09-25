package repositories

import "github.com/intwone/ddd-golang/internal/domain/forum/enterprise"

type QuestionRepositoryInterface interface {
	GetBySlug(slug string) (*enterprise.Question, error)
	GetByID(id string) (*enterprise.Question, error)
	GetManyRecent(page int64) (*[]enterprise.Question, error)
	Create(question *enterprise.Question) error
	Save(question *enterprise.Question) error
	DeleteByID(id string) error
}
