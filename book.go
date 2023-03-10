package booksApi

type Book struct {
	ID     int    `json:"-"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}
