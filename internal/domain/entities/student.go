package entities

import uuid "github.com/satori/go.uuid"

type Student struct {
	ID   *string
	Name string
}

func NewStudent(name string, id ...string) *Student {
	student := &Student{
		Name: name,
	}

	if len(id) > 0 {
		student.ID = &id[0]
	} else {
		generatedId := uuid.NewV4().String()
		student.ID = &generatedId
	}

	return student
}
