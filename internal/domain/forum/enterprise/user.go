package enterprise

import (
	"errors"
	"strings"
	"time"

	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
)

type User struct {
	id        *vo.UniqueID
	name      string
	email     *vo.Email
	password  *vo.Password
	role      string
	createdAt time.Time
	updatedAt *time.Time
}

func NewUser(name string, email string, password string, role string, id ...string) (*User, error) {
	e, err := vo.NewEmail(email)

	if e == nil {
		return nil, err
	}

	p, errs := vo.NewPassword(password)

	if p == nil {
		errs := strings.Join(errs, ",")
		return nil, errors.New(errs)
	}

	user := User{
		name:      name,
		role:      role,
		email:     e,
		password:  p,
		createdAt: time.Now(),
	}

	if len(id) > 0 {
		user.id = vo.NewUniqueID(id[0])
	} else {
		user.id = vo.NewUniqueID()
	}

	return &user, nil
}

func (s *User) GetID() string {
	return s.id.ToStringUniqueID()
}

func (s *User) GetName() string {
	return s.name
}

func (s *User) GetEmail() string {
	return s.email.Value
}

func (s *User) GetPassword() string {
	return s.password.ToStringPassword()
}

func (s *User) GetRole() string {
	return s.role
}

func (s *User) SetPassword(password string) {
	s.password = &vo.Password{Value: password}
}
