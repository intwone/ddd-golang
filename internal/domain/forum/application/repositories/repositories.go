package repositories

import "github.com/intwone/ddd-golang/internal/domain/forum/enterprise"

type RepositoryInterface interface {
	Create(answer *enterprise.Answer)
}
