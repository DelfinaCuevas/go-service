package usecase

import "github.com/eiizu/go-service/entity"

type Books struct {
	books map[string]entity.Book
}

func NewBooks() *Books {
	return &Books{}
}

func (bk *Books) GetBook(name string) (*entity.Book, error) {
	book := bk.books[name]
	return &book, nil
}

func (bk *Books) GetBooks() (*entity.Book, error) {
	return &entity.Book{
		Name:     "",
		Pages:    0,
		Category: "",
	}, nil
}
func (bk *Books) CreateBook(name string, pages int, category string) (*entity.Book, error) {
	book := bk.books[name]
	return &book, nil
}
func (bk *Books) UpdateBook(name string, pages int, category string) (*entity.Book, error) {
	book := bk.books[name]
	return &book, nil
}

func (bk *Books) DeleteBook(name string) (*entity.Book, error) {
	book := bk.books[name]
	return &book, nil
}
