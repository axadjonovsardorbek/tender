package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/axadjonovsardorbek/tender/pkg/models"
	hp "github.com/axadjonovsardorbek/tender/pkg/utils"
	"github.com/gin-gonic/gin"
)

// CreateTender godoc
// @Summary Create Tender
// @Description Create Tender
// @Tags Tender
// @Accept application/json
// @Produce application/json
// @Param admin body models.CreateTenderReq true "Create Tender"
// @Success 201 {object} string "Create Tender"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /tenders [post]
func (h *Handler) CreateTender(c *gin.Context) {
	var body models.CreateTenderReq

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		hp.SmsSender(c, err, http.StatusBadRequest)
		slog.Error("Error binding request body: ", err)
		return
	}

	if body.Budget <= 0 || body.Deadline == "" || body.Description == "" || body.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input"})
		hp.SmsSender(c, err, http.StatusBadRequest)
		slog.Error("Invalid input")
		return
	}

	req := &models.Tender{}

	res, err := h.Clients.Tender.CreateTender(context.Background(), *req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err})
		slog.Error("Error creating tender: ", err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

// ListTender godec
// @Summary List Tender 
// @Description List Tender
// @Tags Tender
// @Accept application/json
// @Produce application/json
// @Param admin body models.CreateTenderReq true "Update Tender"
// @Success 200 {object} string "Success list tenders"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /tenders [get]
func (h *Handler) ListTenders(c *gin.Context) {
	req := &models.GetAllTenderReq{}

    res, err := h.Clients.Tender.ListTenders(context.Background(), req)
    if err!= nil {
        c.JSON(500, gin.H{"Error": err})
        slog.Error("Error listing tenders: ", err)
        return
    }

    c.JSON(http.StatusOK, res)
}

// UpdateTender godec
// @Summary Update Tender
// @Description Update Tender
// @Tags Tender
// @Accept application/json
// @Produce application/json
// @Success 200 {object} string "Success update tender"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /tenders [put]


