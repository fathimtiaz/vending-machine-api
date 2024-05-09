package domain

import (
	"context"
	"reflect"
	"testing"
)

func TestMakeTransaction(t *testing.T) {
	type args struct {
		ctx      context.Context
		billsSum int
		products []Product
	}
	tests := []struct {
		name       string
		args       args
		wantResult Transaction
	}{
		{
			args: args{
				billsSum: 1000,
				products: []Product{
					{Name: "A", Price: 220},
					{Name: "B", Price: 20},
				},
			},
			wantResult: Transaction{
				Items: []TransactionItem{
					{ProductName: "A", Qty: 4, Price: 880},
					{ProductName: "B", Qty: 6, Price: 120},
				},
				Price: 1000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := MakeTransaction(tt.args.ctx, tt.args.billsSum, tt.args.products); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("MakeTransaction() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
