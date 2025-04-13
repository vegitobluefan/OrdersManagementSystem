package main

import "sync"

var (
	store   = map[string]*Item{}
	storeMu sync.Mutex
)

func SaveItem(item *Item) {
	storeMu.Lock()
	defer storeMu.Unlock()
	store[item.Name] = item
}

func GetItem(name string) *Item {
	storeMu.Lock()
	defer storeMu.Unlock()
	return store[name]
}

func ListItems() []*Item {
	storeMu.Lock()
	defer storeMu.Unlock()
	var items []*Item
	for _, item := range store {
		items = append(items, item)
	}
	return items
}

func DecreaseItem(name string, amount int32) bool {
	storeMu.Lock()
	defer storeMu.Unlock()

	item, ok := store[name]
	if !ok || item.Quantity < amount {
		return false
	}
	item.Quantity -= amount
	return true
}
