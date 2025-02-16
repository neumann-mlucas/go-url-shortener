package repository

import (
	"database/sql"
	"fmt"
	"internal/model"
	"internal/utils"
)

type ShortUrlRepository interface {
	CreateShortUrl(shorturl *model.ShortUrl) error
	GetShortUrlById(id int64) (*model.ShortUrl, error)
}

type shortUrlRepository struct {
	db *sql.DB
}

func NewShortUrlRepository(db *sql.DB) ShortUrlRepository {
	return &shortUrlRepository{db: db}
}

func (r *shortUrlRepository) CreateShortUrl(fullurl string) error {
	var id int
	var query string

	query = "SELECT max(id) FROM urls"
	err := r.db.QueryRow(query, id).Scan(id)
	if err != nil {
		return err
	}

	hash, err := utils.ToHash(id)
	if err != nil {
		return err
	}

	query = "INSERT INTO users (id, hash, url) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, id, hash, fullurl)

	fmt.Println(result)
	return err
}

func (r *shortUrlRepository) GetShortUrlByID(id int64) (*model.ShortUrl, error) {
	var user model.Url
	query := "SELECT id, hash, url FROM urls WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&user.id, &user.hash, &user.url)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
