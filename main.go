package main

import (
	"fmt"
	"net/http"
	"tracker/config"
	"tracker/database"
	"tracker/handler"
	"tracker/repository"
	"tracker/routes"
	"tracker/service"
)

/*
create webserver
register handler
login handler
createTransaction handler
*/

func main() {
	//load up variables
	config.LoadEnv()

	// connect to database
	database.ConnectDB()

	// initialise the repo
	userRepo := &repository.UserRepo{}
	transactionRepo := &repository.TransactionRepo{}
	budgetRepo := &repository.BudgetRepo{}

	// initialise the service
	userService := &service.UserService{Repo: userRepo}
	transactionService := &service.TransactionService{Repo: transactionRepo}
	budgetService := &service.BudgetService{Repo: budgetRepo}

	// initialise the handler
	userHandler := &handler.UserHandler{Service: userService}
	transactionHandler := &handler.TransactionHandler{Service: transactionService}
	budgetHandler := &handler.BudgetHandler{Service: budgetService}

	// define routes
	router := routes.SetupRouter(userHandler, transactionHandler, budgetHandler)

	// start the server
	fmt.Println("server is running on localhost:8080...")
	http.ListenAndServe(":8080", router)
	
}
