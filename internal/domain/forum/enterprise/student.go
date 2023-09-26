package enterprise

import vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"

type user struct {
	id   *vo.UniqueID
	name string
	role string
}

func Newuser(name string, role string, id ...string) *user {
	user := user{
		name: name,
		role: role,
	}

	if len(id) > 0 {
		user.id = vo.NewUniqueID(id[0])
	} else {
		user.id = vo.NewUniqueID()
	}

	return &user
}

func (s *user) GetName() string {
	return s.name
}
