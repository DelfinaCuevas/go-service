package router

import (
	"net/http"

	"github.com/gorilla/mux"
)


// StatusController -
type StatusController interface {
	HandlerStatusz(w http.ResponseWriter, r *http.Request)
	HandlerHealthz(w http.ResponseWriter, r *http.Request)
}

// SomethingController -
type SomethingController interface {
	HandlerSomething(w http.ResponseWriter, r *http.Request)
	HandlerSomethingWithRepo(w http.ResponseWriter, r *http.Request)
}

// New -
func New(somethingC SomethingController, statusC StatusController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/statusz", statusC.HandlerStatusz).Methods(http.MethodGet)
	r.HandleFunc("/healthz", statusC.HandlerHealthz).Methods(http.MethodGet)

	r.HandleFunc("/account/{account}", somethingC.HandlerSomethingWithRepo).Methods(http.MethodGet)
	r.HandleFunc("/operation", somethingC.HandlerSomething).Methods(http.MethodPost)

	return r
}
