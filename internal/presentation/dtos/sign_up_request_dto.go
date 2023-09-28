package dtos

import "github.com/intwone/ddd-golang/internal/presentation/validations"

type SignUpRequestDTO struct {
	Name     string           `json:"name" binding:"required,min=3"`
	Email    string           `json:"email" binding:"required,email"`
	Password string           `json:"password" binding:"required"`
	Role     validations.Role `json:"role" binding:"required"`
}
