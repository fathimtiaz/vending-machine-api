package application

import (
	"context"
	"vending-machine-api/domain"
)

type TransactionRepoFake struct {
	err    error
	result []domain.Product
}

func (f *TransactionRepoFake) GetProducts(ctx context.Context) ([]domain.Product, error) {
	return f.result, f.err
}
