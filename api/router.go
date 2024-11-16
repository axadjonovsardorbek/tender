package api

import (
	"fmt"
	"time"

	"github.com/axadjonovsardorbek/tender/api/handlers"

	_ "github.com/axadjonovsardorbek/tender/api/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Pima
// @version 1.0
// @description API for Pima
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewApi(h *handlers.Handler) *gin.Engine {
	fmt.Println("qqqqqqqqqqqqqqqqqqqqq")
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())


	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Println("ddddddddddddddd")

	router.POST("tender/create", h.CreateTender)

	router.POST("/img-upload", h.UploadFile)

	return router
}
