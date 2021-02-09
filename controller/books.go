package controller

import (
	"encoding/json"
	"net/http"

	"github.com/eiizu/go-service/entity"
	"github.com/labstack/echo"
)

type BooksUseCase interface {
	GetBook(string) (*entity.Book, error)
	GetBooks() ([]entity.Book, error)
	CreateBook(entity.Book) (*entity.Book, error)
	UpdateBook(string, entity.Book) (*entity.Book, error)
	DeleteBook(string) (*entity.Book, error)
}

type Books struct {
	UseCase BooksUseCase
}

func NewBooks(book BooksUseCase) *Books {
	return &Books{
		UseCase: book,
	}
}

func (b *Books) GetBook(c echo.Context) error {
	resp, err := b.UseCase.GetBook(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (b *Books) GetBooks(c echo.Context) error {
	resp, err := b.UseCase.GetBooks()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (b *Books) CreateBook(c echo.Context) error {
	var data entity.Book

	decoder := json.NewDecoder(c.Request().Body)

	if err := decoder.Decode(&data); err != nil {
		return c.String(http.StatusBadRequest, "invalid json")
	}

	switch {
	case data.Tittle == "":
		return c.String(http.StatusBadRequest, "invalid name")
	case data.Pages <= 0:
		return c.String(http.StatusBadRequest, "invalid addres")
	case data.Category == "":
		return c.String(http.StatusBadRequest, "invalid phone")
	case data.Author == "":
		return c.String(http.StatusBadRequest, "invalid phone")
	case data.Copies == 0:
		return c.String(http.StatusBadRequest, "invalid phone")
	}

	resp, err := b.UseCase.CreateBook(data)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (b *Books) UpdateBook(c echo.Context) error {
	var data entity.Book
	decoder := json.NewDecoder(c.Request().Body)
	if err := decoder.Decode(&data); err != nil {
		return c.String(http.StatusBadRequest, "invalid json")
	}

	resp, err := b.UseCase.UpdateBook(c.Param("id"), data)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (b *Books) DeleteBook(c echo.Context) error {
	resp, err := b.UseCase.DeleteBook(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusNoContent, resp)
}
