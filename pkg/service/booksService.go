package service

import (
	"booksApi"
	"booksApi/pkg/repository"
)

type BooksService struct {
	repo repository.Books
}

func NewBooksService(repo repository.Books) *BooksService {
	return &BooksService{repo: repo}
}

func (s *BooksService) GetAll() ([]booksApi.Book, error) {
	books, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return books, nil
}
