# Step-by-Step Guide: Building `go-crud-evo` (v1)

This guide explains how to build the initial version of the `go-crud-evo` microservice. This version uses a simple, monolithic structure where all logic resides in `main.go`.

## Prerequisites

- Go 1.23+ installed
- Docker & Docker Compose installed
- A code editor (VS Code recommended)

---

## Step 1: Project Initialization

Create a new directory and initialize the Go module.

```bash
mkdir go-crud-evo
cd go-crud-evo
go mod init go-crud-evo
```

## Step 2: Create the Application

We will implement the entire application in a single file `main.go`. This includes database connection, HTTP handling, and business logic.

**File:** `main.go`

```go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type NumberRequest struct {
	Number int `json:"number"`
}

func main() {
	// 1. Get configuration from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	// 2. Connect to the Database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 3. Ensure DB is ready (simple ping)
	if err = db.Ping(); err != nil {
		log.Printf("Warning: Could not ping DB: %v", err)
	}

	// 4. Create the table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS numbers (value INT)`)
	if err != nil {
		log.Fatal(err)
	}

	// 5. Define the HTTP Handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Only allow POST requests
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the JSON request body
		var req NumberRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Insert the number into the database
		_, err = db.Exec("INSERT INTO numbers (value) VALUES ($1)", req.Number)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Retrieve all numbers sorted by value
		rows, err := db.Query("SELECT value FROM numbers ORDER BY value ASC")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var numbers []int
		for rows.Next() {
			var n int
			if err := rows.Scan(&n); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			numbers = append(numbers, n)
		}

		// Respond with the list of numbers
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(numbers)
	})

	// 6. Start the Server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Step 3: Dockerize the Application

Create a `Dockerfile` to build a lightweight container for our Go app.

**File:** `Dockerfile`

```dockerfile
# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
```

## Step 4: Orchestrate with Docker Compose

Define the application and database services.

**File:** `docker-compose.yml`

```yaml
services:
  app:
    build: .
    ports:
      - "8081:8080"
    environment:
      - DB_HOST=db
      - DB_PORT=5432
      - DB_USER=postgres
      - DB_PASSWORD=postgres
      - DB_NAME=numbers_db
    depends_on:
      - db
    restart: on-failure

  db:
    image: postgres:15-alpine
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
      - POSTGRES_DB=numbers_db
    ports:
      - "5433:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

volumes:
  postgres_data:
```

## Step 5: Run It!

1.  **Start the services:**
    ```bash
    docker-compose up --build
    ```

2.  **Test the API:**
    ```bash
    curl -X POST http://localhost:8081/ \
      -H "Content-Type: application/json" \
      -d '{"number": 100}'
    ```
