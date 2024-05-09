package domain

import "context"

type Transaction struct {
	Items []TransactionItem `json:"items"`
	Price int
}

type TransactionItem struct {
	ProductId   int64  `json:"product_id"`
	ProductName string `json:"product_name"`
	Qty         int    `json:"qty"`
	Price       int    `json:"price"`
}

func NewTransaction(ctx context.Context, billsSum int, products []Product) (result Transaction) {
	for _, product := range products {
		qty := billsSum / product.Price

		if qty > 0 {
			item := TransactionItem{
				ProductId:   product.Id,
				ProductName: product.Name,
				Qty:         qty,
				Price:       qty * product.Price,
			}

			result.Items = append(result.Items, item)
			result.Price += item.Price
		}

		billsSum %= product.Price
	}

	return
}
