package entities

import vo "github.com/intwone/ddd-golang/internal/domain/entities/value_objects"

type Student struct {
	id   *vo.UniqueID
	name string
}

func NewStudent(name string, id ...string) *Student {
	student := Student{
		name: name,
	}

	if len(id) > 0 {
		student.id = vo.NewUniqueID(id[0])
	} else {
		student.id = vo.NewUniqueID()
	}

	return &student
}

func (s *Student) GetName() string {
	return s.name
}
