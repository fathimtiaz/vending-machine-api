package main

import (
	"log"
	"vending-machine-api/api/handler"
	"vending-machine-api/application"
	"vending-machine-api/repository/sql"
)

func main() {
	connStr := ""

	sqlRepo, err := sql.NewSQLRepository("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return
	}

	transactionService := application.NewTransactionService(sqlRepo)

	handler.InitRoutes(transactionService)
}
