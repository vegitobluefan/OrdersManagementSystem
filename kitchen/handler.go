package main

import (
	"encoding/json"
	"net/http"
)

func ReceiveOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	order.Status = "received"
	SaveOrder(&order)

	go ProcessOrder(order.ID)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(order)
}

func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(GetAllOrders())
}
