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
// @Param tender body models.CreateTenderReq true "Create Tender"
// @Success 201 {object} string "Create Tender"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /client/tenders [post]
func (h *Handler) CreateTender(c *gin.Context) {
	var body models.CreateTenderReq

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		hp.SmsSender(c, err, http.StatusBadRequest)
		slog.Error("Error binding request body: ", "err", err)
		return
	}

	if body.Budget <= 0 || body.Deadline == "" || body.Description == "" || body.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input"})
		hp.SmsSender(c, err, http.StatusBadRequest)
		slog.Error("Invalid input")
		return
	}

	user_id := hp.GetUserId(c)

	req := &models.Tender{
		Title:       body.Title,
		Description: body.Description,
		Budget:      int64(body.Budget),
		Deadline:    body.Deadline,
		FileUrl:     body.FileUrl,
		ClientID:    user_id,
	}

	_, err = h.Clients.Tender.CreateTender(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"Error": err})
		slog.Error("Error creating tender: ", "err", err)
		hp.SmsSender(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, req.Title)
}

// ListTender godec
// @Summary List Tender
// @Description List Tender
// @Tags Tender
// @Accept application/json
// @Produce application/json
// @Success 200 {object} string "Success list tenders"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /client/tenders [get]
func (h *Handler) ListTenders(c *gin.Context) {
	user_id := hp.GetUserId(c)

	req := &models.GetAllTenderReq{
		ClientID: user_id,
	}

	cacheKey := "tenders:"

	res := models.GetAllTenderRes{}

	err := hp.GetCachedData(c, h.Redis, cacheKey, &res)
	if err == nil {
		slog.Info("tenders list retrieved from cache")
		c.JSON(200, res)
		return
	}

	resp, err := h.Clients.Tender.ListTenders(context.Background(), req)
	if err != nil {
		if err.Error() == "not found" {
			c.JSON(500, gin.H{"message": "not_found"})
			hp.SmsSender(c, err, http.StatusInternalServerError)
			return
		} else {
			c.JSON(500, gin.H{"message": err})
			slog.Error("Error listing tenders: ", "err", err)
			hp.SmsSender(c, err, http.StatusInternalServerError)
			return
			
		}
		
	}

	res = *resp

	hp.CacheData(c, h.Redis, cacheKey, res)

	c.JSON(http.StatusOK, res)
}

// UpdateTender godec
// @Summary Update Tender
// @Description Update Tender
// @Tags Tender
// @Accept application/json
// @Produce application/json
// @Param id query string false "Tender ID"
// @Param tender body models.UpdateStatus true "Create Tender"
// @Success 200 {object} string "Tender status updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /client/tenders/{id} [put]
func (h *Handler) UpdateTender(c *gin.Context) {
	var body models.UpdateStatus

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		hp.SmsSender(c, err, http.StatusBadRequest)
		slog.Error("Error binding request body: ", "err", err)
		return
	}

	if body.Status != "closed" && body.Status != "open" && body.Status != "awarded" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid tender status"})
		hp.SmsSender(c, err, http.StatusBadRequest)
		slog.Error("Invalid status")
		return
	}

	req := &models.UpdateTenderReq{
		ID:     c.Query("id"),
		Status: body.Status,
	}

	_, err = h.Clients.Tender.UpdateTender(context.Background(), req)
	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"message": "Tender not found"})
			hp.SmsSender(c, err, http.StatusNotFound)
			return
		}
		c.JSON(500, gin.H{"Error": err})
		slog.Error("Error updating tender: ", "err", err)
		hp.SmsSender(c, err, http.StatusInternalServerError)
		return
	}

	// Clear relevant cache keys
	h.Redis.Del(c, "tenders:")

	c.JSON(http.StatusOK, gin.H{"message": "Tender status updated"})
}

// DeleteTender godec
// @Summary Delete Tender
// @Description Delete Tender
// @Tags Tender
// @Accept application/json
// @Produce application/json
// @Param id query string false "Tender ID"
// @Success 200 {object} string "Tender deleted successfully"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /client/tenders/{id} [delete]
func (h *Handler) DeleteTender(c *gin.Context) {
	req := &models.ById{
		ID: c.Query("id"),
	}

	_, err := h.Clients.Tender.DeleteTender(context.Background(), req)
	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"message": "Tender not found or access denied"})
			hp.SmsSender(c, err, http.StatusNotFound)
			return
		}
		c.JSON(500, gin.H{"Error": err})
		slog.Error("Error deleting tender: ", "err", err)
		hp.SmsSender(c, err, http.StatusInternalServerError)
		return
	}

	// Clear relevant cache keys
	h.Redis.Del(c, "tenders:")

	c.JSON(http.StatusOK, gin.H{"message": "Tender deleted successfully"})
}
