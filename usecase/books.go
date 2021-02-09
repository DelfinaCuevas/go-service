package usecase

import (
	"fmt"

	"github.com/eiizu/go-service/entity"
)

type StoreBook interface {
	GetBooks() ([]entity.Book, error)
	GetBook(string) (*entity.Book, error)
	CreateBook(entity.Book) (*entity.Book, error)
	UpdateBook(string, entity.Book) (*entity.Book, error)
	DeleteBook(string) (*entity.Book, error)
}

type Books struct {
	store StoreBook
}

func NewBooks(db StoreBook) *Books {
	var bk Books
	bk.store = db
	return &bk
}

func (bk *Books) GetBook(key string) (*entity.Book, error) {
	book, err := bk.store.GetBook(key)
	return book, err
}

func (bk *Books) GetBooks() ([]entity.Book, error) {
	book, err := bk.store.GetBooks()
	return book, err
}

func (bk *Books) CreateBook(data entity.Book) (*entity.Book, error) {
	book, err := bk.store.CreateBook(data)
	return book, err
}

func (bk *Books) UpdateBook(key string, data entity.Book) (*entity.Book, error) {
	book, err := bk.store.UpdateBook(key, data)
	return book, err
}

func (bk *Books) DeleteBook(key string) (*entity.Book, error) {
	book, err := bk.store.DeleteBook(key)
	if err != nil {
		return nil, fmt.Errorf("The book was erased")
	}
	return book, err
}
