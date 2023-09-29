package value_objects_test

import (
	"testing"

	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
	"github.com/stretchr/testify/require"
)

func TestRole_IsValid(t *testing.T) {
	t.Run("should return true when role is valid", func(t *testing.T) {
		role1 := "student"
		result1 := vo.IsValidRole(role1)

		role2 := "instructor"
		result2 := vo.IsValidRole(role2)

		require.Equal(t, true, result1)
		require.Equal(t, true, result2)
	})

	t.Run("should return false when role is invalid", func(t *testing.T) {
		role := "invalid_role"
		result := vo.IsValidRole(role)

		require.Equal(t, false, result)
	})
}
