package shared_test

import (
	"testing"

	"github.com/intwone/ddd-golang/internal/domain/forum/shared"
	"github.com/stretchr/testify/require"
)

func TestWatchedList_GetCurrentItems(t *testing.T) {
	watchedList := shared.NewWatchedList([]interface{}{1, 2, 3, 4})

	require.Equal(t, 4, len(watchedList.GetCurrentItems()))
}

func TestWatchedList_Add(t *testing.T) {
	watchedList := shared.NewWatchedList([]interface{}{1, 2, 3, 4})
	watchedList.Add(5)

	require.Equal(t, 5, len(watchedList.GetCurrentItems()))
	require.Equal(t, []interface{}{5}, watchedList.GetNewItems())
}

func TestWatchedList_Remove(t *testing.T) {
	watchedList := shared.NewWatchedList([]interface{}{1, 2, 3, 4})
	watchedList.Remove(2)

	require.Equal(t, 3, len(watchedList.GetCurrentItems()))
	require.Equal(t, []interface{}{2}, watchedList.GetRemovedItems())
}

func TestWatchedList(t *testing.T) {
	t.Run("should be anle to add an item event if it was removed befere", func(t *testing.T) {
		watchedList := shared.NewWatchedList([]interface{}{1, 2, 3, 4})
		watchedList.Remove(2)
		watchedList.Add(2)

		require.Equal(t, 0, len(watchedList.GetRemovedItems()))
		require.Equal(t, 0, len(watchedList.GetNewItems()))
	})

	t.Run("should be able to remove an item even if it was added before", func(t *testing.T) {
		watchedList := shared.NewWatchedList([]interface{}{1, 2, 3, 4})
		watchedList.Add(5)
		watchedList.Remove(5)

		require.Equal(t, 4, len(watchedList.GetCurrentItems()))
		require.Equal(t, 0, len(watchedList.GetNewItems()))
		require.Equal(t, 0, len(watchedList.GetRemovedItems()))
	})

	t.Run("should be able to update watched list items", func(t *testing.T) {
		watchedList := shared.NewWatchedList([]interface{}{1, 2, 3, 4})
		watchedList.Update([]interface{}{1, 4, 5})

		require.Equal(t, []interface{}{2, 3}, watchedList.GetRemovedItems())
		require.Equal(t, []interface{}{5}, watchedList.GetNewItems())
	})
}
