package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "vending-machine-api/docs"

	_ "github.com/lib/pq"

	"vending-machine-api/api/handler"
	"vending-machine-api/application"
	"vending-machine-api/helper"
	"vending-machine-api/repository/sql"
)

var (
	logger *helper.Logger

	sqlRepo *sql.SQLRepository

	transactionService *application.TransactionService
)

func main() {
	var err error

	logger = &helper.Logger{
		Info:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		Error: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}

	connStr := "postgres://postgres:postgres@localhost:6432/database?sslmode=disable"

	if sqlRepo, err = sql.NewSQLRepository("postgres", connStr); err != nil {
		logger.Error.Fatal("error initializing repository: ", err)
		return
	}

	transactionService = application.NewTransactionService(sqlRepo)

	handler.InitRoutes(transactionService, logger)

	fmt.Println("server started at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
