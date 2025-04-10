package main

import "sync"

var (
	mu     sync.Mutex
	orders = make(map[string]*Order)
)

func SaveOrder(order *Order) {
	mu.Lock()
	defer mu.Unlock()
	orders[order.ID] = order
}

func GetOrder(id string) *Order {
	mu.Lock()
	defer mu.Unlock()
	return orders[id]
}

func GetAllOrders() []*Order {
	mu.Lock()
	defer mu.Unlock()
	list := make([]*Order, 0, len(orders))
	for _, o := range orders {
		list = append(list, o)
	}
	return list
}
