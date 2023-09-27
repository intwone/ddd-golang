package enterprise

import (
	"time"

	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
)

type User struct {
	id        *vo.UniqueID
	name      string
	role      string
	createdAt time.Time
	updatedAt *time.Time
}

func NewUser(name string, role string, id ...string) *User {
	user := User{
		name:      name,
		role:      role,
		createdAt: time.Now(),
	}

	if len(id) > 0 {
		user.id = vo.NewUniqueID(id[0])
	} else {
		user.id = vo.NewUniqueID()
	}

	return &user
}

func (s *User) GetID() string {
	return s.id.ToString()
}

func (s *User) GetName() string {
	return s.name
}

func (s *User) GetRole() string {
	return s.role
}
