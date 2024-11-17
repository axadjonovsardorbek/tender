package tender

import (
	"context"
	"database/sql"
	"errors"

	"github.com/axadjonovsardorbek/tender/pkg/models"
)

type TenderI interface {
	Create(context.Context, *models.Tender) (*models.Void, error)
	GetById(context.Context, *models.ById) (*models.Tender, error)
	GetAll(context.Context, *models.GetAllTenderReq) (*models.GetAllTenderRes, error)
	Update(context.Context, *models.UpdateTenderReq) (*models.Void, error)
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
		INSERT INTO tenders 
			(title, description, deadline, budget, file_url, client_id)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.DB.ExecContext(ctx, query, tender.Title, tender.Description, tender.Deadline, tender.Budget, tender.FileUrl, tender.ClientID)
	return nil, err
}

func (r *TenderRepository) GetById(ctx context.Context, id *models.ById) (*models.Tender, error) {
	query := `SELECT 
				id, 
				title, 
				description, 
				deadline, 
				budget, 
				status, 
				client_id,
				file_url
			FROM 
				tenders 
			WHERE id = $1 AND deleted_at = 0`
	tender := &models.Tender{}
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&tender.ID,
		&tender.Title,
		&tender.Description,
		&tender.Deadline,
		&tender.Budget,
		&tender.ClientID,
		&tender.FileUrl,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("not found")
	}
	return tender, err
}

func (r *TenderRepository) GetAll(ctx context.Context, req *models.GetAllTenderReq) (*models.GetAllTenderRes, error) {
	query := `SELECT 
				COUNT(id) OVER () AS total_count,
				id, 
				title, 
				description, 
				deadline, 
				budget, 
				status, 
				client_id,
				file_url
			FROM 
				tenders
			WHERE deleted_at = 0`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var count int32
	tenders := &models.GetAllTenderRes{
		Tenders: make([]*models.Tender, 0),
	}
	for rows.Next() {
		tender := &models.Tender{}
		if err := rows.Scan(
			&count,
			&tender.ID,
			&tender.Title,
			&tender.Description,
			&tender.Deadline,
			&tender.Budget,
			&tender.Status,
			&tender.ClientID,
			&tender.FileUrl,
		); err != nil {
			return nil, err
		}
		if err == sql.ErrNoRows {
			return nil, errors.New("not found")
		}
		tenders.Tenders = append(tenders.Tenders, tender)
		tenders.TotalCount = int64(count)
	}
	if len(tenders.Tenders) == 0 {
		return nil, errors.New("not found")
	}

	return tenders, nil
}

func (r *TenderRepository) Update(ctx context.Context, req *models.UpdateTenderReq) (*models.Void, error) {
	query := `UPDATE tenders SET status=$1, updated_at=now() WHERE id=$2`
	_, err := r.DB.ExecContext(ctx, query, req.Status, req.ID)
	return nil, err
}

func (r *TenderRepository) Delete(ctx context.Context, id *models.ById) (*models.Void, error) {
	query := `UPDATE tenders SET deleted_at=EXTRACT(EPOCH FROM NOW()) WHERE id=$1 AND deleted_at = 0`
	_, err := r.DB.ExecContext(ctx, query, id.ID)
	return nil, err
}
