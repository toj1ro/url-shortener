package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"url-shortener/internals/config"
)

var (
	ErrNotFound     = errors.New("short URL not found")
	ErrInvalidKey   = errors.New("invalid short key")
	ErrDBConnection = errors.New("database connection error")
)

type SavedUrl struct {
	id  uuid.UUID
	url string
}

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(cfg config.PostgresConfig) (PostgresRepo, error) {
	db, _ := sql.Open("postgres", cfg.ConnString())

	return PostgresRepo{db: db}, nil
}

func (r *PostgresRepo) Create(longURL string) (string, error) {
	var shortUrl string

	err := r.db.QueryRow(`INSERT INTO urls (url) VALUES ($1) RETURNING id`, longURL).Scan(&shortUrl)

	if err != nil {
		return "", ErrDBConnection
	}

	return shortUrl, nil
}

func (r *PostgresRepo) Get(id string) (string, error) {
	var longURL string

	err := r.db.QueryRow(
		`SELECT url FROM urls WHERE id = $1`,
		id,
	).Scan(&longURL)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrNotFound
		}
		return "", fmt.Errorf("error getting URL: %w", err)
	}

	return longURL, nil
}
