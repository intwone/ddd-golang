package repositories

import "github.com/intwone/ddd-golang/internal/domain/forum/enterprise"

type UserRepositoryInterface interface {
	GetByEmail(email string) (*enterprise.User, error)
	Create(user *enterprise.User) error
}
