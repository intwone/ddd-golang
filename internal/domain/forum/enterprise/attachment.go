package enterprise

import (
	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
)

type Attachment struct {
	id    *vo.UniqueID
	title string
	link  string
}

func NewAttachment(title string, link string, id ...string) *Attachment {
	attachment := Attachment{
		title: title,
		link:  link,
	}

	if len(id) > 0 {
		attachment.id = vo.NewUniqueID(id[0])
	} else {
		attachment.id = vo.NewUniqueID()
	}

	return &attachment
}

func (a *Attachment) GetID() string {
	return a.id.ToString()
}

func (a *Attachment) GetTitle() string {
	return a.title
}

func (a *Attachment) GetLink() string {
	return a.link
}
