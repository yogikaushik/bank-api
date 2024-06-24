package routes

import (
	"github.com/yogikaushik/bank-api/controller"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, accountController *controller.AccountController, transactionController *controller.TransactionController) {
	r.HandleFunc("/accounts", accountController.CreateAccount).Methods("POST")
	r.HandleFunc("/accounts/{id}", accountController.GetAccount).Methods("GET")
	r.HandleFunc("/transactions", transactionController.CreateTransaction).Methods("POST")
}
