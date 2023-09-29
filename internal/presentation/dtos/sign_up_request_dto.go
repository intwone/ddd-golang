package dtos

type SignUpRequestDTO struct {
	Name     string `json:"name" binding:"required,min=3"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}
