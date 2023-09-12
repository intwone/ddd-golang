package repositories

import "github.com/intwone/ddd-golang/internal/domain/entities"

type RepositoryInterface interface {
	Create(answer entities.Answer)
}
