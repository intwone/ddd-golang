package value_objects_test

import (
	"testing"

	"github.com/google/uuid"
	vo "github.com/intwone/ddd-golang/internal/domain/forum/enterprise/value_objects"
	"github.com/stretchr/testify/require"
)

func TestUnique_ID_ToString(t *testing.T) {
	t.Run("should create an uuid", func(t *testing.T) {
		id := vo.NewUniqueID()
		idToString := id.ToString()

		_, err := uuid.Parse(idToString)

		require.Nil(t, err)
	})
}
