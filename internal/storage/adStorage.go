package storage

import (
	"context"
	"database/sql"
	"marketplace/internal/logger"
	"marketplace/internal/models"
	"time"
)

type (
	AdStorage interface {
		CreateAdvertisement(ctx context.Context, title, text, imageURL string, price float64, userID int) (*models.Advertisement, error)
		GetAdList(ctx context.Context, params AdsParams) ([]models.Advertisement, error)
	}

	adStorage struct {
		*sql.DB
	}
)

func (s *adStorage) CreateAdvertisement(ctx context.Context, title, text, imageURL string, price float64, userID int64) (*models.Advertisement, error) {
	query := `
	INSERT INTO ads (title, text, image_url, price, user_id)
	VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	ad := &models.Advertisement{
		Title:    title,
		Text:     text,
		ImageURL: imageURL,
		Price:    price,
		AuthorID: userID,
	}

	err := s.QueryRowContext(ctxWithTimeout, query, title, text, imageURL, price, userID).Scan(&ad.ID, &ad.CreatedAt)
	if err != nil {
		logger.Logger.Errorf("create advertisement error: %v", err)
		return nil, err
	}

	return ad, nil
}

func (s *adStorage) GetAdList(ctx context.Context, params AdsParams) ([]models.Advertisement, error) {
	ads := make([]models.Advertisement, 0)

	query := `
	SELECT a.title, a.text, a.image_url, a.price, u.login AS author_login
	FROM ads a
	JOIN users u ON a.user_id = u.id
	WHERE a.price BETWEEN $1 AND $2
	ORDER_BY 
		CASE WHEN $3 = 'price' THEN a.price ELSE a.created_at END 
	$4
	LIMIT $5 OFFSET $6
	`

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := s.QueryContext(ctxWithTimeout, query,
		params.MinPrice,
		params.MaxPrice,
		params.Order,
		params.Duration,
		params.Limit,
		(params.Page-1)*params.Limit,
	)

	if err != nil {
		logger.Logger.Errorf("get ad list error: %v", err)
		return nil, err
	}

	for rows.Next() {
		var temp models.Advertisement
		err = rows.Scan(
			&temp.Title,
			&temp.Text,
			&temp.ImageURL,
			&temp.Price,
			&temp.AuthorID,
		)
		if err != nil {
			return nil, err
		}
		ads = append(ads, temp)
	}

	return ads, nil
}
