package service

import "booksApi/pkg/repository"

type Books interface {
}

type Book interface {
}

type Service struct {
	Books
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
