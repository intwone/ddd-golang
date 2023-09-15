package enterprise

import (
	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
)

type QuestionComment struct {
	Comment
	questionID *vo.UniqueID
}

func NewQuestionComment(content string, questionID string) *QuestionComment {
	questionComment := QuestionComment{
		Comment:    *NewComment(content),
		questionID: vo.NewUniqueID(questionID),
	}

	return &questionComment
}

func (qc *QuestionComment) GetQuestionID() vo.UniqueID {
	return *qc.questionID
}
