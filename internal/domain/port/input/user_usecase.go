package input

import (
	"context"

	"gottest/internal/domain/entity"
)

type UserUseCase interface {
	RegisterUser(ctx context.Context, id, email, name, password string) error
	GetUser(ctx context.Context, id string) (*entity.User, error)
}
