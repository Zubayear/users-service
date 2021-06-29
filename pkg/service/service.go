package service

import (
	"context"

	"github.com/Zubayear/bruce-almighty/external/entity"
	"github.com/Zubayear/bruce-almighty/external/repo"
)

type Service interface {
	CreateUser(ctx context.Context, name string) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
}

type UserService struct {
	repository repo.UserRepository
}

func (s *UserService) CreateUser(ctx context.Context, name string) (string, error) {
	e, err := s.repository.CreateUser(&entity.User{Name: name})
	if err != nil {
		return "", err
	}
	return e, nil
}

func (s *UserService) GetUser(ctx context.Context, id string) (string, error) {
	e, err := s.repository.GetUser(id)
	if err != nil {
		return "", err
	}
	return e, nil
}

func New(repo repo.UserRepository) Service {
	return &UserService{
		repository: repo,
	}
}
