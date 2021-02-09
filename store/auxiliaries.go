package store

import (
	"fmt"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

func (st *Store) GetCathegory(cathegory string) (int, error) {
	rows, err := st.DB.Query(`SELECT id FROM public.cathegory where cathegory=$1`, cathegory)
	if err != nil {
		return 0, fmt.Errorf("The table is empty")
	}
	var id int
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	return id, err
}

func (st *Store) GetAuthor(author string) (int, error) {
	rows, err := st.DB.Query(`SELECT id FROM public.author where author=$1`, author)
	if err != nil {
		return 0, fmt.Errorf("The table is empty")
	}
	var id int
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

func (st *Store) GetOrder(book string) (int, error) {
	if book == "0" {
		return 0, nil
	}
	rows, err := st.DB.Query(`SELECT loan_id FROM public.order where book_id=$1`, book)
	if err != nil {
		return 0, fmt.Errorf("The table is empty")
	}
	var id int
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

func (st *Store) GetAvailableBooks(books []string) (available []string, err error) {
	var x int
	//arreglar
	var status bool
	var id int
	query := fmt.Sprintf(`SELECT available, id
	FROM public.books
	WHERE tittle IN (%s)`, strings.Join(books, ", "))
	rows, err := st.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("no")
	}
	rows.Next()
	rows.Scan(&status, &id)
	if status == false {
		available = append(available, "No disponible")
		x = 1
	}
	if status == true {
		available = append(available, strconv.Itoa(id))
	}

	if x == 1 {
		return available, fmt.Errorf("Libros no disponibles")
	}
	return available, nil
}

func (st *Store) CreateOrder(books []string, id int) error {
	for _, book := range books {
		_, err := st.DB.Query(`INSERT INTO public.orders(book_id, loan_id)
			VALUES ($1, $2)`, book, id)
		if err != nil {
			return fmt.Errorf("Cannot add")
		}
	}
	return nil
}
