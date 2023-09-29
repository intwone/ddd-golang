package constants

// Database
const (
	NoRowsFound = "sql: no rows in result set"
)

// Email
const (
	InvalidEmailError             = "invalid email"
	EmailAlreadyTakenError        = "already taken"
	EmailOrPasswordIncorrectError = "email or password incorrect"
)

// UUID
const (
	InvalidUUIDError = "invalid uuid"
)

// Password
const (
	NotContainMinimumCaracteresPasswordError   = "must contain at least eight characters long"
	NotContainUpperCaseCharacterePasswordError = "must contain at least one uppercase character"
	NotContainSpecialCharacterePasswordError   = "must contain at least one special character"
	PasswordAreNotTheSame                      = "password are not the same"
)

// Role
const (
	InvalidRoleError = "must be student or instructor"
)

// Question
const (
	QuestionNotFoundError = "question not found"
)

// Other
const (
	NotAllowedError         = "not allowed"
	UnexpectedError         = "unexpected error"
	InvalidFieldTypeError   = "invalid field type"
	InvalidFieldsError      = "invalid fields"
	ConvertFieldsError      = "convert fields error"
	OccurredSameErrorsError = "ocurred same errors"
	GenerateHashError       = "error to generate hash"
	UnMarshalJSONError      = "unmarshal json error"
)
