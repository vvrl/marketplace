package storage

import "database/sql"

type Storage struct {
	DB   *sql.DB
	User *userStorage
	Ad   *adStorage
}

func New(db *sql.DB) *Storage {
	return &Storage{
		DB:   db,
		User: &userStorage{db},
		Ad:   &adStorage{db},
	}
}
