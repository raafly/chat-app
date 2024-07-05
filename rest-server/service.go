package restserver

import (
	"context"
	"math/rand"
	"time"

	"github.com/raafly/realtime-app/helper"
)

type AuthService interface {
	Create( dto *UserReq) error
	Login( dto *UserReq) (int, error)
	// GetProfile( id string) (*User, error)
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

func (s *AuthServiceImpl) Create(dto *UserReq) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dto.Password = s.pasword.HashPassword(dto.Password)

	err := s.repo.Create(ctx, dto)
	if err != nil {
		return helper.ErrInternalServerError()
	}

	return nil
}

func (s *AuthServiceImpl) Login(dto *UserReq) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := s.repo.FindByEmail(ctx, dto.Name)
	if err != nil {
		return 0, helper.ErrNotFound("name not found")
	}

	if s.pasword.ComparePassword(resp.Password, dto.Password) != nil {
		return 0, helper.ErrBadRequest("name or password is incocrret")
	}

	token := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	return token.Int(), nil
}

func (s *AuthServiceImpl) GetContacts(id string) (*[]Contact, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := s.repo.GetContacts(ctx, id)
	if err != nil {
		return nil, helper.ErrNotFound("contact not found")
	}

	return resp, nil
}

func (s *AuthServiceImpl) GetHistory(userIO string, contactID string) (*[]Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := s.repo.GetHistory(ctx, userIO, contactID)
	if err != nil {
		return nil, helper.ErrNotFound("message not found")
	}

	return resp, nil
}