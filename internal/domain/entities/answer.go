package entities

import (
	"time"

	vo "github.com/intwone/ddd-golang/internal/domain/entities/value_objects"
)

type Answer struct {
	id         *vo.UniqueID
	content    string
	authorID   *vo.UniqueID
	questionID *vo.UniqueID
	createdAt  time.Time
	updatedAt  *time.Time
}

func NewAnswer(content string, authorID string, questionID string, id ...string) *Answer {
	answer := Answer{
		content:    content,
		authorID:   vo.NewUniqueID(authorID),
		questionID: vo.NewUniqueID(questionID),
		createdAt:  time.Now(),
	}

	if len(id) > 0 {
		answer.id = vo.NewUniqueID(id[0])
	} else {
		answer.id = vo.NewUniqueID()
	}

	return &answer
}

func (a *Answer) GetContent() string {
	return a.content
}

func (a *Answer) GetAuthorID() vo.UniqueID {
	return *a.authorID
}

func (a *Answer) GetQuestionID() vo.UniqueID {
	return *a.questionID
}
