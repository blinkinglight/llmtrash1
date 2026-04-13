package repository

import (
	"context"

	"gorm.io/gorm"
	"gotest/internal/domain/entity"
)

type SQLiteUserRepository struct {
	db *gorm.DB
}

func NewSQLiteUserRepository(db *gorm.DB) *SQLiteUserRepository {
	// Migrate the schema
	db.AutoMigrate(&entity.User{})
	return &SQLiteUserRepository{db: db}
}

func (r *SQLiteUserRepository) Save(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *SQLiteUserRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
