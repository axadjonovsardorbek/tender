package tender

import (
	"context"
	"database/sql"

	"github.com/axadjonovsardorbek/tender/pkg/models"
)

type TenderI interface {
	Create(context.Context, *models.Tender) (*models.Void, error)
	GetById(context.Context, *models.ById) (*models.Tender, error)
	GetAll(context.Context, *models.GetAllTenderReq) (*models.GetAllTenderRes, error)
	Update(context.Context, *models.ById) (*models.Void, error)
	Delete(context.Context, *models.ById) (*models.Void, error)
}

// TenderRepository is the Postgres implementation of the Repository
type TenderRepository struct {
	DB *sql.DB
}

func NewTenderRepository(db *sql.DB) *TenderRepository {
	return &TenderRepository{DB: db}
}

func (r *TenderRepository) Create(ctx context.Context, tender *models.Tender) (*models.Void, error) {
	query := `
		INSERT INTO tenders (title, description, deadline, budget, status, client_id)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
	`
	var id int64
	err := r.DB.QueryRowContext(ctx, query, tender.Title, tender.Description, tender.Deadline, tender.Budget, "open", tender.ClientID).Scan(&id)
	return nil, err
}

func (r *TenderRepository) GetById(ctx context.Context, id *models.ById) (*models.Tender, error) {
	query := `SELECT id, title, description, deadline, budget, status, client_id FROM tenders WHERE id = $1`
	tender := &models.Tender{}
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&tender.ID, &tender.Title, &tender.Description, &tender.Deadline, &tender.Budget, &tender.ClientID)
	return tender, err
}

func (r *TenderRepository) GetAll(ctx context.Context, req *models.GetAllTenderReq) (*models.GetAllTenderRes, error) {
	query := `SELECT id, title, description, deadline, budget, status, client_id FROM tenders`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tenders *models.GetAllTenderRes
	for rows.Next() {
		tender := &models.Tender{}
		if err := rows.Scan(&tender.ID, &tender.Title, &tender.Description, &tender.Deadline, &tender.Budget, &tender.ClientID); err != nil {
			return nil, err
		}
		tenders.Tenders = append(tenders.Tenders, tender)
	}
	return tenders, nil
}

func (r *TenderRepository) Update(ctx context.Context, id *models.ById) (*models.Void, error) {
	query := `UPDATE tenders SET title=$1, description=$2, deadline=$3, budget=$4, status=$5 WHERE id=$6`
	_, err := r.DB.ExecContext(ctx, query, id.ID)
	return nil, err
}

func (r *TenderRepository) Delete(ctx context.Context, id *models.ById) (*models.Void, error) {
	query := `DELETE FROM tenders WHERE id=$1`
	_, err := r.DB.ExecContext(ctx, query, id.ID)
	return nil, err
}
