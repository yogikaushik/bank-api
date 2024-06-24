package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	router "github.com/yogikaushik/bank-api/http"
	"github.com/yogikaushik/bank-api/repository"
	"github.com/yogikaushik/bank-api/routes"
	"github.com/yogikaushik/bank-api/service"

	"github.com/yogikaushik/bank-api/controller"

	_ "github.com/go-sql-driver/mysql"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/yogikaushik/bank-api/docs"
)

// @title Bank API
// @version 1.0
// @description This is a sample server for a bank API.
// @host localhost:8082
// @BasePath /
func main() {
	// Database connection
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	defer db.Close()

	// Repositories
	accountRepo := repository.NewAccountRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	// Services
	accountService := service.NewAccountService(accountRepo)
	transactionService := service.NewTransactionService(transactionRepo)

	// Controllers
	accountController := controller.NewAccountController(accountService)
	transactionController := controller.NewTransactionController(transactionService)

	// Router
	router := router.NewRouter()
	routes.RegisterRoutes(router, accountController, transactionController)

	// Swagger route
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	// Run the server
	fmt.Println("Server running at http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", router))

}
