package app

import (
	"log"
	"net/http"

	"github.com/axadjonovsardorbek/tender/api"
	"github.com/axadjonovsardorbek/tender/api/handlers"
	"github.com/axadjonovsardorbek/tender/clients"
	"github.com/axadjonovsardorbek/tender/config"
	"github.com/axadjonovsardorbek/tender/platform"

	"github.com/gorilla/mux"
)

type App struct {
	Router      *mux.Router
	Storage     *platform.Storage
	RedisClient *platform.Redis
	WsHub       *platform.WebSocketHub
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
	a.RedisClient = redisClient

	// Initialize WebSocket
	wsHub := platform.NewWebSocketHub()
	a.WsHub = wsHub
	go wsHub.Run()

	services, err := clients.NewClients(cfg, stg)
	if err != nil {
		log.Fatalf("error while connecting clients. err: %s", err.Error())
	}

	handler := handlers.NewHandler(*services)

	// Setup Router
	a.Router = mux.NewRouter()

	// Setup router
	api.RegisterRoutes(a.Router, handler)
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
