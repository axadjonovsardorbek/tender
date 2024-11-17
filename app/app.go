package app

import (
	"log"
	"net/http"

	"github.com/axadjonovsardorbek/tender/api"
	"github.com/axadjonovsardorbek/tender/api/handlers"
	"github.com/axadjonovsardorbek/tender/clients"
	"github.com/axadjonovsardorbek/tender/config"
	"github.com/axadjonovsardorbek/tender/platform"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type App struct {
	Router      *gin.Engine
	Storage     *platform.Storage
	RedisClient *redis.Client
	WsHub       *platform.WebSocketHub
	MinIO       *platform.MinIO
}

func (a *App) Initialize(cfg *config.Config) {
	// Initialize database
	stg, err := platform.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	a.Storage = stg

	// Initialize Redis
	redisClient := platform.ConnectRedis(cfg)
	a.RedisClient = redisClient.Client

	// Initialize WebSocket
	wsHub := platform.NewWebSocketHub()
	a.WsHub = wsHub
	go wsHub.Run()

	//Initialize MinIO
	minioClient, err := platform.MinIOConnect(cfg)
    if err!= nil {
        log.Fatalf("Failed to connect to MinIO: %v", err)
    }
    a.MinIO = minioClient

 
	// Initialize clients
	services, err := clients.NewClients(cfg, stg)
	if err != nil {
		log.Fatalf("error while connecting clients. err: %s", err.Error())
	}

	handler := handlers.NewHandler(*services, minioClient, a.RedisClient)

	a.Router = api.NewApi(handler)
}

func (a *App) Run(serverPort string) {
	log.Printf("Server is running on port %s...", serverPort)
	if err := http.ListenAndServe(serverPort, a.Router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (a *App) Close() {
	// Close database connection
	if a.Storage != nil {
		a.Storage.Close()
	}
	// Close Redis connection
	if a.RedisClient != nil {
		a.RedisClient.Close()
	}
}
