package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"go-crud-practice/handlers"

	_ "github.com/lib/pq"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello world!")
}

func main() {
	fmt.Println("hello world!")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("DB init error: %v", err)
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

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS numbers (value INT)`)
	if err != nil {
		log.Printf("failed to create a table: %v", err)
	}

	handler := handlers.NewHandler(db)
	http.HandleFunc("/hello", handler.HelloHandler)
	http.HandleFunc("/number", handler.NumberHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
