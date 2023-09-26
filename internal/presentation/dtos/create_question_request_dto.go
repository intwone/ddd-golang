package dtos

type CreateQuestionRequestDTO struct {
	AuthorID string `json:"author_id" binding:"required,uuid"`
	Title    string `json:"title" binding:"required,min=1,max=100"`
	Content  string `json:"content" binding:"required,min=1"`
}
