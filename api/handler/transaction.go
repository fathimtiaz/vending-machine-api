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
			return errors.New("invalid payment bills denomination")
		}
	}

	return nil
}

func (h *transactionHandler) Transaction(w http.ResponseWriter, r *http.Request) {
	var rw *ResponseWriter = &ResponseWriter{w: w}

	if r.Method != http.MethodPost {
		rw.httpStatus = http.StatusMethodNotAllowed
		rw.WriteResponse(&Response{
			Message: "method not allowed",
		})
		return
	}

	if r.Method == http.MethodPost {
		h.NewTransaction(rw, r)
		return
	}
}

// NewTransaction
//
//	@Summary		Create a New Transaction
//	@Description	takes array of bills in payment_bills body field and calculate transaction items
//	@Accept			json
//	@Produce		json
//	@Param			payment_bills	body		[]int			true	"int array: allowed 2000 & 5000"
//	@Success		200		{object}	string			"ok"
//	@Failure		400		{object}	string			"invalid parameters"
//	@Failure		500		{object}	string			"internal error"
//	@Router			/transaction [post]
func (h *transactionHandler) NewTransaction(rw *ResponseWriter, r *http.Request) {
	var req NewTransactionReq
	var resp Response
	var err error
	var transaction domain.Transaction

	defer rw.WriteResponse(&resp)

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error.Printf("error decoding request body: %+v", err)
		resp.Message = "invalid request"
		rw.httpStatus = http.StatusBadRequest
		return
	}

	if err = req.validate(); err != nil {
		resp.Message = err.Error()
		rw.httpStatus = http.StatusBadRequest
		return
	}

	if transaction, err = h.transactionService.NewTransaction(r.Context(), req.NewTransactionSpec); err != nil {
		h.logger.Error.Printf("error creating new transaction: %+v", err)
		resp.Message = "internal error"
		rw.httpStatus = http.StatusInternalServerError
		return
	}

	rw.httpStatus = http.StatusOK
	resp.Message = "success"
	resp.Data = transaction
}
