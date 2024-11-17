package utils

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserId(c *gin.Context) (string) {
	id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		slog.Error("Unauthorized user")
		return ""
	}

	return fmt.Sprint(id)
}
