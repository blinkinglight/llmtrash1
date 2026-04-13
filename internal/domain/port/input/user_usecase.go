package input

import (
	"context"
	"gotest/internal/domain/entity"
)

type UserUseCase interface {
	RegisterUser(ctx context.Context, id, email, name string) error
	GetUser(ctx context.Context, id string) (*entity.User, error)
}
