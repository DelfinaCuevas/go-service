package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/eiizu/go-service/controller"
	"github.com/eiizu/go-service/router"
	"github.com/eiizu/go-service/service"
	"github.com/eiizu/go-service/store"
	"github.com/eiizu/go-service/usecase"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const (
	// AppName application name
	AppName = "go-service-demo"
)

func main() {

	logger := logrus.New()

	//postgres://postgres:Alma2097@localhost:5432/Library?sslmode=disable

	db, err := store.New("user=postgres password=Alma2097 dbname=Library sslmode=disable")
	if err != nil {
		logger.Fatal("Something went wrong")
	}
	// Service init
	someService := service.NewSomeService("something")

	// UseCase init
	somethingUC := usecase.NewSomething(someService)
	statusUC := usecase.NewStatus(AppName)
	userUC := usecase.NewUsers(db)
	bookUC := usecase.NewBooks(db)
	loanUC := usecase.NewLoans(db)

	// Controller init
	somethingC := controller.NewSomething(somethingUC)
	statusC := controller.NewStatus(statusUC)
	userC := controller.NewUsers(userUC) // se inicializa el user controller
	bookC := controller.NewBooks(bookUC)
	loanC := controller.NewLoans(loanUC)

	// Create router
	r := router.New(somethingC, statusC, userC, bookC, loanC)

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
