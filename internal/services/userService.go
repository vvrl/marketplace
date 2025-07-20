package services

import "marketplace/internal/storage"

type (
	UserService interface {
		Login()
		Register()
	}

	userService struct {
		storage storage.UserStorage
	}
)

func NewUserService(storage storage.UserStorage) UserService {
	return &userService{storage: storage}
}

func (s *userService) Login() {

}

func (s *userService) Register() {

}
