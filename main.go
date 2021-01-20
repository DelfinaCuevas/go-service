package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/eiizu/go-service/controller"
	"github.com/eiizu/go-service/router"
	"github.com/eiizu/go-service/service"
	"github.com/eiizu/go-service/usecase"
	"github.com/sirupsen/logrus"
)

const (
	// AppName application name
	AppName = "go-service-demo"
)

func main() {

	logger := logrus.New()

	// Service init
	someService := service.NewSomeService("something")

	// UseCase init
	somethingUC := usecase.NewSomething(someService)
	statusUC := usecase.NewStatus(AppName)
	userUC := usecase.NewUsers()
	bookUC := usecase.NewBooks()

	// Controller init
	somethingC := controller.NewSomething(somethingUC)
	statusC := controller.NewStatus(statusUC)
	userC := controller.NewUsers(userUC) // se inicializa el user controller
	bookC := controller.NewBooks(bookUC)

	// Create router
	r := router.New(somethingC, statusC, userC, bookC)

	// Define stop signal for the end of excecution
	stop := make(chan os.Signal, 1)
	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)

	go func() {
		address := ":8080"
		if err := r.Start(address); err != nil {
			logger.Fatal("something went wrong")
		}
	}()

	<-stop

	logger.Info("shutting down server...")
}
