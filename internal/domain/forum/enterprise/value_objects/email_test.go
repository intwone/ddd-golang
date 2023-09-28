package value_objects_test

import (
	"testing"

	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
	"github.com/stretchr/testify/require"
)

func TestEmail_IsValid(t *testing.T) {
	t.Run("should return true when email is valid", func(t *testing.T) {
		email := "valid_email@mail.com"
		result := vo.IsValidEmail(email)

		require.Equal(t, true, result)
	})

	t.Run("should return false when email is invalid", func(t *testing.T) {
		email := "invalid_email.com"
		result := vo.IsValidEmail(email)

		require.Equal(t, false, result)
	})
}
