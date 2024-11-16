package handlers

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/axadjonovsardorbek/tender/pkg/models"
	"github.com/gin-gonic/gin"
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
func (h *Handler) CreateTender(c *gin.Context) {
	fmt.Println("ssssssssssssssssssssssssssss")
	var req models.Tender

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"Error": err})
		slog.Error("Error binding request body: ", err)
		return
	}

	res, err := h.Clients.Tender.CreateTender(context.TODO(), req)
	if err!= nil {
        c.JSON(500, gin.H{"Error": err})
        slog.Error("Error creating tender: ", err)
        return
    }

	c.JSON(201, res)
}
