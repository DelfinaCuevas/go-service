package store

import (
	"fmt"
	"strconv"
	"time"

	"github.com/eiizu/go-service/entity"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func (st *Store) GetLoans() (map[int]entity.Loan, error) {
	query := `SELECT L.id, B.tittle, U.email, L.date_begin, L.date_end, L.status, L.comments
				FROM public.orders O
				LEFT JOIN public.books B ON B.id = O.book_id
				LEFT JOIN public.loans L ON O.loan_id = L.id
				LEFT JOIN public.users U ON U.id = L.user_id`
	ln := make(map[int]entity.Loan)
	rows, err := st.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("The table is empty")
	}

	defer rows.Close()
	var book string
	for rows.Next() {
		var aux entity.Loan
		err := rows.Scan(&aux.Id, &book, &aux.Loan_User, &aux.Date_Begin, &aux.Date_End, &aux.State, &aux.Coments)
		aux.Loan_Book = ln[aux.Id].Loan_Book
		aux.Loan_Book = append(aux.Loan_Book, book)
		ln[aux.Id] = aux
		if err != nil {
			return nil, err
		}
	}
	return ln, err
}

func (st *Store) GetLoan(query string, book string, user string) (map[int]entity.Loan, error) {
	id_book, _ := st.GetBook(book)
	id_loan, _ := st.GetOrder(strconv.Itoa(id_book.Id))
	user_, _ := st.GetUser(user)
	ln := make(map[int]entity.Loan)
	rows, err := st.DB.Query(query, id_loan, user_.Id)
	if err != nil {
		return nil, fmt.Errorf("no")
	}

	defer rows.Close()

	for rows.Next() {
		var aux entity.Loan
		err := rows.Scan(&aux.Id, &book, &aux.Loan_User, &aux.Date_Begin, &aux.Date_End, &aux.State, &aux.Coments)
		if err != nil {
			return nil, fmt.Errorf("no")
		}
		aux.Loan_Book = ln[aux.Id].Loan_Book
		aux.Loan_Book = append(aux.Loan_Book, book)
		ln[aux.Id] = aux
	}
	return ln, err
}

func (st *Store) CreateLoan(data entity.Loan) (*entity.Loan, error) {
	id := uuid.New() //arreglar
	var date time.Time
	date = time.Now()
	user, err := st.GetUser(data.Loan_User)
	if err != nil {
		return nil, fmt.Errorf("Invalid User")
	}

	books, err := st.GetAvailableBooks(data.Loan_Book)
	if err != nil {
		return nil, fmt.Errorf("Prestamo no efectuado libros no disponibles")
	}

	rows, err := st.DB.Query(`INSERT INTO public.loans(
		user_id, date_begin, status, comments)
		VALUES ($1, $2, $3, $4)`, user.Id, date, data.State, data.Coments)

	defer rows.Close()

	rows.Next()
	rows.Scan(&data.Id)

	err = st.CreateOrder(books, data.Id)

	if err != nil {
		return nil, err
	}

	query := `SELECT L.id, B.tittle, U.email, L.date_begin, L.date_end, L.status, L.comments
	FROM public.orders O
	LEFT JOIN public.books B ON B.id = O.book_id
	LEFT JOIN public.loans L ON O.loan_id = L.id
	LEFT JOIN public.users U ON U.id = L.user_id
	WHERE L.id = $1`
	rows, err = st.DB.Query(query, data.Id)
	var book string
	for rows.Next() {
		err := rows.Scan(&data.Id, &book, &data.Loan_User, &data.Date_Begin, &data.Date_End, &data.State, &data.Coments)
		data.Loan_Book = append(data.Loan_Book, book)
		if err != nil {
			return nil, err
		}
	}
	return &data, err
}

func (st *Store) UpdateLoan(id string, data entity.Loan) (*entity.Loan, error) {
	return nil, nil
}
