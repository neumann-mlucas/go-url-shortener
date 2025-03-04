package repository

import (
	"database/sql"
	"errors"
	"fmt"

	model "github.com/neumann-mlucas/go-url-shortener/internal/model"
	utils "github.com/neumann-mlucas/go-url-shortener/internal/utils"
)

var ErrNotFound = errors.New("value not found in the database")

type ShortUrlRepository interface {
	CreateShortUrl(shorturl string) error
	GetShortUrlByID(id int64) (*model.ShortUrl, error)
	GetShortUrlByHash(hash string) (*model.ShortUrl, error)
	GetShortUrls(limit int) ([]*model.ShortUrl, error)
}

type shortUrlRepository struct {
	db *sql.DB
}

func NewShortUrlRepository(db *sql.DB) ShortUrlRepository {
	return &shortUrlRepository{db: db}
}

// CreateShortUrl creates a new short URL entry in the database generating a unique hash for the URL
func (r *shortUrlRepository) CreateShortUrl(fullurl string) error {
	var result sql.Result

	insert_query := "INSERT INTO urls (hash, url) VALUES (?, ?)"
	result, err := r.db.Exec(insert_query, "dummy", fullurl)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	hash, err := utils.ToHash(uint64(id))
	if err != nil {
		return err
	}

	update_query := "UPDATE urls SET hash = ? WHERE id = ?"
	_, err = r.db.Exec(update_query, hash, id)
	if err != nil {
		return err
	}

	return err
}

// GetShortUrlByID retrieves a short URL record from the database by its unique ID.
func (r *shortUrlRepository) GetShortUrlByID(id int64) (*model.ShortUrl, error) {
	var shorturl model.ShortUrl
	query := "SELECT id, hash, url, active FROM urls WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&shorturl.Id, &shorturl.Hash, &shorturl.Url, &shorturl.Active)
	if err != nil {
		return nil, ErrNotFound
	}
	if !shorturl.Active {
		return nil, ErrNotFound
	}
	return &shorturl, nil
}

// GetShortUrlByID retrieves a short URL record from the database by its unique hash.
func (r *shortUrlRepository) GetShortUrlByHash(hash string) (*model.ShortUrl, error) {
	var shorturl model.ShortUrl
	id, err := utils.ToID(hash)
	if err != nil {
		return nil, err
	}

	query := "SELECT id, hash, url, active FROM urls WHERE id = ?"
	err = r.db.QueryRow(query, id).Scan(&shorturl.Id, &shorturl.Hash, &shorturl.Url, &shorturl.Active)
	if err != nil {
		return nil, ErrNotFound
	}
	if !shorturl.Active {
		return nil, ErrNotFound
	}
	fmt.Println(err, shorturl)
	return &shorturl, nil
}

func (r *shortUrlRepository) GetShortUrls(limit int) ([]*model.ShortUrl, error) {
	rows, err := r.db.Query("SELECT id, hash, url, active FROM urls ORDER BY id LIMIT ?", limit)
	if err != nil {
		return nil, err
	}

	urls := []*model.ShortUrl{}
	for rows.Next() {
		shorturl := &model.ShortUrl{}
		if err := rows.Scan(&shorturl.Id, &shorturl.Hash, &shorturl.Url, &shorturl.Active); err != nil {
			return nil, err
		}
		urls = append(urls, shorturl)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(urls) == 0 {
		return nil, ErrNotFound
	}
	return urls, nil
}
