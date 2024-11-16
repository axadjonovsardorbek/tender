package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/axadjonovsardorbek/tender/pkg/models"
)

// CreateTender godoc
// @Summary Create Tender
// @Description Create Tender
// @Tags Tender
// @Accept json
// @Produce json
// @Param admin body models.Tender true "Create Tender"
// @Success 201 {object} string "Create Tender"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /tender/create [post]
func (h *Handler) CreateTender(w http.ResponseWriter, r *http.Request) {
	var req models.Tender
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.Clients.Tender.CreateTender(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
