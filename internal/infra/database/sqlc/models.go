// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package postgres

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	UserRoleStudent    UserRole = "student"
	UserRoleInstructor UserRole = "instructor"
)

func (e *UserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRole(s)
	case string:
		*e = UserRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UserRole: %T", src)
	}
	return nil
}

type NullUserRole struct {
	UserRole UserRole
	Valid    bool // Valid is true if UserRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserRole) Scan(value interface{}) error {
	if value == nil {
		ns.UserRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserRole), nil
}

type Answer struct {
	AnswerID   uuid.UUID
	AuthorID   uuid.UUID
	QuestionID uuid.UUID
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Attachment struct {
	AttachmentID uuid.UUID
	QuestionID   uuid.NullUUID
	AnswerID     uuid.NullUUID
	Title        string
	Link         string
}

type Comment struct {
	CommentID  uuid.UUID
	AuthorID   uuid.UUID
	QuestionID uuid.NullUUID
	AnswerID   uuid.NullUUID
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Question struct {
	QuestionID   uuid.UUID
	AuthorID     uuid.UUID
	BestAnswerID uuid.NullUUID
	Slug         string
	Title        string
	Content      string
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type User struct {
	UserID uuid.UUID
	Name   string
	Role   UserRole
}
