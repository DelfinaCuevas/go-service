package store

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/eiizu/go-service/entity"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func (st *Store) GetLoans() (map[string]entity.Loan, error) {
	query, _, _ := sq.
		Select(`L.uuid, B.tittle, U.email, L.date_begin, L.date_end, L.status, L.comments`).
		From(`public.orders O`).
		LeftJoin(`public.loans L ON O.uuid_loan = L.uuid`).
		LeftJoin(`public.books B ON B.id = O.book_id`).
		LeftJoin(`public.users U ON U.id = L.user_id`).ToSql()
	rows, err := st.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("The table is empty")
	}
	defer rows.Close()

	ln := make(map[string]entity.Loan)
	var book string
	for rows.Next() {
		var aux entity.Loan
		err := rows.Scan(&aux.Uuid, &book, &aux.Loan_User, &aux.Date_Begin, &aux.Date_End, &aux.State, &aux.Coments)
		if err != nil {
			return nil, err
		}
		aux.Loan_Book = ln[aux.Uuid].Loan_Book
		aux.Loan_Book = append(aux.Loan_Book, book)
		ln[aux.Uuid] = aux
	}
	return ln, err
}

func (st *Store) GetLoan(parameters map[string]string) (map[string]entity.Loan, error) {

	uid, err := st.GetUuid(parameters)
	if err != nil {
		return nil, err
	}
	queryvar := fmt.Sprintf("L.uuid IN ('%s')", strings.Join(uid, "','"))

	query, _, _ := sq.
		Select(`L.uuid, B.tittle, U.email, L.date_begin, L.date_end, L.status, L.comments`).
		From(`public.orders O`).
		LeftJoin(`public.loans L ON O.uuid_loan = L.uuid`).
		LeftJoin(`public.books B ON B.id = O.book_id`).
		LeftJoin(`public.users U ON U.id = L.user_id`).
		Where(queryvar).
		PlaceholderFormat(sq.Dollar).ToSql()

	rows, err := st.DB.Query(query)
	if err != nil {
		return nil, err
	}

	ln := make(map[string]entity.Loan)
	var book string
	defer rows.Close()

	for rows.Next() {
		var aux entity.Loan
		err := rows.Scan(&aux.Uuid, &book, &aux.Loan_User, &aux.Date_Begin, &aux.Date_End, &aux.State, &aux.Coments)
		if err != nil {
			return nil, err
		}
		aux.Loan_Book = ln[aux.Uuid].Loan_Book
		aux.Loan_Book = append(aux.Loan_Book, book)
		ln[aux.Uuid] = aux
	}
	return ln, err
}

func (st *Store) GetLoan_(parameters map[string]string) (map[string]entity.Loan, error) {

	query, _, _ := sq.
		Select(`L.uuid, B.tittle, U.email, L.date_begin, L.date_end, L.status, L.comments`).
		From(`public.orders O`).
		LeftJoin(`public.loans L ON O.uuid_loan = L.uuid`).
		LeftJoin(`public.books B ON B.id = O.book_id`).
		LeftJoin(`public.users U ON U.id = L.user_id`).
		Where(`L.user_id = (SELECT id FROM public.users WHERE email = $1) AND O.book_id = (SELECT id FROM public.books WHERE tittle = $2)`).
		PlaceholderFormat(sq.Dollar).ToSql()

	rows, err := st.DB.Query(query, parameters["user"], parameters["book"])
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	ln := make(map[string]entity.Loan)
	var book string
	for rows.Next() {
		var aux entity.Loan
		err := rows.Scan(&aux.Uuid, &book, &aux.Loan_User, &aux.Date_Begin, &aux.Date_End, &aux.State, &aux.Coments)
		if err != nil {
			return nil, err
		}
		aux.Loan_Book = ln[aux.Uuid].Loan_Book
		aux.Loan_Book = append(aux.Loan_Book, book)
		ln[aux.Uuid] = aux
	}

	return ln, err
}

func (st *Store) CreateLoan(data entity.Loan) (*entity.Loan, error) {
	var date time.Time
	date = time.Now()
	id := uuid.New().String()
	data.Uuid = id
	data.State = "On loan"
	data.Date_End = "00/00/0000"

	user, err := st.GetUser(data.Loan_User)
	if err != nil {
		return nil, err
	}
	books, err := st.GetAvailableBooks(data.Loan_Book)
	if err != nil {
		return nil, err
	}

	query, _, _ := sq.
		Insert(`public.loans`).
		Columns(`user_id`, ` date_begin`, ` status`, ` comments`, ` uuid`, ` date_end`).
		Values(`$1`, `$2`, `$3`, `$4`, `$5`, `$6`).
		PlaceholderFormat(sq.Dollar).ToSql()

	rows, err := st.DB.Query(query, user.Id, date, data.State, data.Coments, data.Uuid, data.Date_End)

	defer rows.Close()
	err = st.CreateOrder(books, data.Uuid)
	if err != nil {
		return nil, err
	}

	query, _, _ = sq.
		Select(`L.uuid, B.tittle, U.email, L.date_begin, L.date_end, L.status, L.comments`).
		From(`public.orders O`).
		LeftJoin(`public.loans L ON O.uuid_loan = L.uuid`).
		LeftJoin(`public.books B ON B.id = O.book_id`).
		LeftJoin(`public.users U ON U.id = L.user_id`).
		Where("L.uuid = $1").
		PlaceholderFormat(sq.Dollar).ToSql()

	rows, err = st.DB.Query(query, data.Uuid)
	if err != nil {
		return nil, err
	}

	var aux entity.Loan
	var book string
	for rows.Next() {
		err := rows.Scan(&aux.Uuid, &book, &aux.Loan_User, &aux.Date_Begin, &aux.Date_End, &aux.State, &aux.Coments)
		if err != nil {
			return nil, err
		}
		aux.Loan_Book = append(aux.Loan_Book, book)
	}
	return &aux, err
}
func (st *Store) UpdateLoan(data entity.Loan) (*entity.Loan, error) {
	var date time.Time
	date = time.Now()
	query, _, _ := sq.
		Update(`public.Loan`).
		Set(`status`, `$1`).
		Set(`comments`, `$2`).
		Set(`date_end`, `$3`).
		Where(`uuid = $4`).
		PlaceholderFormat(sq.Dollar).ToSql()
	rows, err := st.DB.Query(query, data.State, data.Coments, date, data.Uuid)
	if err != nil {
		return nil, err
	}

	var aux entity.Loan
	var book string
	for rows.Next() {
		err := rows.Scan(&aux.Uuid, &book, &aux.Loan_User, &aux.Date_Begin, &aux.Date_End, &aux.State, &aux.Coments)
		if err != nil {
			return nil, err
		}
		aux.Loan_Book = append(aux.Loan_Book, book)
	}
	return &aux, err
}

func (st *Store) GetUuid(parameters map[string]string) ([]string, error) {
	var uid []string
	var queryvar string

	if parameters["uuid"] != "" {
		uid = append(uid, parameters["uuid"])
		return uid, nil
	}

	query, _, _ := sq.
		Select("uuid_loan").
		From("public.orders").
		Where("book_id = (SELECT id FROM public.books WHERE tittle = $1)").
		PlaceholderFormat(sq.Dollar).ToSql()
	queryvar = parameters["book"]

	if parameters["user"] != "" {
		query, _, _ = sq.
			Select(`uuid`).
			From(`public.loans`).
			Where(`user_id = (SELECT id FROM public.users WHERE email = $1)`).
			PlaceholderFormat(sq.Dollar).ToSql()
		queryvar = parameters["user"]
	}

	rows, err := st.DB.Query(query, queryvar)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		uid = append(uid, id)
	}
	return uid, err
}

func (st *Store) GetAvailableBooks(books []string) (available []string, err error) {
	var status bool
	var id int
	query := fmt.Sprintf(`SELECT available, id
	FROM public.books
	WHERE tittle IN ('%s')`, strings.Join(books, "','"))
	rows, err := st.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&status, &id)
		if status == false {
			return nil, err
		}
		available = append(available, strconv.Itoa(id))
	}
	return available, nil
}

func (st *Store) GetOrder(book string) ([]string, error) {
	if book == "0" {
		return nil, fmt.Errorf("No books")
	}
	rows, err := st.DB.Query(`SELECT uuid_loan FROM public.orders where book_id = $1`, book)
	if err != nil {
		return nil, fmt.Errorf("The table is empty")
	}
	var id []string
	defer rows.Close()

	for rows.Next() {
		var uuid string
		err := rows.Scan(&uuid)
		if err != nil {
			return nil, err
		}
		id = append(id, uuid)
	}
	return id, nil
}

func (st *Store) CreateOrder(books []string, id string) error {

	for _, book := range books {
		_, err := st.DB.Query(`INSERT INTO public.orders(book_id, uuid_loan)
			VALUES ($1, $2)`, book, id)
		if err != nil {
			return fmt.Errorf("Cannot add")
		}
	}
	return nil
}
