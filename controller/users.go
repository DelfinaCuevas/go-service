package controller

import (
	"encoding/json"
	"net/http"

	"github.com/eiizu/go-service/entity"
	"github.com/labstack/echo"
)

type UserUseCase interface {
	GetUser(string) (*entity.User, error)
	GetUsers() ([]entity.User, error)
	CreateUser(string, string, string, string) (*entity.User, error)
	UpdateUser(string, string, string, string) (*entity.User, error)
	DeleteUser(string) (*entity.User, error)
}

type Users struct {
	UseCase UserUseCase
}

func NewUsers(user UserUseCase) *Users {
	return &Users{
		UseCase: user,
	}
}

func (u *Users) GetUser(c echo.Context) error {
	id := c.Param("id")
	resp, err := u.UseCase.GetUser(id) //aqui va el id
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (u *Users) GetUsers(c echo.Context) error {
	resp, err := u.UseCase.GetUsers()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (u *Users) CreateUser(c echo.Context) error {
	var data entity.User

	decoder := json.NewDecoder(c.Request().Body)

	if err := decoder.Decode(&data); err != nil {
		return c.String(http.StatusBadRequest, "invalid json")
	}

	switch {
	case data.Email == "":
		return c.String(http.StatusBadRequest, "invalid email")
	case data.Name == "":
		return c.String(http.StatusBadRequest, "invalid name")
	case data.Address == "":
		return c.String(http.StatusBadRequest, "invalid addres")
	case data.Phone == "":
		return c.String(http.StatusBadRequest, "invalid phone")
	}

	resp, err := u.UseCase.CreateUser(data.Email, data.Name, data.Address, data.Phone)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (u *Users) UpdateUser(c echo.Context) error { //Funcion update(Put) va asi?
	id := c.Param("id")

	var data entity.User

	decoder := json.NewDecoder(c.Request().Body)

	if err := decoder.Decode(&data); err != nil {
		return c.String(http.StatusBadRequest, "invalid json")
	}

	resp, err := u.UseCase.UpdateUser(id, data.Name, data.Address, data.Phone) //aqui va algo?
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (u *Users) DeleteUser(c echo.Context) error { //Funcion delete(DELETE) va asi?
	id := c.Param("id")
	resp, err := u.UseCase.DeleteUser(id) //aqui va algo?
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusNoContent, resp)
}
