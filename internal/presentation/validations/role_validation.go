package validations

type Role string

const (
	Student    Role = "student"
	Instructor Role = "instructor"
)

func (r Role) Validate() bool {
	switch r {
	case Student, Instructor:
		return true
	default:
		return false
	}
}
