package main

import (
	"log"
	"net/http"

	"OrdersManagementSystem/kitchen"
)

func main() {
	http.HandleFunc("/order", kitchen.ReceiveOrderHandler)
	http.HandleFunc("/orders", kitchen.GetOrdersHandler)

	log.Println("сервис кухни запущен на :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
