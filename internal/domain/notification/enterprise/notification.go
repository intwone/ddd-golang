package enterprise

import (
	"time"

	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
)

type Notification struct {
	id          *vo.UniqueID
	recipientID *vo.UniqueID
	title       string
	content     string
	createdAt   time.Time
	readAt      *time.Time
}

type NotificationOptionalParams struct {
	ID     string
	ReadAt *time.Time
}

func NewNotification(title string, content string, recipientID string, params ...NotificationOptionalParams) *Notification {
	notification := Notification{
		recipientID: vo.NewUniqueID(recipientID),
		title:       title,
		content:     content,
		createdAt:   time.Now(),
	}

	for _, param := range params {
		if param.ID != "" {
			notification.id = vo.NewUniqueID(param.ID)
		}
	}

	if notification.id == nil {
		notification.id = vo.NewUniqueID()
	}

	return &notification
}

func (n *Notification) GetID() string {
	return n.id.ToString()
}

func (n *Notification) GetRecipientID() string {
	return n.recipientID.ToString()
}

func (n *Notification) GeTitle() string {
	return n.title
}

func (n *Notification) GetContent() string {
	return n.content
}

func (n *Notification) SetReadAt() time.Time {
	return *n.readAt
}
