package service

import (
	"context"

	"gotest/internal/domain/entity"
	"gotest/internal/domain/port/output"
)

type UserService struct {
	repo output.UserRepository
}

func NewUserService(repo output.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, id, email, name, password string) error {
	user, err := entity.NewUser(id, email, name, password)
	if err != nil {
		return err
	}

	return s.repo.Save(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, id string) (*entity.User, error) {
	return s.repo.GetByID(ctx, id)
}
