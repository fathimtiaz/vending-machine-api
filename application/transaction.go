package application

import (
	"context"
	"vending-machine-api/domain"
)

type ITransactionRepo interface {
	GetProducts(ctx context.Context) ([]domain.Product, error)
}

type TransactionService struct {
	repo ITransactionRepo
}

func NewTransactionService(repo ITransactionRepo) *TransactionService {
	return &TransactionService{repo}
}

type NewTransactionSpec struct {
	PaymentBills []int `json:"payment_bills"`
}

func (s NewTransactionSpec) sum() (result int) {
	for _, bill := range s.PaymentBills {
		result += bill
	}

	return
}

func (s *TransactionService) NewTransaction(ctx context.Context, spec NewTransactionSpec) (result domain.Transaction, err error) {
	var products []domain.Product

	if products, err = s.repo.GetProducts(ctx); err != nil {
		return
	}

	result = domain.MakeTransaction(ctx, spec.sum(), products)

	return
}
