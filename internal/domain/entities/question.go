package entities

import (
	"time"

	vo "github.com/intwone/ddd-golang/internal/domain/entities/value_objects"
)

type Question struct {
	id           *vo.UniqueID
	slug         vo.Slug
	title        string
	content      string
	bestAnswerID *vo.UniqueID
	authorID     *vo.UniqueID
	createdAt    time.Time
	updatedAt    *time.Time
}

func NewQuestion(slug vo.Slug, title string, content string, authorId string, id ...string) *Question {
	question := Question{
		slug:      slug,
		title:     title,
		content:   content,
		authorID:  vo.NewUniqueID(authorId),
		createdAt: time.Now(),
	}

	if len(id) > 0 {
		question.id = vo.NewUniqueID(id[0])
	} else {
		question.id = vo.NewUniqueID()
	}

	return &question
}

func (q *Question) GetSlug() vo.Slug {
	return q.slug
}

func (q *Question) GetTitle() string {
	return q.title
}

func (q *Question) GetContent() string {
	return q.content
}

func (q *Question) GetBestAnswerID() vo.UniqueID {
	return *q.bestAnswerID
}

func (q *Question) GetAuthorID() vo.UniqueID {
	return *q.authorID
}
