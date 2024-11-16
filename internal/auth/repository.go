package auth

import (
	"context"
	"database/sql"

	"github.com/axadjonovsardorbek/tender/pkg/models"
)

type AuthI interface {
	Register(context.Context, *models.Register) (*models.Void, error)
	Login(context.Context, *models.Login) (*models.Token, error)
}

// TenderRepository is the Postgres implementation of the Repository
type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (r *AuthRepository) Register(ctx context.Context, tender *models.Register) (*models.Void, error) {
	query := `
		INSERT INTO tenders (title, description, deadline, budget, status, client_id)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
	`
	var id int64
	err := r.DB.QueryRowContext(ctx, query, tender.Email).Scan(&id)
	return nil, err
}

func (r *AuthRepository) Login(ctx context.Context, id *models.Login) (*models.Token, error) {
	query := `SELECT id, title, description, deadline, budget, status, client_id FROM tenders WHERE id = $1`
	tender := &models.Token{}
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&tender.Token)
	return tender, err
}
