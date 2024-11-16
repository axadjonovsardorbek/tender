package utils

import (
	"fmt"
	"log/slog"
	"net/http"
	"regexp"

	sms "github.com/axadjonovsardorbek/tender/pkg/helper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const emailRegex = `^[a-zA-Z0-9._]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func ClaimData(c *gin.Context, data string) string {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return ""
	}

	res := claims.(jwt.MapClaims)[data].(string)

	return res
}

func SmsSender(c *gin.Context, err error, code int) {
	url := c.Request.URL
	message := fmt.Sprint("Status code: ", code, "\nEndpoint: ", url, "\nError: ", err)
	err = sms.SendMessage(message)
	if err == nil {
		slog.Info("Error is successfully sent to group by bot")
	}
}
