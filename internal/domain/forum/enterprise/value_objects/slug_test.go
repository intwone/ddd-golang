package value_objects_test

import (
	"testing"

	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
	"github.com/stretchr/testify/require"
)

func TestSlug_CreateFromText(t *testing.T) {
	t.Run("should create a slug", func(t *testing.T) {
		slug := vo.NewSlug("An Exemple TiTle--")
		slugFormatted := slug.CreateFromText()

		require.Equal(t, "an-exemple-title", slugFormatted)
	})
}
