package repository

import (
	"context"
	"errors"
	"sync"

	"gotest/internal/domain/entity"
)

type InMemoryUserRepository struct {
	users map[string]*entity.User
	mu    sync.RWMutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*entity.User),
	}
}

func (r *InMemoryUserRepository) Save(ctx context.Context, user *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, ok := r.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}
