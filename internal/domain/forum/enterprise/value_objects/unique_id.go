package value_objects

import (
	uuid "github.com/satori/go.uuid"
)

type UniqueID struct {
	Value *string
}

func NewUniqueID(value ...string) *UniqueID {
	uniqueID := UniqueID{}

	if len(value) > 0 {
		uniqueID.Value = &value[0]
	} else {
		generatedID := uuid.NewV4().String()
		uniqueID.Value = &generatedID
	}

	return &uniqueID
}

func (u *UniqueID) ToStringUniqueID() string {
	if u != nil {
		return *u.Value
	}

	return ""
}
