package controller

import (
	"encoding/json"
	"net/http"

	"github.com/eiizu/go-service/entity"
	"github.com/labstack/echo"
)

//go:generate mockgen -destination=./mocks/mock_usecase_loans.go -package=mocks github.com/eiizu/go-service/controller LoansUseCase
type LoansUseCase interface {
	GetLoans(map[string]string) (map[string]entity.Loan, error)
	CreateLoan(entity.Loan) (*entity.Loan, error)
	UpdateLoan(entity.Loan) (*entity.Loan, error)
}

type Loans struct {
	UseCaseLoan LoansUseCase
}

func NewLoans(loan LoansUseCase) *Loans {
	return &Loans{
		UseCaseLoan: loan,
	}
}

func (l *Loans) GetLoans(c echo.Context) error {
	parameters := make(map[string]string)
	parameters["book"] = c.QueryParam("book")
	parameters["user"] = c.QueryParam("user")
	parameters["uuid"] = c.QueryParam("uuid")
	resp, err := l.UseCaseLoan.GetLoans(parameters)
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

	data.Uuid = c.Param("id")
	res, err := l.UseCaseLoan.UpdateLoan(data)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
