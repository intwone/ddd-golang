package entities

import (
	uuid "github.com/satori/go.uuid"
)

type Answer struct {
	ID      *string
	Content string
}

func NewAnswer(content string, id ...string) *Answer {
	answer := &Answer{
		Content: content,
	}

	if len(id) > 0 {
		answer.ID = &id[0]
	} else {
		generatedId := uuid.NewV4().String()
		answer.ID = &generatedId
	}

	return answer
}
