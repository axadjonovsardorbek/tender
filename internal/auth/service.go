package auth

import (
	"context"

	"github.com/axadjonovsardorbek/tender/pkg/models"
)

type AuthService struct {
	Auth AuthRepo
}

func (s *AuthService) Register(ctx context.Context, req *models.RegisterReq) (*models.TokenRes, error) {
	res, err := s.Auth.Register(ctx, req)

	if err != nil {
		return nil, err
	}
	return res, err
}
func (s *AuthService) Login(ctx context.Context, req *models.LoginReq) (*models.TokenRes, error) {
	res, err := s.Auth.Login(ctx, req)

	if err != nil {
		return nil, err
	}
	return res, err
}
func (s *AuthService) GetProfile(ctx context.Context, id string) (*models.UserRes, error) {
	res, err := s.Auth.GetProfile(ctx, id)

	if err != nil {
		return nil, err
	}
	return res, err
}
func (s *AuthService) UpdateProfile(ctx context.Context, req *models.UpdateReq) (*models.Void, error) {
	_, err := s.Auth.UpdateProfile(ctx, req)

	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (s *AuthService) DeleteProfile(ctx context.Context, id string) (*models.Void, error) {
	_, err := s.Auth.DeleteProfile(ctx, id)

	if err != nil {
		return nil, err
	}
	return nil, nil
}
