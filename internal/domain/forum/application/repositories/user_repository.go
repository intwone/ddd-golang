package repositories

import "github.com/intwone/ddd-golang/internal/domain/forum/enterprise"

type UserRepositoryInterface interface {
	Create(user *enterprise.User) error
}
