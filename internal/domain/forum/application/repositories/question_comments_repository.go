package repositories

import "github.com/intwone/ddd-golang/internal/domain/forum/enterprise"

type QuestionCommentsRepositoryInterface interface {
	Create(questionComment *enterprise.QuestionComment) error
}
