package main

func AddOrUpdateItem(name string, quantity int32) *Item {
	item := GetItem(name)
	if item == nil {
		item = &Item{Name: name, Quantity: quantity}
	} else {
		item.Quantity += quantity
	}
	SaveItem(item)
	return item
}
