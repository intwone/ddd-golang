package dtos

import (
	"time"
)

type QuestionResponseDTO struct {
	ID           string     `json:"question_id"`
	AuthorID     string     `json:"author_id"`
	BestAnswerID string     `json:"best_answer_id,omitempty"`
	Slug         string     `json:"slug"`
	Title        string     `json:"title"`
	Content      string     `json:"content"`
	Attachments  []string   `json:"attachments,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}
