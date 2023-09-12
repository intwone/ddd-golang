package value_objects

import "github.com/gosimple/slug"

type Slug struct {
	Value string
}

func NewSlug(value string) *Slug {
	slug := Slug{
		Value: value,
	}

	return &slug
}

func (s *Slug) CreateFromText() string {
	formattedText := slug.Make(s.Value)
	return formattedText
}
