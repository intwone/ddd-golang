package repositories

import "github.com/intwone/ddd-golang/internal/domain/forum/enterprise"

type AnswerRepositoryInterface interface {
	Create(answer *enterprise.Answer)
}
