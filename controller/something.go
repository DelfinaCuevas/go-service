package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

// SomethingUseCase -
type SomethingUseCase interface {
	DoSomething(string) (map[string]int, error)
	DoSomethingWithRepo(string) (map[string]string, error)
}

// Something -
type Something struct {
	UseCase SomethingUseCase
	Render  *render.Render
}

// Request -
type Request struct {
	Info string `json:"info"`
}

// NewSomething -
func NewSomething(uc SomethingUseCase, r *render.Render) *Something {
	return &Something{
		UseCase: uc,
		Render:  r,
	}
}

// HandlerSomething -
func (c *Something) HandlerSomething(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var data Request
	if err := decoder.Decode(&data); err != nil {
		logrus.WithError(err).Info("decoding")
		c.Render.Text(w, http.StatusBadRequest, "invalid json")
		return
	}

	if data.Info == "" {
		c.Render.Text(w, http.StatusBadRequest, "invalid info")
		return
	}

	resp, err := c.UseCase.DoSomething(data.Info)
	if err != nil {
		c.Render.Text(w, http.StatusBadRequest, "something went wrong")
		return
	}

	c.Render.JSON(w, http.StatusOK, resp)
}

// HandlerSomethingWithRepo -
func (c *Something) HandlerSomethingWithRepo(w http.ResponseWriter, r *http.Request) {
	account := mux.Vars(r)["account"]

	resp, err := c.UseCase.DoSomethingWithRepo(account)
	if err != nil {
		c.Render.Text(w, http.StatusBadRequest, "something went wrong")
		return
	}

	c.Render.JSON(w, http.StatusOK, resp)
}
