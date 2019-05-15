package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/eiizu/go-service/controller"
	"github.com/eiizu/go-service/router"
	"github.com/eiizu/go-service/repository"
	"github.com/eiizu/go-service/usecase"

	"github.com/codegangsta/negroni"
	"github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)


const (
	// AppName application name
	AppName = "go-service-demo"
)

func main() {

	repo := repository.New()

	logger := logrus.New()
	myRender := render.New()

	// UseCase init
	somethingUC := usecase.NewSomething(repo)
	statusUC := usecase.NewStatus(AppName)

	// Controller init
	somethingC := controller.NewSomething(somethingUC, myRender)
	statusC := controller.NewStatus(statusUC, myRender)

	// Create router
	r := router.New(somethingC, statusC)

	n := negroni.Classic()
	n.UseHandler(r)

	// Define stop signal for the end of excecution
	stop := make(chan os.Signal, 1)
	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)

	go func() {
		address := "localhost:7071"
		logger.WithField("Address", address).Info("Starting server")
		if err := http.ListenAndServe(address, n); err != nil {
			logger.Fatal("something went wrong")
		}
	}()

	<-stop

	logger.Info("Shutting down server...")
}
