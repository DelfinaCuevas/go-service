package usecase

import (
	"fmt"

	"github.com/eiizu/go-service/entity"
)

//crea el user en la base de datos

type Users struct {
	users map[string]entity.User
}

// NewSomething -
func NewUsers() *Users {
	var us Users
	us.users = make(map[string]entity.User)
	us.users["alfredcue@gmail.com"] = entity.User{
		Email:   "alfredcue@gmail.com",
		Name:    "Alfredo",
		Address: "Vista al atardecer",
		Phone:   "3310057182",
	}
	us.users["edith@gmail.com"] = entity.User{
		Email:   "edith@gmail.com",
		Name:    "Edith",
		Address: "Vista al atardecer",
		Phone:   "1234567891",
	}
	us.users["ramiro@gmail.com"] = entity.User{
		Email:   "ramiro@gmail.com",
		Name:    "Ramiro",
		Address: "Av. Isla Gomera",
		Phone:   "3364681012",
	}
	us.users["miguel@gmail.com"] = entity.User{
		Email:   "miguel@gmail.com",
		Name:    "Miguel",
		Address: "Paseo de los Granaderos",
		Phone:   "3317670422",
	}
	return &us
}

func (us *Users) GetUser(email string) (*entity.User, error) { // regresar por id
	_, ok := us.users[email]
	if !ok {
		return nil, fmt.Errorf("User dosen't exist")
	}
	user := us.users[email]
	return &user, nil
}

func (us *Users) GetUsers() ([]entity.User, error) { //regresar un arreglo de usuarios
	user := []entity.User{}
	for _, aux := range us.users {
		user = append(user, aux)
	}
	return user, nil
}

func (us *Users) CreateUser(email string, name string, addres string, phone string) (*entity.User, error) {
	_, ok := us.users[email]
	if ok {
		return nil, fmt.Errorf("User already exist")
	}
	us.users[email] = entity.User{
		Email:   email,
		Name:    name,
		Address: addres,
		Phone:   phone,
	}
	user := us.users[email]
	return &user, nil
}

func (us *Users) UpdateUser(email string, name string, address string, phone string) (*entity.User, error) {
	_, ok := us.users[email]
	if ok {
		if name == "" {
			name = us.users[email].Name
		}
		if address == "" {
			address = us.users[email].Address
		}
		if phone == "" {
			phone = us.users[email].Phone
		}
		us.users[email] = entity.User{
			Email:   email,
			Name:    name,
			Address: address,
			Phone:   phone,
		}
		user := us.users[email]
		return &user, nil
	}
	return nil, fmt.Errorf("User doesn't exist")
}

func (us *Users) DeleteUser(email string) (*entity.User, error) {
	_, ok := us.users[email]
	if ok {
		delete(us.users, email)
		return nil, fmt.Errorf("User deleted")
	}
	return nil, fmt.Errorf("User dosen't exist")
}
