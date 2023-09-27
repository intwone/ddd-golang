package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/intwone/ddd-golang/internal/domain/forum/application/repositories"
	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
	s "github.com/intwone/ddd-golang/internal/infra/database/sqlc"
)

type UserSQLCRepository struct {
	db *s.Queries
}

func NewUserSQLCRepository(db *s.Queries) repositories.UserRepositoryInterface {
	return &UserSQLCRepository{
		db: db,
	}
}

func (r *UserSQLCRepository) Create(user *enterprise.User) error {
	userID, err := uuid.Parse(user.GetID())

	if err != nil {
		return err
	}

	createUserErr := r.db.CreateUser(context.Background(), s.CreateUserParams{
		UserID: userID,
		Name:   user.GetName(),
		Role:   s.UserRole(user.GetRole()),
	})

	if createUserErr != nil {
		return createUserErr
	}

	return nil
}
