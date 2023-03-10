package service

import (
	"booksApi"
	"booksApi/pkg/repository"
)

type BookService struct {
	repo repository.Book
}

func NewBookService(repo repository.Book) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetById(id int) (booksApi.Book, error) {
	book, err := s.repo.GetById(id)
	if err != nil {
		return booksApi.Book{}, err
	}
	return book, nil
}

func (s *BookService) Create(book booksApi.Book) (int, error) {
	id, err := s.repo.Create(book)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *BookService) Update(id int, input booksApi.UpdateBookInput) error {
	err := s.repo.Update(id, input)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookService) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
