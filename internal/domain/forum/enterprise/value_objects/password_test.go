package value_objects_test

import (
	"testing"

	"github.com/intwone/ddd-golang/internal/constants"
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
		require.Equal(t, errs[0].Error(), constants.NotContainMinimumCaracteresPasswordError)
	})

	t.Run("should return false when password not contains at least 8 characters and not contains at least 1 uppercase caractere", func(t *testing.T) {
		password := "test@12"
		result, errs := vo.IsValidPassword(password)

		require.Equal(t, false, result)
		require.Equal(t, errs[0].Error(), constants.NotContainMinimumCaracteresPasswordError)
		require.Equal(t, errs[1].Error(), constants.NotContainUpperCaseCharacterePasswordError)
	})

	t.Run("should return false when password not contains at least 8 characters and not contains at least 1 uppercase caractere and not contain at least one special character", func(t *testing.T) {
		password := "test12"
		result, errs := vo.IsValidPassword(password)

		require.Equal(t, false, result)
		require.Equal(t, errs[0].Error(), constants.NotContainMinimumCaracteresPasswordError)
		require.Equal(t, errs[1].Error(), constants.NotContainUpperCaseCharacterePasswordError)
		require.Equal(t, errs[2].Error(), constants.NotContainSpecialCharacterePasswordError)
	})
}
