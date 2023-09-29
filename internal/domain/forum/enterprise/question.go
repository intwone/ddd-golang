package enterprise

import (
	"strings"
	"time"

	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
)

type Question struct {
	id           *vo.UniqueID
	authorID     *vo.UniqueID
	bestAnswerID *vo.UniqueID
	slug         *vo.Slug
	title        string
	content      string
	attachments  *QuestionAttachmentsList
	createdAt    time.Time
	updatedAt    *time.Time
}

type QuestionOptionalParams struct {
	ID           string
	BestAnswerID string
	Attachments  QuestionAttachmentsList
}

func NewQuestion(title string, content string, authorID string, params ...QuestionOptionalParams) *Question {
	question := Question{
		authorID:  vo.NewUniqueID(authorID),
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}

	slug := vo.NewSlug(title)
	question.slug = &vo.Slug{Value: slug.CreateFromText()}

	for _, param := range params {
		if param.ID != "" {
			question.id = vo.NewUniqueID(param.ID)
		}

		if param.BestAnswerID != "" {
			question.bestAnswerID = vo.NewUniqueID(param.BestAnswerID)
		}

		if len(param.Attachments.GetCurrentItems()) > 0 {
			question.attachments = &param.Attachments
		} else {
			question.attachments = NewQuestionAttachmentsList([]interface{}{})
		}
	}

	if question.id == nil {
		question.id = vo.NewUniqueID()
	}

	return &question
}

func (q *Question) GetID() string {
	return q.id.ToStringUniqueID()
}

func (q *Question) GetAuthorID() string {
	return q.authorID.ToStringUniqueID()
}

func (q *Question) GetSlug() vo.Slug {
	return *q.slug
}

func (q *Question) GetTitle() string {
	return q.title
}

func (q *Question) GetContent() string {
	return q.content
}

func (q *Question) GetAttachments() QuestionAttachmentsList {
	return *q.attachments
}

func (q *Question) GetBestAnswerID() string {
	return q.bestAnswerID.ToStringUniqueID()
}

func (q *Question) GetCreatedAt() *time.Time {
	return &q.createdAt
}

func (q *Question) GetUpdatedAt() *time.Time {
	return q.updatedAt
}

func (q *Question) GetExcerpt() string {
	maxLength := 117

	if len(q.content) > maxLength {
		return strings.TrimRight(q.content[:maxLength], " ") + "..."
	}
	return q.content
}

func (q *Question) SetTitle(title string) {
	q.title = title
	q.update()
}

func (q *Question) SetContent(content string) {
	q.content = content
	q.update()
}

func (q *Question) SetAttachments(attachments QuestionAttachmentsList) {
	q.attachments = &attachments
	q.update()
}

func (q *Question) SetBestAnswerID(bestAnswerID vo.UniqueID) {
	q.bestAnswerID = &bestAnswerID
	q.update()
}

func (q *Question) CanModify(authorID string) bool {
	return q.authorID.ToStringUniqueID() == authorID
}

func (q *Question) update() {
	now := time.Now()
	q.updatedAt = &now
}
