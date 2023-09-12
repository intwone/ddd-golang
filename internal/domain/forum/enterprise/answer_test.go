package enterprise_test

import (
	"strings"
	"testing"

	"github.com/intwone/ddd-golang/internal/domain/forum/enterprise"
	"github.com/stretchr/testify/require"
)

func TestSlug_CreateFromText(t *testing.T) {
	t.Run("should return a string with 120 caracters when answer content has more 117 caracters", func(t *testing.T) {
		answer := enterprise.NewAnswer(strings.Repeat("test", 150), "1", "1")
		excerpt := answer.GetExcerpt()

		require.Equal(t, 120, len(excerpt))
	})
}
