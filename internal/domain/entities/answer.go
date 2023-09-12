package entities

import (
	uuid "github.com/satori/go.uuid"
)

type Answer struct {
	ID         *string
	Content    string
	AuthorID   string
	QuestionID string
}

func NewAnswer(content string, authorID string, questionID string, id ...string) *Answer {
	answer := &Answer{
		Content:    content,
		AuthorID:   authorID,
		QuestionID: questionID,
	}

	if len(id) > 0 {
		answer.ID = &id[0]
	} else {
		generatedId := uuid.NewV4().String()
		answer.ID = &generatedId
	}

	return answer
}
