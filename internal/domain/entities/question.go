package entities

import uuid "github.com/satori/go.uuid"

type Question struct {
	ID       *string
	Title    string
	Content  string
	AuthorID string
}

func NewQuestion(title string, content string, authorId string, id ...string) *Question {
	question := &Question{
		Title:    title,
		Content:  content,
		AuthorID: authorId,
	}

	if len(id) > 0 {
		question.ID = &id[0]
	} else {
		generatedId := uuid.NewV4().String()
		question.ID = &generatedId
	}

	return question
}
