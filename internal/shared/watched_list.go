package shared

type WatchedList struct {
	currentItems []interface{}
	initial      []interface{}
	new          []interface{}
	removed      []interface{}
}

func NewWatchedList(initialItems []interface{}) *WatchedList {
	return &WatchedList{
		currentItems: initialItems,
		initial:      initialItems,
		new:          []interface{}{},
		removed:      []interface{}{},
	}
}

func (wl *WatchedList) GetCurrentItems() []interface{} {
	return wl.currentItems
}

func (wl *WatchedList) GetNewItems() []interface{} {
	return wl.new
}

func (wl *WatchedList) GetRemovedItems() []interface{} {
	return wl.removed
}

func (wl *WatchedList) compareItems(a, b interface{}) bool {
	return a == b
}

func (wl *WatchedList) IsCurrentItem(item interface{}) bool {
	for _, v := range wl.currentItems {
		if wl.compareItems(item, v) {
			return true
		}
	}

	return false
}

func (wl *WatchedList) IsNewItem(item interface{}) bool {
	for _, v := range wl.new {
		if wl.compareItems(item, v) {
			return true
		}
	}

	return false
}

func (wl *WatchedList) IsRemovedItem(item interface{}) bool {
	for _, v := range wl.removed {
		if wl.compareItems(item, v) {
			return true
		}
	}

	return false
}

func (wl *WatchedList) RemoveFromNew(item interface{}) {
	var newNewItems []interface{}

	for _, v := range wl.new {
		if !wl.compareItems(v, item) {
			newNewItems = append(newNewItems, v)
		}
	}

	wl.new = newNewItems
}

func (wl *WatchedList) RemoveFromCurrent(item interface{}) {
	var newCurrentItems []interface{}

	for _, v := range wl.currentItems {
		if !wl.compareItems(item, v) {
			newCurrentItems = append(newCurrentItems, v)
		}
	}

	wl.currentItems = newCurrentItems
}

func (wl *WatchedList) RemoveFromRemoved(item interface{}) {
	var newRemovedItems []interface{}

	for _, v := range wl.removed {
		if !wl.compareItems(item, v) {
			newRemovedItems = append(newRemovedItems, v)
		}
	}

	wl.removed = newRemovedItems
}

func (wl *WatchedList) WasAddedInitially(item interface{}) bool {
	for _, v := range wl.initial {
		if wl.compareItems(item, v) {
			return true
		}
	}

	return false
}

func (wl *WatchedList) Exists(item interface{}) bool {
	return wl.IsCurrentItem(item)
}

func (wl *WatchedList) Add(item interface{}) {
	if wl.IsRemovedItem(item) {
		wl.RemoveFromRemoved(item)
	}

	if !wl.IsNewItem(item) && !wl.WasAddedInitially(item) {
		wl.new = append(wl.new, item)
	}

	if !wl.IsCurrentItem(item) {
		wl.currentItems = append(wl.currentItems, item)
	}
}

func (wl *WatchedList) Remove(item interface{}) {
	wl.RemoveFromCurrent(item)

	if wl.IsNewItem(item) {
		wl.RemoveFromNew(item)
		return
	}

	if !wl.IsRemovedItem(item) {
		wl.removed = append(wl.removed, item)
	}
}

func (wl *WatchedList) Update(items []interface{}) {
	var newItems []interface{}

	for _, a := range items {
		found := false

		for _, b := range wl.GetCurrentItems() {
			if wl.compareItems(a, b) {
				found = true
				break
			}
		}

		if !found {
			newItems = append(newItems, a)
		}
	}

	var removedItems []interface{}

	for _, a := range wl.GetCurrentItems() {
		found := false

		for _, b := range items {
			if wl.compareItems(a, b) {
				found = true
				break
			}
		}

		if !found {
			removedItems = append(removedItems, a)
		}
	}

	wl.currentItems = items
	wl.new = newItems
	wl.removed = removedItems
}
