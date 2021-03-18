package usecase

import (
	"fmt"

	"github.com/eiizu/go-service/entity"
)

type StoreLoan interface {
	GetLoan(map[string]string) (map[string]entity.Loan, error)
	GetLoan_(map[string]string) (map[string]entity.Loan, error)
	GetLoans() (map[string]entity.Loan, error)
	CreateLoan(entity.Loan) (*entity.Loan, error)
	UpdateLoan(entity.Loan) (*entity.Loan, error)
}

type Loans struct {
	store StoreLoan
}

func NewLoans(db StoreLoan) *Loans {
	var ln Loans
	ln.store = db
	return &ln
}

func (ln *Loans) GetLoans(parameters map[string]string) (map[string]entity.Loan, error) {

	if parameters["book"] == "" && parameters["user"] == "" && parameters["uuid"] == "" {
		loan, err := ln.store.GetLoans()
		return loan, err
	}
	if parameters["book"] != "" && parameters["user"] != "" {
		loan, err := ln.store.GetLoan_(parameters)
		return loan, err
	}
	loan, err := ln.store.GetLoan(parameters)

	return loan, err
}

func (ln *Loans) CreateLoan(data entity.Loan) (*entity.Loan, error) {
	switch {
	case data.Loan_User == "":
		return nil, fmt.Errorf("Invalid User")
	case len(data.Loan_Book) == 0:
		return nil, fmt.Errorf("No books for the loan")
	}

	loan, err := ln.store.CreateLoan(data)

	return loan, err
}

func (ln *Loans) UpdateLoan(data entity.Loan) (*entity.Loan, error) {

	switch {
	case data.Uuid == "":
		return nil, fmt.Errorf("invalid id")
	case data.Coments == "":
		return nil, fmt.Errorf("invalid coments")
	case data.State == "":
		return nil, fmt.Errorf("invalid state")
	}

	loan, err := ln.store.UpdateLoan(data)

	return loan, err

}
