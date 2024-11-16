package internal

import "github.com/axadjonovsardorbek/tender/clients"

type Handler struct {
	Clients *clients.Clients
}

func NewHandler(client clients.Clients) *Handler {
	return &Handler{
		Clients: &client,
	}
}
