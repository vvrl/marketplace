package services

import "marketplace/internal/storage"

type (
	AdService interface {
		PostAd()
		GetAdList()
	}

	adService struct {
		storage storage.AdStorage
	}
)

func NewAdService(storage storage.AdStorage) AdService {
	return &adService{storage: storage}
}

func (s *adService) PostAd() {

}

func (s *adService) GetAdList() {

}
