package dtos

type CreateQuestionRequestDTO struct {
	Title   string `json:"title" binding:"required,min=1,max=100"`
	Content string `json:"content" binding:"required,min=1"`
}
