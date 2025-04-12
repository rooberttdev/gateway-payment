package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rooberttdev/gateway-payment/internal/repository"
	"github.com/rooberttdev/gateway-payment/internal/service"
	"github.com/rooberttdev/gateway-payment/internal/web/server"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// String de conexão com o banco de dados
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "gateway"),
		getEnv("DB_SSL_MODE", "disable"),
	)	
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()
	
	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(*accountRepository)
	port := getEnv("HTTP_PORT", "8080")
	srv := server.NewServer(*accountService, port)
	srv.ConfigureRoutes()
	
	srv.Start()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	
}
