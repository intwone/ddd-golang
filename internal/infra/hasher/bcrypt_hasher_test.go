package hasher_test

import (
	"testing"

	"github.com/intwone/ddd-golang/internal/infra/hasher"
	"github.com/stretchr/testify/require"
)

func TestBcryptHasher_Hash(t *testing.T) {
	t.Run("should hash a value", func(t *testing.T) {
		value := "test123"
		_, err := hasher.NewBcryptHasher().Hash(value)

		require.Nil(t, err)
	})

	t.Run("should return true when the comparison of values ​​is true", func(t *testing.T) {
		value := "test123"
		valueHashed, _ := hasher.NewBcryptHasher().Hash(value)

		result := hasher.NewBcryptHasher().Compare(value, *valueHashed)

		require.Equal(t, true, result)
	})

	t.Run("should return false when the comparison of values ​​is false", func(t *testing.T) {
		value1 := "test123"
		value2 := "test111"
		valueHashed, _ := hasher.NewBcryptHasher().Hash(value1)

		result := hasher.NewBcryptHasher().Compare(value2, *valueHashed)

		require.Equal(t, false, result)
	})
}
