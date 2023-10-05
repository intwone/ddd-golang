package dtos

import "time"

type AnswerResponseDTO struct {
	ID          string     `json:"id"`
	AuthorID    string     `json:"author_id"`
	QuestionID  string     `json:"question_id"`
	Content     string     `json:"content"`
	Attachments []string   `json:"attachments,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
