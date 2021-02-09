package controller

import (
	"encoding/json"
	"net/http"

	"github.com/eiizu/go-service/entity"
	"github.com/labstack/echo"
)

type LoansUseCase interface {
	GetLoans(string, string) (map[int]entity.Loan, error)
	CreateLoan(entity.Loan) (*entity.Loan, error)
	UpdateLoan(string, entity.Loan) (*entity.Loan, error)
}

type Loans struct {
	UseCaseLoan LoansUseCase
	UseCaseBook BooksUseCase
	UseCaseUser UserUseCase
}

func NewLoans(loan LoansUseCase) *Loans {
	return &Loans{
		UseCaseLoan: loan,
	}
}

func (l *Loans) GetLoans(c echo.Context) error {
	resp, err := l.UseCaseLoan.GetLoans(c.QueryParam("book"), c.QueryParam("user")) // buscar por parametro, sacar del body
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (l *Loans) CreateLoan(c echo.Context) error {
	var data entity.Loan
	decoder := json.NewDecoder(c.Request().Body)

	if err := decoder.Decode(&data); err != nil {
		return c.String(http.StatusBadRequest, "invalid json")
	}

	switch {
	case data.Loan_User == "":
		return c.String(http.StatusBadRequest, "invalid addres")
	}
	data.Date_End = ""
	data.State = "Loan"

	res, er := l.UseCaseLoan.CreateLoan(data)
	if er != nil {
		return c.String(http.StatusBadRequest, er.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func (l *Loans) UpdateLoan(c echo.Context) error {
	var data entity.Loan

	decoder := json.NewDecoder(c.Request().Body)
	if err := decoder.Decode(&data); err != nil {
		return c.String(http.StatusBadRequest, "invalid json")
	}

	switch {
	case data.Date_End == "":
		return c.String(http.StatusBadRequest, "invalid date")
	case data.Coments == "":
		return c.String(http.StatusBadRequest, "invalid coments")
	case data.State == "":
		return c.String(http.StatusBadRequest, "invalid state")
	}

	resp, err := l.UseCaseLoan.UpdateLoan(c.Param("id"), data)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
