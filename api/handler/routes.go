package handler

import (
	"net/http"
	"vending-machine-api/application"
)

func InitRoutes(
	transactionService *application.TransactionService,
) {
	transactionHndlr := NewTransactionHandler(transactionService)
	http.HandleFunc("/transaction", transactionHndlr.NewTransaction)
}
