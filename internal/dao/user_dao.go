package dao

import (
	"context"
	"errors"
	"fmt"

	"gomind/internal/model"

	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (d *UserDao) Create(ctx context.Context, user *model.User) error {
	if err := d.db.WithContext(ctx).Create(user).Error; err != nil {
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}

func (d *UserDao) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := d.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("get user by username: %w", err)
	}
	return &user, nil
}
