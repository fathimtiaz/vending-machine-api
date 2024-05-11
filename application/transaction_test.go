package application

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"vending-machine-api/domain"
)

func TestTransactionService_NewTransaction(t *testing.T) {
	type args struct {
		ctx  context.Context
		spec NewTransactionSpec
	}
	tests := []struct {
		name       string
		s          *TransactionService
		args       args
		wantResult domain.Transaction
		wantErr    bool
	}{
		{
			name: "error at GetProducts",
			s: NewTransactionService(&TransactionRepoFake{
				err: errors.New("expected error"),
			}),
			wantErr: true,
		},
		{
			name: "success",
			s: NewTransactionService(&TransactionRepoFake{
				result: []domain.Product{{Name: "A", Price: 2000}},
			}),
			args: args{spec: NewTransactionSpec{PaymentBills: []int{4000}}},
			wantResult: domain.Transaction{
				Items: []domain.TransactionItem{{ProductName: "A", Qty: 2, Price: 4000}},
				Price: 4000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.s.NewTransaction(tt.args.ctx, tt.args.spec)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionService.NewTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("TransactionService.NewTransaction() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
