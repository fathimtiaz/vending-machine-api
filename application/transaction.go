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
