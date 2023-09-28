package constants

// Database
const (
	NoRowsFound = "no rows in result set"
)

// Email
const (
	InvalidEmailError      = "invalid email"
	EmailAlreadyTakenError = "email already taken"
)

// UUiD
const (
	InvalidUUIDError = "invalid uuid"
)

// Password
const (
	NotContainMinimumCaracteresPasswordError   = "the password must contain at least eight characters long"
	NotContainUpperCaseCharacterePasswordError = "the password must contain at least one uppercase character"
	NotContainSpecialCharacterePasswordError   = "the password must contain at least one special character"
)

// Role
const (
	InvalidRoleError = "role must be student or instructor"
)

// Question
const (
	QuestionNotFoundError = "question not found"
)

// Other
const (
	NotAllowedError = "not allowed"
)
