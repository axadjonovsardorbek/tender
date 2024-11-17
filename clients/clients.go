package clients

import (
	"github.com/axadjonovsardorbek/tender/config"
	"github.com/axadjonovsardorbek/tender/internal/auth"
	"github.com/axadjonovsardorbek/tender/internal/bid"
	"github.com/axadjonovsardorbek/tender/internal/tender"
	"github.com/axadjonovsardorbek/tender/platform"
)

type Clients struct {
	Auth   *auth.AuthService
	Tender *tender.TenderService
	Bid    *bid.BidService
}

func NewClients(cfg *config.Config, conn *platform.Storage) (*Clients, error) {

	return &Clients{
		Tender: tender.NewTenderService(&conn.TenderS),
		Auth:   auth.NewAuthService(&conn.AuthS),
		Bid:    bid.NewBidService(&conn.BidS),
	}, nil
}
