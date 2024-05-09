package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"vending-machine-api/api/handler"
	"vending-machine-api/application"
	"vending-machine-api/repository/sql"
)

func main() {
	connStr := "postgres://postgres:postgres@localhost:6432/database?sslmode=disable"

	sqlRepo, err := sql.NewSQLRepository("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return
	}

	transactionService := application.NewTransactionService(sqlRepo)

	handler.InitRoutes(transactionService)

	fmt.Println("server started at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
