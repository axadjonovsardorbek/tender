package tender

import (
	"context"

	"github.com/axadjonovsardorbek/tender/pkg/models"
)

type Service interface {
	CreateTender(ctx context.Context, req *models.Tender) (*models.Void, error)
	GetTender(ctx context.Context, req *models.ById) (*models.Tender, error)
	ListTenders(ctx context.Context, req *models.GetAllTenderReq) (*models.GetAllTenderRes, error)
	UpdateTender(ctx context.Context, req *models.ById) (*models.Void, error)
	DeleteTender(ctx context.Context, req *models.ById) (*models.Void, error)
}

type TenderService struct {
	storage TenderI
}

func NewTenderService(stg *TenderI) *TenderService {
	return &TenderService{storage: *stg}
}

func (s *TenderService) CreateTender(ctx context.Context, tender models.Tender) (*models.Void, error) {
	_, err := s.storage.Create(ctx, &tender)
    if err!= nil {
        return nil, err
    }
    return &models.Void{}, nil
}

func (s *TenderService) GetTender(ctx context.Context, id *models.ById) (*models.Tender, error) {
    return s.storage.GetById(ctx, id)
}

func (s *TenderService) ListTenders(ctx context.Context, req *models.GetAllTenderReq) (*models.GetAllTenderRes, error) {
    return s.storage.GetAll(ctx, req)
}

func (s *TenderService) UpdateTender(ctx context.Context, id *models.ById) (*models.Void, error) {
	_, err := s.storage.Update(ctx, id)
    if err!= nil {
        return nil, err
    }
    return &models.Void{}, nil
}

func (s *TenderService) DeleteTender(ctx context.Context, id *models.ById) (*models.Void, error) {
	_, err := s.storage.Delete(ctx, id)
    if err!= nil {
        return nil, err
    }
    return &models.Void{}, nil
}
