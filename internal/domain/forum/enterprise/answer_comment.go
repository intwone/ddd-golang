package enterprise

import (
	"time"

	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
)

type AnswerComment struct {
	id        *vo.UniqueID
	content   string
	answerID  *vo.UniqueID
	authorID  *vo.UniqueID
	createdAt time.Time
	updatedAt *time.Time
}

func NewAnswerComment(content string, answerID string, authorID string, id ...string) *AnswerComment {
	answerComment := AnswerComment{
		content:   content,
		answerID:  vo.NewUniqueID(answerID),
		authorID:  vo.NewUniqueID(authorID),
		createdAt: time.Now(),
	}

	if len(id) > 0 {
		answerComment.id = vo.NewUniqueID(id[0])
	} else {
		answerComment.id = vo.NewUniqueID()
	}

	return &answerComment
}

func (ac *AnswerComment) GetID() vo.UniqueID {
	return *ac.id
}

func (ac *AnswerComment) GetContent() string {
	return ac.content
}

func (ac *AnswerComment) GetAnswerID() vo.UniqueID {
	return *ac.answerID
}

func (ac *AnswerComment) GetAuthorID() vo.UniqueID {
	return *ac.authorID
}

func (ac *AnswerComment) SetContent(content string) {
	ac.content = content
	ac.update()
}

func (ac *AnswerComment) update() {
	now := time.Now()
	ac.updatedAt = &now
}
