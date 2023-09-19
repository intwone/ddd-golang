package enterprise

import (
	"strings"
	"time"

	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
)

type Question struct {
	id           *vo.UniqueID
	slug         *vo.Slug
	title        string
	content      string
	attachments  *[]QuestionAttachment
	bestAnswerID *vo.UniqueID
	authorID     *vo.UniqueID
	createdAt    time.Time
	updatedAt    *time.Time
}

type OptionalParams struct {
	ID          string
	Attachments []QuestionAttachment
}

func NewQuestion(title string, content string, authorID string, params ...OptionalParams) *Question {
	question := Question{
		title:     title,
		content:   content,
		authorID:  vo.NewUniqueID(authorID),
		createdAt: time.Now(),
	}

	slug := vo.NewSlug(title)
	question.slug = &vo.Slug{Value: slug.CreateFromText()}

	for _, param := range params {
		question.id = vo.NewUniqueID(param.ID)

		if len(param.Attachments) > 0 {
			question.attachments = &param.Attachments
		} else {
			question.attachments = &[]QuestionAttachment{}
		}
	}

	return &question
}

func (q *Question) GetID() string {
	return q.id.ToString()
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

func (q *Question) GetAttachments() []QuestionAttachment {
	return *q.attachments
}

func (q *Question) GetBestAnswerID() string {
	return q.bestAnswerID.ToString()
}

func (q *Question) GetAuthorID() string {
	return q.authorID.ToString()
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

func (q *Question) SetAttachments(attachments []QuestionAttachment) {
	q.attachments = &attachments
}

func (q *Question) SetBestAnswerID(bestAnswerID vo.UniqueID) {
	q.bestAnswerID = &bestAnswerID
	q.update()
}

func (q *Question) update() {
	now := time.Now()
	q.updatedAt = &now
}
