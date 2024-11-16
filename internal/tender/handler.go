package tender

import (
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
func (h *handler) CreateTender(c *gin.Context) {
	var req models.Tender
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Clients.Tender.CreateTender(c.Request.Context(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, res)
}
