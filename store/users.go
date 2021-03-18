package store

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/eiizu/go-service/entity"
	_ "github.com/lib/pq"
)

func (st *Store) GetUsers() ([]entity.User, error) {
	query, _, _ := sq.
		Select(`*`).
		From(`public.users`).ToSql()

	rows, err := st.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("The table is empty")
	}

	defer rows.Close()

	var usr []entity.User

	for rows.Next() {
		var us entity.User
		err := rows.Scan(&us.Id, &us.Name, &us.Lastname, &us.Email, &us.Address, &us.Phone)
		if err != nil {
			return nil, err
		}
		usr = append(usr, us)
	}
	return usr, err
}

func (st *Store) GetUser(key string) (*entity.User, error) {
	query, _, _ := sq.
		Select(`*`).
		From(`public.users`).
		Where(`email = $1`).ToSql()

	rows, err := st.DB.Query(query, key)
	if err != nil {
		return nil, err
	}

	var us entity.User

	for rows.Next() {
		err := rows.Scan(&us.Id, &us.Name, &us.Lastname, &us.Email, &us.Address, &us.Phone)
		if err != nil {
			return nil, err
		}
	}

	if us.Name == "" {
		return nil, fmt.Errorf("User doesn't exist!")
	}
	return &us, nil
}

func (st *Store) CreateUser(data entity.User) (*entity.User, error) {
	us, err := st.GetUser(data.Email)
	if err == nil {
		return nil, fmt.Errorf("The User already exist")
	}

	query, _, _ := sq.
		Insert(`public.users`).
		Columns(`name`, `lastname`, `email`, `address`, `phone`).
		Values(`$1`, `$2`, `$3`, `$4`, `$5`).
		PlaceholderFormat(sq.Dollar).ToSql()

	fmt.Println(query)

	_, err = st.DB.Query(query, data.Name, data.Lastname, data.Email, data.Address, data.Phone)
	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}

	us, err = st.GetUser(data.Email)
	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}

	return us, err
}

func (st *Store) UpdateUser(data entity.User) (*entity.User, error) {
	us, err := st.GetUser(data.Email)
	if err != nil {
		return nil, err
	}
	query, _, _ := sq.
		Update("public.users").
		Set("name", "$1").
		Set("lastname", "$2").
		Set("address", "$3").
		Set("phone", "$4").
		Where("email = $5").
		PlaceholderFormat(sq.Dollar).ToSql()

	_, err = st.DB.Query(query, data.Name, data.Lastname, data.Address, data.Phone, data.Email)
	if err != nil {
		return nil, err
	}
	us, _ = st.GetUser(data.Email)
	return us, err
}

func (st *Store) DeleteUser(key string) (*entity.User, error) {
	us, err := st.GetUser(key)
	if err != nil {
		return nil, err
	}
	query, _, _ := sq.
		Delete(`public.users`).
		Where(`email = $1`).ToSql()

	_, err = st.DB.Query(query, key)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong")
	}
	_, err = st.GetUser(key)
	return us, err
}
