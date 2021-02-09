package store

import (
	"fmt"

	"github.com/eiizu/go-service/entity"
	_ "github.com/lib/pq"
)

func (st *Store) GetUsers() ([]entity.User, error) {
	rows, err := st.DB.Query(`SELECT * FROM users`)
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
	rows, err := st.DB.Query(`SELECT * FROM users WHERE email=$1`, key)
	if err != nil {
		return nil, fmt.Errorf("The table is empty")
	}

	defer rows.Close()
	var us entity.User

	for rows.Next() {
		err := rows.Scan(&us.Id, &us.Name, &us.Lastname, &us.Email, &us.Address, &us.Phone)
		if err != nil {
			return nil, err
		}
	}
	if us.Id == 0 {
		return &us, fmt.Errorf("User doesn't exist!")
	}
	return &us, err
}

func (st *Store) CreateUser(data entity.User) (*entity.User, error) {
	us, err := st.GetUser(data.Email)
	if err == nil {
		return nil, fmt.Errorf("the User already exist")
	}
	_, err = st.DB.Query(`INSERT INTO public.users(name, lastname, email, address, phone) VALUES ($1,$2,$3,$4,$5)`, data.Name, data.Lastname, data.Email, data.Address, data.Phone)
	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}
	us, _ = st.GetUser(data.Email)
	return us, err
}

func (st *Store) UpdateUser(data entity.User) (*entity.User, error) {
	us, err := st.GetUser(data.Email)
	if err != nil {
		return nil, err
	}
	_, err = st.DB.Query(`UPDATE public.users SET name=$1, lastname=$2, address=$3, phone=$4 WHERE email=$5`, data.Name, data.Lastname, data.Address, data.Phone, data.Email)
	if err != nil {
		return nil, fmt.Errorf("User doesn't exist")
	}
	us, _ = st.GetUser(data.Email)
	return us, err
}

func (st *Store) DeleteUser(key string) (*entity.User, error) {
	us, err := st.GetUser(key)
	if err != nil {
		return nil, err
	}
	_, err = st.DB.Query(`DELETE FROM public.users WHERE email=$1`, key)
	if err != nil {
		return nil, fmt.Errorf("Something went wrong")
	}
	us, _ = st.GetUser(key)
	return us, err
}
