package usecase

import (
	"fmt"

	"github.com/eiizu/go-service/entity"
)

type StoreUser interface {
	GetUsers() ([]entity.User, error)
	GetUser(string) (*entity.User, error)
	CreateUser(entity.User) (*entity.User, error)
	UpdateUser(entity.User) (*entity.User, error)
	DeleteUser(string) (*entity.User, error)
}

type Users struct {
	store StoreUser
}

func NewUsers(db StoreUser) *Users {
	var us Users
	us.store = db
	return &us
}

func (us *Users) GetUsers() ([]entity.User, error) { // regresar por id
	user, err := us.store.GetUsers()
	return user, err
}

func (us *Users) GetUser(key string) (*entity.User, error) {
	user, err := us.store.GetUser(key)
	return user, err
}

func (us *Users) CreateUser(data entity.User) (*entity.User, error) {
	user, err := us.store.CreateUser(data)
	return user, err
}

func (us *Users) UpdateUser(data entity.User) (*entity.User, error) {
	user, err := us.store.UpdateUser(data)
	return user, err
}

func (us *Users) DeleteUser(key string) (*entity.User, error) {
	_, err := us.store.DeleteUser(key)
	if err != nil {
		return nil, fmt.Errorf("The User was erased")
	}
	return nil, err
}
