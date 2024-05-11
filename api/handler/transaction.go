package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"vending-machine-api/application"
	"vending-machine-api/domain"
	"vending-machine-api/helper"
)

type transactionHandler struct {
	transactionService *application.TransactionService
	logger             *helper.Logger
}

func NewTransactionHandler(
	transactionService *application.TransactionService,
	logger *helper.Logger,
) *transactionHandler {
	return &transactionHandler{transactionService, logger}
}

type NewTransactionReq struct {
	application.NewTransactionSpec
}

func (r *NewTransactionReq) validate() error {
	if len(r.PaymentBills) <= 0 {
		return errors.New("payment_bills array cannot be empty")
	}

	for _, bill := range r.PaymentBills {
		if !(bill == 2000 || bill == 5000) {
			return errors.New("invalid denomination")
		}
	}

	return nil
}

type NewTransactionRes struct {
	Status string             `json:"status"`
	Data   domain.Transaction `json:"data"`
}

func (h *transactionHandler) NewTransaction(w http.ResponseWriter, r *http.Request) {
	var req NewTransactionReq
	var err error
	var transaction domain.Transaction

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error.Printf("error decoding request body: %+v", err)
		return
	}

	if transaction, err = h.transactionService.NewTransaction(r.Context(), req.NewTransactionSpec); err != nil {
		h.logger.Error.Printf("error creating new transaction: %+v", err)
		return
	}

	if err = json.NewEncoder(w).Encode(transaction); err != nil {
		h.logger.Error.Printf("error encoding response: %+v", err)
		return
	}
}
