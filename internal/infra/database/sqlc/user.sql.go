// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: user.sql

package postgres

import (
	"context"
	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :exec
insert into "users" (user_id, name, role) values ($1, $2, $3)
`

type CreateUserParams struct {
	UserID uuid.UUID
	Name   string
	Role   UserRole
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser, arg.UserID, arg.Name, arg.Role)
	return err
}
