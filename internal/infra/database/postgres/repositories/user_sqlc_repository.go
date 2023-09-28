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

func (r *UserSQLCRepository) GetByEmail(email string) (*enterprise.User, error) {
	result, err := r.db.GetUserByEmail(context.Background(), email)

	if err != nil {
		return nil, err
	}

	user, err := enterprise.NewUser(
		result.Name,
		result.Email,
		result.Password,
		string(result.Role),
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserSQLCRepository) Create(user *enterprise.User) error {
	userID, err := uuid.Parse(user.GetID())

	if err != nil {
		return err
	}

	createUserErr := r.db.CreateUser(context.Background(), s.CreateUserParams{
		UserID:   userID,
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
		Role:     s.UserRole(user.GetRole()),
	})

	if createUserErr != nil {
		return createUserErr
	}

	return nil
}
