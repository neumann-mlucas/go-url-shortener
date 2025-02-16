package repository

import (
	"database/sql"
	"fmt"

	model "github.com/neumann-mlucas/go-url-shortener/internal/model"
	utils "github.com/neumann-mlucas/go-url-shortener/internal/utils"
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

	id += 1

	hash := utils.ToHash(id)
	if err != nil {
		return err
	}

	query = "INSERT INTO users (id, hash, url) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, id, hash, fullurl)

	fmt.Println(result)
	return err
}

func (r *shortUrlRepository) GetShortUrlByID(id int64) (*model.ShortUrl, error) {
	var shorturl model.ShortUrl
	query := "SELECT id, hash, url FROM urls WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&shorturl.Id, &shorturl.Hash, &shorturl.Url)
	if err != nil {
		return nil, err
	}
	return &shorturl, nil
}
