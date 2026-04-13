package output

import (
	"context"
	"gotest/internal/domain/entity"
)

type UserRepository interface {
	Save(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id string) (*entity.User, error)
}
