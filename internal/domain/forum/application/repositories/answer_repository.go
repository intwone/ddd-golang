package repositories

import "github.com/intwone/ddd-golang/internal/domain/forum/enterprise"

type AnswerRepositoryInterface interface {
	GetByID(id string) (enterprise.Answer, error)
	Create(answer *enterprise.Answer)
	DeleteByID(id string) error
}
