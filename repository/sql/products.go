package sql

import (
	"context"
	"vending-machine-api/domain"
)

func (repo *SQLRepository) GetProducts(ctx context.Context) (result []domain.Product, err error) {
	rows, err := repo.db.QueryContext(ctx, `SELECT id, name, price FROM product`)
	if err != nil {
		return
	}

	for rows.Next() {
		var product domain.Product

		if err = rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
			return []domain.Product{}, err
		}

		result = append(result, product)
	}

	return
}
