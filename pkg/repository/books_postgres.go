package repository

import (
	"booksApi"
	"github.com/jmoiron/sqlx"
)

type BooksPostgres struct {
	db *sqlx.DB
}

func NewBookPostgres(db *sqlx.DB) *BooksPostgres {
	return &BooksPostgres{db: db}
}

func (r *BooksPostgres) GetAll() ([]booksApi.Book, error) {
	var books []booksApi.Book
	query := "SELECT * FROM books"
	err := r.db.Select(&books, query)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BooksPostgres) GetById(id int) (booksApi.Book, error) {
	var book booksApi.Book
	query := "SELECT * FROM books WHERE id = $1"
	err := r.db.Get(&book, query, id)
	if err != nil {
		return booksApi.Book{}, err
	}
	return book, nil
}

func (r *BooksPostgres) Create(input booksApi.Book) (int, error) {
	query := "INSERT INTO books (title, author, year) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err := r.db.Get(&id, query, input.Title, input.Author, input.Year)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *BooksPostgres) Update(id int, input booksApi.UpdateBookInput) error {
	query := "UPDATE books SET title = $1, author = $2, year = $3 WHERE id = $4"
	_, err := r.db.Exec(query, input.Title, input.Author, input.Year, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *BooksPostgres) Delete(id int) error {
	query := "DELETE FROM books WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
