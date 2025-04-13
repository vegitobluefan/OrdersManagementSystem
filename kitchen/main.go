package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/order", ReceiveOrderHandler)
	http.HandleFunc("/orders", GetOrdersHandler)

	log.Println("сервис кухни запущен на :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
