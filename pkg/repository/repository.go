package repository

import "github.com/jmoiron/sqlx"

type Books interface {
}

type Book interface {
}

type Repository struct {
	Books
	Book
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
