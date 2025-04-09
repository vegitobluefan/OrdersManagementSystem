package main

import "net/http"

type handler struct {
	// gateway
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HanleCreateOrder)
}

func (h *handler) HanleCreateOrder(w http.ResponseWriter, r *http.Request) {

}
