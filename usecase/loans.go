package usecase

import (
	"github.com/eiizu/go-service/entity"
)

type StoreLoan interface {
	GetLoan(string, string, string) (map[int]entity.Loan, error)
	GetLoans() (map[int]entity.Loan, error)
	CreateLoan(entity.Loan) (*entity.Loan, error)
	UpdateLoan(string, entity.Loan) (*entity.Loan, error)
}

type Loans struct {
	store StoreLoan
}

func NewLoans(db StoreLoan) *Loans {
	var ln Loans
	ln.store = db
	return &ln
}

func (ln *Loans) GetLoans(book string, user string) (loan map[int]entity.Loan, err error) {
	switch {
	case book == "" && user == "":
		{
			loan, err = ln.store.GetLoans()
		}
	case book != "" && user != "":
		{
			query := `SELECT L.id, B.tittle, U.email, L.date_begin, L.date_end, L.status, L.comments
			FROM public.orders O
			LEFT JOIN public.books B ON B.id = O.book_id
			LEFT JOIN public.loans L ON O.loan_id = L.id
			LEFT JOIN public.users U ON U.id = L.user_id
			WHERE L.user_id = $2 AND O.book_id = $1`
			loan, err = ln.store.GetLoan(query, book, user)
		}
	case book != "" || user != "":
		{
			query := `SELECT L.id, B.tittle, U.email, L.date_begin, L.date_end, L.status, L.comments
			FROM public.orders O
			LEFT JOIN public.books B ON B.id = O.book_id
			LEFT JOIN public.loans L ON O.loan_id = L.id
			LEFT JOIN public.users U ON U.id = L.user_id
			WHERE L.user_id = $2 OR O.book_id = $1`
			loan, err = ln.store.GetLoan(query, book, user)
		}
	}
	return
}

func (ln *Loans) CreateLoan(data entity.Loan) (*entity.Loan, error) {
	loan, err := ln.store.CreateLoan(data)
	return loan, err
}

func (ln *Loans) UpdateLoan(id string, data entity.Loan) (*entity.Loan, error) {
	loan, err := ln.store.UpdateLoan(id, data)
	return loan, err
}
