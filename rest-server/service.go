package restserver

import (
	"context"
	"log"
	"time"

	"github.com/raafly/realtime-app/helper"
)

type AuthService interface {
	Create(dto *UserDTO) (string, error)
	VertifyOTP(telp string, otp int) error
	GetContacts( id string) (*[]Contact, error)
	GetHistory( userIO string, contactID string) (*[]Message, error)
}

type AuthServiceImpl struct {
	repo    AuthRepo
	pasword *helper.Password
}

func NewAuthService(repo AuthRepo, password *helper.Password) AuthService {
	return &AuthServiceImpl{
		repo:    repo,
		pasword: password,
	}
}

func (s *AuthServiceImpl) Create(dto *UserDTO) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	otp, err := s.repo.Create(ctx, dto)
	if err != nil {
		return "", err
	}

	return otp, nil
}

func (s *AuthServiceImpl) VertifyOTP(telp string, otp int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println(telp)
	result, err := s.repo.FindByTelp(ctx, telp)
	if err != nil {
		return helper.ErrNotFound("telp not found", nil)
	}

	if result.OTP != otp {
		return nil
	}

	return helper.ErrBadRequest("otp not match", nil)
}

func (s *AuthServiceImpl) GetContacts(id string) (*[]Contact, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := s.repo.GetContacts(ctx, id)
	if err != nil {
		return nil, helper.ErrNotFound("contact not found", nil)
	}

	return resp, nil
}

func (s *AuthServiceImpl) GetHistory(userIO string, contactID string) (*[]Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := s.repo.GetHistory(ctx, userIO, contactID)
	if err != nil {
		return nil, helper.ErrNotFound("message not found", nil)
	}

	return resp, nil
}