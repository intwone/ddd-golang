package entities

import (
	uuid "github.com/satori/go.uuid"
)

type Instructor struct {
	ID   *string
	Name string
}

func NewInstructor(name string, id ...string) *Instructor {
	instructor := &Instructor{
		Name: name,
	}

	if len(id) > 0 {
		instructor.ID = &id[0]
	} else {
		generatedId := uuid.NewV4().String()
		instructor.ID = &generatedId
	}

	return instructor
}
