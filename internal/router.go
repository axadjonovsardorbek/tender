package internal

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router, handler *Handler) {
	r.HandleFunc("/tenders", handler.Clients.Tender.CreateTender()).Methods("POST")
	r.HandleFunc("/tenders/{id:[0-9]+}", handler.GetTender).Methods("GET")
	r.HandleFunc("/tenders", handler.ListTenders).Methods("GET")
}
