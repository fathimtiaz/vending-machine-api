package handler

import (
	"net/http"
	"vending-machine-api/application"
)

type transactionHandler struct {
	transactionService *application.TransactionService
}

func NewTransactionHandler(
	transactionService *application.TransactionService,
) *transactionHandler {
	return &transactionHandler{transactionService}
}

type NewTransactionReq struct {
	paymentBills []int
}

func (h *transactionHandler) NewTransaction(w http.ResponseWriter, r *http.Request) {
}
