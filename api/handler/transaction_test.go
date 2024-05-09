package handler

import (
	"testing"
	"vending-machine-api/application"
)

func TestNewTransactionReq_validate(t *testing.T) {
	tests := []struct {
		name    string
		r       *NewTransactionReq
		wantErr bool
	}{
		{
			r: &NewTransactionReq{
				application.NewTransactionSpec{
					PaymentBills: []int{1000},
				},
			},
			wantErr: true,
		},
		{
			r: &NewTransactionReq{
				application.NewTransactionSpec{
					PaymentBills: []int{2000, 5000},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.validate(); (err != nil) != tt.wantErr {
				t.Errorf("NewTransactionReq.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
