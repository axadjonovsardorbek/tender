package api

import (
	"time"

	"github.com/axadjonovsardorbek/tender/api/handlers"
	"github.com/axadjonovsardorbek/tender/pkg/middleware"

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

	router.GET("/api/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	
	router.POST("/register", h.Register)
	router.POST("/login", h.Login)
	router.GET("/profile", h.Profile).Use(middleware.AuthMiddleware())
	router.PUT("/profile/update", h.UpdateProfile).Use(middleware.AuthMiddleware())
	router.DELETE("/profile/delete", h.DeleteProfile).Use(middleware.AuthMiddleware())
	

	client := router.Group("/client").Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("client"))
	{
		client.POST("/tenders", h.CreateTender)
		client.GET("/tenders", h.ListTenders)
		client.PUT("/tenders/:id", h.UpdateTender)
		client.DELETE("/tenders/:id", h.DeleteTender)
	}

	bid := router.Group("/client/tenders").Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("client"))
	{
		bid.POST("/:id/bids", h.CreateTender)
		bid.GET("/:id/bids", h.ListTenders)
		bid.PUT("/:id/award/:bid_id", h.UpdateTender)
		bid.GET("/:id/bids", h.DeleteTender)
	}

	router.DELETE("/bids/delete", h.DeleteTender)
	
	router.POST("/img-upload", h.UploadFile)

	return router
}
