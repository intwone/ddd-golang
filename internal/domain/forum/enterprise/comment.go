package enterprise

import (
	"time"

	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
)

type Comment struct {
	id        *vo.UniqueID
	authorID  *vo.UniqueID
	content   string
	createdAt time.Time
	updatedAt *time.Time
}

func NewComment(content string, authorID string, id ...string) *Comment {
	comment := Comment{
		authorID:  vo.NewUniqueID(authorID),
		content:   content,
		createdAt: time.Now(),
	}

	if len(id) > 0 {
		comment.id = vo.NewUniqueID(id[0])
	} else {
		comment.id = vo.NewUniqueID()
	}

	return &comment
}

func (c *Comment) GetID() vo.UniqueID {
	return *c.id
}

func (c *Comment) GetAuthorID() vo.UniqueID {
	return *c.authorID
}

func (c *Comment) GetContent() string {
	return c.content
}

func (c *Comment) GetCreatedAt() time.Time {
	return c.createdAt
}

func (c *Comment) GetUpdatedAt() time.Time {
	return *c.updatedAt
}

func (c *Comment) SetContent(content string) {
	c.content = content
	c.update()
}

func (c *Comment) update() {
	now := time.Now()
	c.updatedAt = &now
}
