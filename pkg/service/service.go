package service

import (
	"booksApi"
	"booksApi/pkg/repository"
)

type Books interface {
	GetAll() ([]booksApi.Book, error)
}

type Book interface {
	GetById(id int) (booksApi.Book, error)
	Create(book booksApi.Book) (int, error)
	Update(id int, input booksApi.UpdateBookInput) error
	Delete(id int) error
}

type Service struct {
	Books
	Book
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Books: NewBooksService(repos.Books),
		Book:  NewBookService(repos.Book),
	}

}
