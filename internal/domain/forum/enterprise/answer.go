package enterprise

import (
	"strings"
	"time"

	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
)

type Answer struct {
	id          *vo.UniqueID
	content     string
	attachments *AnswerAttachmentsList
	authorID    *vo.UniqueID
	questionID  *vo.UniqueID
	createdAt   time.Time
	updatedAt   *time.Time
}

type AnswerOptionalParams struct {
	ID          string
	Attachments AnswerAttachmentsList
}

func NewAnswer(content string, authorID string, questionID string, params ...AnswerOptionalParams) *Answer {
	answer := Answer{
		content:    content,
		authorID:   vo.NewUniqueID(authorID),
		questionID: vo.NewUniqueID(questionID),
		createdAt:  time.Now(),
	}

	for _, param := range params {
		if param.ID != "" {
			answer.id = vo.NewUniqueID(param.ID)
		}

		if len(param.Attachments.GetCurrentItems()) > 0 {
			answer.attachments = &param.Attachments
		} else {
			answer.attachments = NewAnswerAttachmentsList([]interface{}{})
		}
	}

	if answer.id == nil {
		answer.id = vo.NewUniqueID()
	}

	return &answer
}

func (a *Answer) GetID() string {
	return a.id.ToStringUniqueID()
}

func (a *Answer) GetContent() string {
	return a.content
}

func (a *Answer) GetAttachments() AnswerAttachmentsList {
	return *a.attachments
}

func (a *Answer) GetAuthorID() string {
	return a.authorID.ToStringUniqueID()
}

func (a *Answer) GetQuestionID() string {
	return a.questionID.ToStringUniqueID()
}

func (a *Answer) GetExcerpt() string {
	maxLength := 117

	if len(a.content) > maxLength {
		return strings.TrimRight(a.content[:maxLength], " ") + "..."
	}

	return a.content
}

func (a *Answer) SetContent(content string) {
	a.content = content
	a.update()
}

func (a *Answer) SetAttachments(attachments AnswerAttachmentsList) {
	a.attachments = &attachments
	a.update()
}

func (a *Answer) CanModify(authorID string) bool {
	return a.authorID.ToStringUniqueID() == authorID
}

func (a *Answer) update() {
	now := time.Now()
	a.updatedAt = &now
}
