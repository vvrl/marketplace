package services

import (
	"context"
	"errors"
	"fmt"
	"marketplace/internal/logger"
	"marketplace/internal/models"
	"marketplace/internal/storage"
)

type (
	UserService interface {
		Register(ctx context.Context, login, password string) (*models.User, error)
		Login(ctx context.Context, login, password string) error
	}

	userService struct {
		storage storage.UserStorage
	}
)

func NewUserService(storage storage.UserStorage) UserService {
	return &userService{storage: storage}
}

func (s *userService) Register(ctx context.Context, login, password string) (*models.User, error) {

	exists, err := s.storage.IsExists(ctx, login)
	if err != nil {
		logger.Logger.Errorf("%v", err)
		return nil, err
	}
	if exists {
		textErr := fmt.Sprintf("user with login: %s alredy exists", login)
		return nil, errors.New(textErr)
	}

	if err = validateCredentials(login, password); err != nil {
		logger.Logger.Errorf("invalid login or password: %v", err)
		return nil, err
	}

	HashPassword, err := HashPassword(password)
	if err != nil {
		logger.Logger.Errorf("password hashing error: %v", err)
		return nil, err
	}

	return s.storage.CreateUser(ctx, login, HashPassword)
}

func (s *userService) Login(ctx context.Context, login, password string) error {
	return nil
}
