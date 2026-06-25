package service

import (
	"context"
	"fmt"
	"strings"

	"gomind/internal/dao"
	"gomind/internal/dto"
	"gomind/internal/model"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService(userDao *dao.UserDao) *UserService {
	return &UserService{userDao: userDao}
}

func (s *UserService) Register(ctx context.Context, req dto.RegisterRequest) (*dto.UserResponse, error) {
	username := strings.TrimSpace(req.Username)
	if username == "" {
		return nil, fmt.Errorf("username is required")
	}

	existingUser, err := s.userDao.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, fmt.Errorf("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}

	user := &model.User{
		UUID:     uuid.NewString(),
		Username: username,
		Password: string(hashedPassword),
	}
	if err := s.userDao.Create(ctx, user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		UserID:   user.UUID,
		Username: user.Username,
	}, nil
}

func (s *UserService) Login(ctx context.Context, req dto.LoginRequest) (*dto.UserResponse, error) {
	username := strings.TrimSpace(req.Username)
	if username == "" {
		return nil, fmt.Errorf("username is required")
	}

	user, err := s.userDao.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}

	return &dto.UserResponse{
		UserID:   user.UUID,
		Username: user.Username,
	}, nil
}
