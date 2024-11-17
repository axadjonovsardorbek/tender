package tender

import (
	"context"

	"github.com/axadjonovsardorbek/tender/pkg/models"
)

type TenderService struct {
	storage TenderI
}

func NewTenderService(stg *TenderI) *TenderService {
	return &TenderService{storage: *stg}
}

func (s *TenderService) CreateTender(ctx context.Context, tender *models.Tender) (*models.Void, error) {
	_, err := s.storage.Create(ctx, tender)
	if err != nil {
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

func (s *TenderService) UpdateTender(ctx context.Context, req *models.UpdateTenderReq) (*models.Void, error) {
	_, err := s.storage.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return &models.Void{}, nil
}

func (s *TenderService) DeleteTender(ctx context.Context, id *models.ById) (*models.Void, error) {
	_, err := s.storage.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	return &models.Void{}, nil
}
