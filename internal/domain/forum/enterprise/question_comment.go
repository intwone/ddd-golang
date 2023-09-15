package enterprise

import (
	"time"

	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
)

type QuestionComment struct {
	id         *vo.UniqueID
	content    string
	questionID *vo.UniqueID
	authorID   *vo.UniqueID
	createdAt  time.Time
	updatedAt  *time.Time
}

func NewQuestionComment(content string, questionID string, authorID string, id ...string) *QuestionComment {
	questionComment := QuestionComment{
		content:    content,
		questionID: vo.NewUniqueID(questionID),
		authorID:   vo.NewUniqueID(authorID),
		createdAt:  time.Now(),
	}

	if len(id) > 0 {
		questionComment.id = vo.NewUniqueID(id[0])
	} else {
		questionComment.id = vo.NewUniqueID()
	}

	return &questionComment
}

func (qc *QuestionComment) GetID() vo.UniqueID {
	return *qc.id
}

func (qc *QuestionComment) GetContent() string {
	return qc.content
}

func (qc *QuestionComment) GetQuestionID() vo.UniqueID {
	return *qc.questionID
}

func (qc *QuestionComment) GetAuthorID() vo.UniqueID {
	return *qc.authorID
}

func (qc *QuestionComment) SetContent(content string) {
	qc.content = content
	qc.update()
}

func (qc *QuestionComment) update() {
	now := time.Now()
	qc.updatedAt = &now
}
