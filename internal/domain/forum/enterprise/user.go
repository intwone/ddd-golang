package enterprise

import (
	"time"

	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
)

type User struct {
	id        *vo.UniqueID
	name      string
	email     *vo.Email
	password  *vo.Password
	role      *vo.Role
	createdAt time.Time
	updatedAt *time.Time
}

func NewUser(name string, email string, password string, role string, id ...string) (*User, []error) {
	e, newEmailErr := vo.NewEmail(email)

	if e == nil {
		errors := []error{newEmailErr}
		return nil, errors
	}

	p, newPasswordErrs := vo.NewPassword(password)

	if p == nil {
		return nil, newPasswordErrs
	}

	r, newRoleErr := vo.NewRole(role)

	if r == nil {
		errors := []error{newRoleErr}
		return nil, errors
	}

	user := User{
		name:      name,
		role:      r,
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
	return s.role.Value
}

func (s *User) SetPassword(password string) {
	s.password = &vo.Password{Value: password}
}
