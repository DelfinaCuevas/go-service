package router

import (
	"github.com/labstack/echo"
)

// StatusController -
type StatusController interface {
	HandlerStatusz(c echo.Context) error
	HandlerHealthz(c echo.Context) error
}

// SomethingController -
type SomethingController interface {
	HandlerSomething(c echo.Context) error
}

type UsersController interface {
	GetUser(c echo.Context) error
	GetUsers(c echo.Context) error
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type BooksController interface {
	GetBooks(c echo.Context) error
	GetBook(c echo.Context) error
	CreateBook(c echo.Context) error
	UpdateBook(c echo.Context) error
	DeleteBook(c echo.Context) error
}

type LoansController interface {
	GetLoans(c echo.Context) error
	CreateLoan(c echo.Context) error
	UpdateLoan(c echo.Context) error
}

// New -
func New(somethingC SomethingController, statusC StatusController, userC UsersController, bookC BooksController, loanC LoansController) *echo.Echo {
	e := echo.New()

	e.GET("/statusz", statusC.HandlerStatusz)
	e.GET("/healthz", statusC.HandlerHealthz)
	e.POST("/operation", somethingC.HandlerSomething)

	//--------USERS--------//
	e.GET("/users", userC.GetUsers)          //Traer todos los usuarios
	e.GET("/users/:id", userC.GetUser)       //Traer un usuario en especifico
	e.POST("/users", userC.CreateUser)       //Crear un usuario nuevo
	e.PUT("/users/:id", userC.UpdateUser)    //Modificar un usuario en especifico
	e.DELETE("/users/:id", userC.DeleteUser) //Eliminar un usuario

	//--------BOOKS--------//
	e.GET("/books", bookC.GetBooks)          //Traer todos los libros
	e.GET("/books/:id", bookC.GetBook)       //Traer un libro en especifico
	e.POST("/books", bookC.CreateBook)       //Crear un libro nuevo
	e.PUT("/books/:id", bookC.UpdateBook)    //Modificar un libro en especifico
	e.DELETE("/books/:id", bookC.DeleteBook) //Eliminar un libro

	//---------LOANS--------////
	e.GET("/loans", loanC.GetLoans)
	e.POST("/loans", loanC.CreateLoan)
	e.PUT("/loans/:id", loanC.UpdateLoan)

	return e
}
