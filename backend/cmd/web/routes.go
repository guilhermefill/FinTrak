package web

import (
	"github.com/gorilla/mux"
	"github.com/guilhermefill/FinTrak/pkg/handlers"
)

var RoutesRegistrar = func(r *mux.Router) {
	r.HandleFunc("/accounts", handlers.GetAccounts).Methods("GET")
	r.HandleFunc("/accounts/{id}", handlers.GetAccountByID).Methods("GET")
	r.HandleFunc("/accounts/create", handlers.CreateAccount).Methods("POST")

	r.HandleFunc("/transfers", handlers.GetTransfers).Methods("GET")
	r.HandleFunc("/transfers/{id}", handlers.GetTransferByID).Methods("GET")
	r.HandleFunc("/transfers/create", handlers.CreateTransfer).Methods("POST")

	r.HandleFunc("/incomes", handlers.GetIncomes).Methods("GET")
	r.HandleFunc("/incomes/{id}", handlers.GetIncomeByID).Methods("GET")
	r.HandleFunc("/incomes/create", handlers.CreateIncome).Methods("POST")

	r.HandleFunc("/expenses", handlers.GetExpenses).Methods("GET")
	r.HandleFunc("/expenses/{id}", handlers.GetExpenseByID).Methods("GET")
	r.HandleFunc("/expenses/create", handlers.CreateExpense).Methods("POST")
	r.HandleFunc("/expenses/categories", handlers.GetCategories).Methods("GET")
	r.HandleFunc("/expenses/categories/create", handlers.CreateCategory).Methods("POST")
}
