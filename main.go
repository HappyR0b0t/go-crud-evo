package main

import (
	"database/sql"
	"fmt"
	"go-crud-evo/internal/handler"
	"go-crud-evo/internal/repository"
	"go-crud-evo/internal/service"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var pingErr error
	for i := 0; i < 10; i++ {
		pingErr = db.Ping()
		if pingErr == nil {
			log.Println("Successfully connected to the database")
			break
		}
		log.Printf("Attempt %d: failed to ping db, retrying in 2s...", i+1)
		time.Sleep(2 * time.Second)
	}

	if pingErr != nil {
		log.Fatalf("failed to ping db after multiple attempts: %v", pingErr)
	}

	// Initialize layers
	// 1. Repository
	numberRepo := repository.NewPostgresNumberRepository(db)

	// 2. Service
	numberService := service.NewDefaultNumberService(numberRepo)

	// 3. Handler
	numberHandler := handler.NewNumberHandler(numberService)

	// Ensure table exists (could be moved to migration tool, but keeping here for simplicity as per original)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS numbers (value INT)`)
	if err != nil {
		log.Fatal(err)
	}

	// Register routes
	http.HandleFunc("/", numberHandler.Handle)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
