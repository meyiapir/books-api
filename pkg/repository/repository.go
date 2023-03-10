package repository

import (
	"booksApi"
	"github.com/jmoiron/sqlx"
)

type Books interface {
	GetAll() ([]booksApi.Book, error)
}

type Book interface {
	GetById(id int) (booksApi.Book, error)
	Create(input booksApi.Book) (int, error)
	Update(id int, input booksApi.UpdateBookInput) error
	Delete(id int) error
}

type Repository struct {
	Books
	Book
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Books: NewBookPostgres(db),
		Book:  NewBookPostgres(db),
	}
}
