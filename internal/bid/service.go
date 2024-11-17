package bid

import (
	"context"

	"github.com/axadjonovsardorbek/tender/pkg/models"
)

type BidService struct {
	Bid BidI
}

func NewBidService(repo *BidI) *BidService {
	return &BidService{Bid: *repo}
}

func (s *BidService) Create(ctx context.Context, req *models.CreateBidReq) (*models.Void, error) {
	res, err := s.Bid.Create(ctx, req)

	if err != nil {
		return nil, err
	}
	return res, err
}
func (s *BidService) GetById(ctx context.Context, req string) (*models.BidRes, error) {
	res, err := s.Bid.GetById(ctx, req)

	if err != nil {
		return nil, err
	}
	return res, err
}
func (s *BidService) GetAll(ctx context.Context, req *models.GetAllBidReq) (*models.GetAllBidRes, error) {
	res, err := s.Bid.GetAll(ctx, req)

	if err != nil {
		return nil, err
	}
	return res, err
}
func (s *BidService) Update(ctx context.Context, req *models.UpdateBidReq) (*models.Void, error) {
	_, err := s.Bid.Update(ctx, req)

	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (s *BidService) Delete(ctx context.Context, req *models.DeleteBidReq) (*models.Void, error) {
	_, err := s.Bid.Delete(ctx, req)

	if err != nil {
		return nil, err
	}
	return nil, nil
}
