package handler

import (
	"net/http"
	"vending-machine-api/application"
	"vending-machine-api/helper"

	httpSwagger "github.com/swaggo/http-swagger"
)

func InitRoutes(
	transactionService *application.TransactionService,
	logger *helper.Logger,
) {
	transactionHndlr := NewTransactionHandler(transactionService, logger)
	http.HandleFunc("/transaction", transactionHndlr.Transaction)

	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
}
