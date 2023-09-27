package dtos

import "github.com/intwone/ddd-golang/internal/presentation/validations"

type CreateUserRequestDTO struct {
	Name string           `json:"name" binding:"required,min=3"`
	Role validations.Role `json:"role" binding:"required"`
}
