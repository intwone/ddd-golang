package enterprise

import (
	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
)

type Instructor struct {
	id   *vo.UniqueID
	name string
}

func NewInstructor(name string, id ...string) *Instructor {
	instructor := Instructor{
		name: name,
	}

	if len(id) > 0 {
		instructor.id = vo.NewUniqueID(id[0])
	} else {
		instructor.id = vo.NewUniqueID()
	}

	return &instructor
}

func (i *Instructor) GetName() string {
	return i.name
}
