package store

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/eiizu/go-service/entity"
	_ "github.com/lib/pq"
)

func (st *Store) GetBooks() ([]entity.Book, error) {
	query, _, _ := sq.
		Select(`B.id, B.tittle, B.pages, B.copies, B.available, C.cathegory, A.author`).
		From(`public.books B`).
		LeftJoin(`public.cathegory C ON B.cathegory_id = C.id`).
		LeftJoin(`public.author A ON B.author_id = A.id`).
		OrderBy(`B.id`).ToSql()

	rows, err := st.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("The table is empty")
	}

	defer rows.Close()

	var usr []entity.Book

	for rows.Next() {
		var us entity.Book
		err := rows.Scan(&us.Id, &us.Tittle, &us.Pages, &us.Copies, &us.Available, &us.Category, &us.Author)
		if err != nil {
			return nil, err
		}
		usr = append(usr, us)
	}
	return usr, err
}

func (st *Store) GetBook(key string) (*entity.Book, error) {
	query, _, _ := sq.
		Select(`B.id, B.tittle, B.pages, B.copies, B.available, C.cathegory, A.author`).
		From(`public.books B`).
		LeftJoin(`public.cathegory C ON B.cathegory_id = C.id`).
		LeftJoin(`public.author A ON B.author_id = A.id`).
		Where(`B.tittle = $1`).PlaceholderFormat(sq.Dollar).ToSql()

	rows, err := st.DB.Query(query, key)
	if err != nil {
		return nil, fmt.Errorf("The table is empty")
	}

	defer rows.Close()
	var us entity.Book

	for rows.Next() {
		err := rows.Scan(&us.Id, &us.Tittle, &us.Pages, &us.Copies, &us.Available, &us.Category, &us.Author)
		if err != nil {
			return nil, err
		}
	}
	if us.Id == 0 {
		return &us, fmt.Errorf("Book doesn't exist!")
	}
	return &us, err
}

func (st *Store) CreateBook(data entity.Book) (*entity.Book, error) {
	bk, err := st.GetBook(data.Tittle)
	if err == nil {
		return nil, fmt.Errorf("the Book already exist")
	}
	id_author, err := st.GetAuthor(data.Author)
	if err != nil {
		return nil, fmt.Errorf("Author unknow")
	}
	id_cathegory, er := st.GetCathegory(data.Category)
	if er != nil {
		return nil, fmt.Errorf("Cathegory unknow")
	}
	query, _, _ := sq.Insert(`public.books`).
		Columns(`pages`, `copies`, `available`, `cathegory_id`, `author_id`, `tittle`).
		Values(`$1`, `$2`, `$3`, `$4`, `$5`, `$6`).
		PlaceholderFormat(sq.Dollar).ToSql()

	_, err = st.DB.Query(query, data.Pages, data.Copies, data.Available, id_cathegory,
		id_author, data.Tittle)

	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}
	bk, _ = st.GetBook(data.Tittle)
	return bk, err
}

func (st *Store) UpdateBook(key string, data entity.Book) (*entity.Book, error) {
	bk, err := st.GetBook(key)
	if err != nil {
		return nil, err
	}
	if data.Author == "" {
		data.Author = bk.Author
	}
	if data.Pages == 0 {
		data.Pages = bk.Pages
	}
	if data.Copies == 0 {
		data.Copies = bk.Copies
	}
	if data.Category == "" {
		data.Category = bk.Category
	}
	id_author, er := st.GetAuthor(data.Author)
	if er != nil {
		return nil, fmt.Errorf("Author unknow")
	}
	id_cathegory, er := st.GetCathegory(data.Category)
	if er != nil {
		return nil, fmt.Errorf("Cathegory unknow")
	}

	query, _, _ := sq.
		Update(`public.books`).
		Set(`pages`, `$1`).
		Set(`copies`, `$2`).
		Set(`available`, `$3`).
		Set(`cathegory_id`, `$4`).
		Set(`author_id`, `$5`).
		Set(`tittle`, `$6`).
		Where(`id = $7`).
		PlaceholderFormat(sq.Dollar).ToSql()

	_, err = st.DB.Query(query, data.Pages, data.Copies, data.Available, id_cathegory, id_author, data.Tittle, bk.Id)

	if err != nil {
		return nil, fmt.Errorf("User doesn't exist")
	}
	bk, _ = st.GetBook(data.Tittle)
	return bk, err
}

func (st *Store) DeleteBook(key string) error {
	_, err := st.GetBook(key)
	if err != nil {
		return err
	}
	query, _, _ := sq.
		Delete(`public.books`).
		Where(`tittle=$1`).ToSql()

	_, err = st.DB.Query(query, key)
	if err != nil {
		return fmt.Errorf("Something went wrong")
	}
	_, err = st.GetBook(key)
	return err
}

func (st *Store) GetCathegory(cathegory string) (int, error) {
	query, _, _ := sq.
		Select(`id`).
		From(`public.cathegory`).
		Where(`cathegory=$1`).ToSql()
	rows, err := st.DB.Query(query, cathegory)
	if err != nil {
		return 0, fmt.Errorf("The table is empty")
	}
	defer rows.Close()

	var id int
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	return id, err
}

func (st *Store) GetAuthor(author string) (int, error) {
	query, _, _ := sq.
		Select(`id`).
		From(`public.author`).
		Where(`author=$1`).ToSql()

	rows, err := st.DB.Query(query, author)
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
