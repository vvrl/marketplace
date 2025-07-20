package services

import "marketplace/internal/storage"

type Services struct {
	UserService UserService
	AdService   AdService
}

func NewServices(s *storage.Storage) *Services {
	return &Services{
		UserService: NewUserService(s.User),
		AdService:   NewAdService(s.Ad),
	}
}
