package main

import (
	"log"
	"time"
)

func ProcessOrder(id string) {
	time.Sleep(2 * time.Second) // имитация задержки
	order := GetOrder(id)
	if order == nil {
		log.Printf("Order %s not found", id)
		return
	}

	order.Status = "in_progress"
	log.Printf("Order %s is in progress", id)

	time.Sleep(5 * time.Second) // готовка
	order.Status = "done"
	log.Printf("Order %s is done!", id)
}
