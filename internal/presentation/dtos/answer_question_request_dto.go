package dtos

type AnswerQuestionRequestDTO struct {
	AuthorID   string `json:"author_id" binding:"required,uuid"`
	QuestionID string `json:"question_id" binding:"required,uuid"`
	Content    string `json:"content" binding:"required,min=1"`
}
