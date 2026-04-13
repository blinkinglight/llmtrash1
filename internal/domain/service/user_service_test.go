package service

import (
	"context"
	"testing"

	"gottest/internal/domain/entity"
)

type MockUserRepository struct {
	SaveFunc    func(ctx context.Context, user *entity.User) error
	GetByIDFunc func(ctx context.Context, id string) (*entity.User, error)
}

func (m *MockUserRepository) Save(ctx context.Context, user *entity.User) error {
	return m.SaveFunc(ctx, user)
}

func (m *MockUserRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	return m.GetByIDFunc(ctx, id)
}

func TestUserService_RegisterUser(t *testing.T) {
	repo := &MockUserRepository{
		SaveFunc: func(ctx context.Context, user *entity.User) error {
			return nil
		},
	}
	service := NewUserService(repo)

	err := service.RegisterUser(context.Background(), "1", "test@example.com", "Test User", "password123")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
