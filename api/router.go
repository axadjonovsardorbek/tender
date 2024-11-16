package api

import (
	"net/http"

	"github.com/axadjonovsardorbek/tender/api/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Tender
// @version 1.0
// @description API for Tender
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func RegisterRoutes(r *mux.Router, h *handlers.Handler) {
	// Swagger UI ko'rsatish
	r.Handle("/api/swagger/", http.StripPrefix("/api/swagger/", http.FileServer(http.Dir("./swaggerui"))))
	r.Handle("/swagger/*", httpSwagger.WrapHandler)

	// Tenders endpoint
	r.HandleFunc("/tenders", h.CreateTender).Methods("POST")
	// r.HandleFunc("/tenders/{id:[0-9]+}", handler.GetTender).Methods("GET")
	// r.HandleFunc("/tenders", handler.ListTenders).Methods("GET")
}
