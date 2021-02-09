package store

import (
	"fmt"

	"github.com/eiizu/go-service/entity"
	_ "github.com/lib/pq"
)

func (st *Store) GetBooks() ([]entity.Book, error) {
	rows, err := st.DB.Query(`SELECT B.id, B.tittle, B.pages, B.copies, B.available, C.cathegory, A.author 
								FROM public.books B LEFT JOIN public.cathegory C ON B.cathegory_id = C.id 
								LEFT JOIN public.author A ON B.author_id = A.id`)
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
	rows, err := st.DB.Query(`SELECT B.id, B.tittle, B.pages, B.copies, B.available, C.cathegory, A.author 
								FROM public.books B LEFT JOIN public.cathegory C ON B.cathegory_id = C.id 
								LEFT JOIN public.author A ON B.author_id = A.id where B.tittle = $1`, key)
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

	_, err = st.DB.Query(`INSERT INTO public.books(pages, copies, available, cathegory_id, author_id, tittle) 
							VALUES ($1,$2,$3,$4,$5,$6)`, data.Pages, data.Copies, data.Available, id_cathegory,
		id_author, data.Tittle)

	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}
	bk, _ = st.GetBook(data.Tittle)
	return bk, err
}

func (st *Store) UpdateBook(key string, data entity.Book) (*entity.Book, error) {
	us, err := st.GetBook(key)
	if err != nil {
		return nil, err
	}
	id_author, er := st.GetAuthor(data.Author)
	if er != nil {
		return nil, fmt.Errorf("Author unknow")
	}
	id_cathegory, er := st.GetCathegory(data.Category)
	if er != nil {
		return nil, fmt.Errorf("Cathegory unknow")
	}

	_, err = st.DB.Query(`UPDATE public.books SET pages=$1, copies=$2, available=$3, cathegory_id=$4, author_id=$5, tittle=$6
							WHERE id = $7`, data.Pages, data.Copies, data.Available, id_cathegory, id_author, data.Tittle, us.Id)

	if err != nil {
		return nil, fmt.Errorf("User doesn't exist")
	}
	us, _ = st.GetBook(data.Tittle)
	return us, err
}

func (st *Store) DeleteBook(key string) (*entity.Book, error) {
	fmt.Printf(key)
	us, err := st.GetBook(key)
	if err != nil {
		return nil, err
	}
	_, err = st.DB.Query(`DELETE FROM public.books WHERE tittle=$1`, key)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong")
	}
	_, err = st.GetBook(key)
	return us, err
}

//SELECT B.id, B.tittle, B.pages, B.copies, B.available, C.description as cathegory, A.name as author FROM public.books B LEFT JOIN public.cathegory C ON B.cathegory_id = C.id LEFT JOIN public.author A ON B.author_id = A.id;
//INSERT INTO public.books(pages, copies, available, cathegory_id, author_id, tittle) VALUES ('" + data.Pages + "','" + data.Copies + "', '" + data.Available + "', '" + id_cathegory + "', '" + id_author + "', '"+ data.Tittle + "')
