package repository

type Books interface {
}

type Book interface {
}

type Repository struct {
	Books
	Book
}

func NewRepository() *Repository {
	return &Repository{}
}
