package handler

import (
	"net/http"
	"vending-machine-api/application"
	"vending-machine-api/helper"
)

func InitRoutes(
	transactionService *application.TransactionService,
	logger *helper.Logger,
) {
	transactionHndlr := NewTransactionHandler(transactionService, logger)
	http.HandleFunc("/transaction", transactionHndlr.NewTransaction)
}
