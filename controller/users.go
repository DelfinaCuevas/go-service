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
	CreateUser(entity.User) (*entity.User, error)
	UpdateUser(entity.User) (*entity.User, error)
	DeleteUser(string) (string, error)
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
	resp, err := u.UseCase.GetUser(c.Param("id"))
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
	usr, err := u.UseCase.CreateUser(data)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, usr)
}

func (u *Users) UpdateUser(c echo.Context) error {
	var data entity.User

	decoder := json.NewDecoder(c.Request().Body)

	if err := decoder.Decode(&data); err != nil {
		return c.String(http.StatusBadRequest, "invalid json")
	}

	data.Email = c.Param("id")
	resp, err := u.UseCase.UpdateUser(data)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (u *Users) DeleteUser(c echo.Context) error {
	resp, err := u.UseCase.DeleteUser(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusNoContent, resp)
}
