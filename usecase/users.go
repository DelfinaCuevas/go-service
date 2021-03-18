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

func (us *Users) GetUser(key string) (*entity.User, error) {
	if key == "" {
		return nil, fmt.Errorf("invalid key")
	}
	user, err := us.store.GetUser(key)
	return user, err
}

func (us *Users) GetUsers() ([]entity.User, error) {
	user, err := us.store.GetUsers()
	return user, err
}

func (us *Users) CreateUser(data entity.User) (*entity.User, error) {
	switch {
	case data.Email == "":
		return nil, fmt.Errorf("invalid email")
	case data.Name == "":
		return nil, fmt.Errorf("invalid name")
	case data.Address == "":
		return nil, fmt.Errorf("invalid address")
	case data.Phone == "":
		return nil, fmt.Errorf("invalid phone")
	}
	user, err := us.store.CreateUser(data)
	return user, err
}

func (us *Users) UpdateUser(data entity.User) (*entity.User, error) {
	if data.Email == "" {
		return nil, fmt.Errorf("invalid user")
	}
	user, err := us.store.UpdateUser(data)
	return user, err
}

func (us *Users) DeleteUser(key string) (string, error) {
	if key == "" {
		return "", fmt.Errorf("invalid user")
	}
	_, err := us.store.DeleteUser(key)
	if err != nil {
		return "", fmt.Errorf("Something went wrong")
	}
	return "The user was erased", err
}
