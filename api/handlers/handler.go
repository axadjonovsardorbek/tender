package handlers

import (
	"github.com/axadjonovsardorbek/tender/clients"
	"github.com/axadjonovsardorbek/tender/platform"
)

type Handler struct {
	Clients *clients.Clients
	MinIO   *platform.MinIO
}

func NewHandler(client clients.Clients, minio *platform.MinIO) *Handler {
	return &Handler{
		Clients: &client,
		MinIO:   minio,
	}
}
