package clients

import (
	"github.com/axadjonovsardorbek/tender/config"
	"github.com/axadjonovsardorbek/tender/internal/tender"
	"github.com/axadjonovsardorbek/tender/platform"
)

type Clients struct {
	// Auth   *auth.AuthService
	Tender *tender.TenderService
}

func NewClients(cfg *config.Config, conn *platform.Storage) (*Clients, error) {

	return &Clients{
		Tender: tender.NewTenderService(&conn.TenderS),
	}, nil
}
