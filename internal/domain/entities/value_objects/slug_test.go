package value_objects_test

import (
	"testing"

	valueobjects "github.com/intwone/ddd-golang/internal/domain/entities/value_objects"
	"github.com/stretchr/testify/require"
)

func TestSlug_CreateFromText(t *testing.T) {
	t.Run("should create a slug", func(t *testing.T) {
		slug := valueobjects.NewSlug("An Exemple TiTle--")
		slugFormatted := slug.CreateFromText()

		require.Equal(t, "an-exemple-title", slugFormatted)
	})
}
