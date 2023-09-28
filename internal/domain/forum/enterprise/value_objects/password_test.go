package value_objects_test

import (
	"testing"

	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
	"github.com/stretchr/testify/require"
)

func TestPassword_IsValid(t *testing.T) {
	t.Run("should return true when password is valid", func(t *testing.T) {
		password := "Test@123"
		result, _ := vo.IsValidPassword(password)

		require.Equal(t, true, result)
	})

	t.Run("should return false when password not contains at least 8 characters", func(t *testing.T) {
		password := "Test@12"
		result, errs := vo.IsValidPassword(password)

		require.Equal(t, false, result)
		require.Equal(t, errs[0], "the password must be at least 8 characters long")
	})

	t.Run("should return false when password not contains at least 8 characters and not contains at least 1 uppercase caractere", func(t *testing.T) {
		password := "test@12"
		result, errs := vo.IsValidPassword(password)

		require.Equal(t, false, result)
		require.Equal(t, errs[0], "the password must be at least 8 characters long")
		require.Equal(t, errs[1], "the password must contain at least 1 uppercase character")
	})

	t.Run("should return false when password not contains at least 8 characters and not contains at least 1 uppercase caractere and not contain at least one special character", func(t *testing.T) {
		password := "test12"
		result, errs := vo.IsValidPassword(password)

		require.Equal(t, false, result)
		require.Equal(t, errs[0], "the password must contain at least eight characters long")
		require.Equal(t, errs[1], "the password must contain at least one uppercase character")
		require.Equal(t, errs[2], "the password must contain at least one special character")
	})
}
