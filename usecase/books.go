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
	DeleteBook(string) error
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
	if key == "" {
		return nil, fmt.Errorf("Invalid Book")
	}
	book, err := bk.store.GetBook(key)
	return book, err
}

func (bk *Books) GetBooks() ([]entity.Book, error) {
	book, err := bk.store.GetBooks()
	return book, err
}

func (bk *Books) CreateBook(data entity.Book) (*entity.Book, error) {
	switch {
	case data.Tittle == "":
		return nil, fmt.Errorf("invalid tittle")
	case data.Pages <= 0:
		return nil, fmt.Errorf("invalid number of pages")
	case data.Category == "":
		return nil, fmt.Errorf("invalid cathegory")
	case data.Author == "":
		return nil, fmt.Errorf("invalid Author")
	case data.Copies == 0:
		return nil, fmt.Errorf("invalid number of copies")
	}
	book, err := bk.store.CreateBook(data)
	return book, err
}

func (bk *Books) UpdateBook(key string, data entity.Book) (*entity.Book, error) {
	if key == "" {
		return nil, fmt.Errorf("Invalid Book")
	}
	book, err := bk.store.UpdateBook(key, data)
	return book, err
}

func (bk *Books) DeleteBook(key string) (string, error) {
	if key == "" {
		return "", fmt.Errorf("Invalid Book")
	}
	err := bk.store.DeleteBook(key)
	if err != nil {
		return "the book was erased", nil
	}
	return "", fmt.Errorf("something went wrong")
}
