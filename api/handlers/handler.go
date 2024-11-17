package handlers

import (
	"github.com/axadjonovsardorbek/tender/clients"
	"github.com/axadjonovsardorbek/tender/platform"
	"github.com/go-redis/redis/v8"
)

type Handler struct {
	Clients *clients.Clients
	MinIO   *platform.MinIO
	Redis   *redis.Client
}

func NewHandler(client clients.Clients, minio *platform.MinIO, redis *redis.Client) *Handler {
	return &Handler{
		Clients: &client,
		MinIO:   minio,
		Redis:   redis,
	}
}
