package auth

import "github.com/axadjonovsardorbek/tender/pkg/models"

type AuthService struct {
	Auth AuthRepo
}

func (s *AuthService) Register(req models.RegisterReq)(*models.TokenRes, error){
	res, err := s.Auth.Register(&req)

	if err != nil {
		return nil, err
	}
	return res, err
}
func (s *AuthService) Login(req models.LoginReq)(*models.TokenRes, error){
	res, err := s.Auth.Login(&req)

	if err != nil {
		return nil, err
	}
	return res, err
}
func (s *AuthService) GetProfile(id string)(*models.UserRes, error){
	res, err := s.Auth.GetProfile(id)

	if err != nil {
		return nil, err
	}
	return res, err
}
func (s *AuthService) DeleteProfile(id string)(error){
	err := s.Auth.DeleteProfile(id)

	if err != nil {
		return err
	}
	return nil
}