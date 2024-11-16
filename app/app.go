package app

import (
	"log"
	"net/http"

	"github.com/axadjonovsardorbek/tender/config"
	"github.com/axadjonovsardorbek/tender/platform"

	"github.com/gorilla/mux"
)

type App struct {
	Router      *mux.Router
	DB          *platform.Database
	RedisClient *platform.Redis
	WsHub       *platform.WebSocketHub
}

func (a *App) Initialize(cfg *config.Config) {
	// Initialize database
	db, err := platform.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	a.DB = db

	// Initialize Redis
	redisClient := platform.ConnectRedis(cfg)
	a.RedisClient = redisClient

	// Initialize WebSocket
	wsHub := platform.NewWebSocketHub()
	a.WsHub = wsHub
	go wsHub.Run()

	// Setup Router
	a.Router = mux.NewRouter()
	// TODO: Add routes here
}

func (a *App) Run(serverPort string) {
	log.Printf("Server is running on port %s...", serverPort)
	if err := http.ListenAndServe(serverPort, a.Router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (a *App) Close() {
	// Close database connection
	if a.DB != nil {
		a.DB.Close()
	}
	// Close Redis connection
	if a.RedisClient != nil {
		a.RedisClient.Close()
	}
}
