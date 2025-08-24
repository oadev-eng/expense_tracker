package routes

import (
	"tracker/handler"
	"tracker/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handler.UserHandler, transactionHandler *handler.TransactionHandler, budgetHandler *handler.BudgetHandler) *mux.Router {
	r := mux.NewRouter()

	//public routes
	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/login", userHandler.LoginUser).Methods("POST")

	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// Authenticated user routes
	protected.HandleFunc("/trans", transactionHandler.CreateTransaction).Methods("POST")
	protected.HandleFunc("/gettrans", transactionHandler.GetTransactionbyID).Methods("GET")
	protected.HandleFunc("/budget", budgetHandler.CreateBudget).Methods("POST")
	protected.HandleFunc("/gettransuid", transactionHandler.GetTransactionsbyUserID).Methods("GET")

	return r
}
