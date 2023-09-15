package enterprise

import (
	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
)

type AnswerComment struct {
	Comment
	answerID *vo.UniqueID
}

func NewAnswerComment(content string, answerID string) *AnswerComment {
	answerComment := AnswerComment{
		Comment:  *NewComment(content),
		answerID: vo.NewUniqueID(answerID),
	}

	return &answerComment
}

func (ac *AnswerComment) GetAnswerID() vo.UniqueID {
	return *ac.answerID
}
