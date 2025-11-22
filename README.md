# Go CRUD Evo

A simple, containerized Go microservice demonstrating Clean Architecture principles. This service accepts integer inputs, stores them in a PostgreSQL database, and returns the complete list of stored numbers sorted in ascending order.

## Features

- **REST API**: Single endpoint to process numbers.
- **Clean Architecture**: Separation of concerns into Handler, Service, and Repository layers.
- **PostgreSQL Integration**: Persistent storage for numbers.
- **Dockerized**: Fully containerized application and database using Docker Compose.

## Architecture

The project follows a layered architecture:

1.  **Handler Layer** (`internal/handler`): Handles HTTP requests and responses.
2.  **Service Layer** (`internal/service`): Contains business logic.
3.  **Repository Layer** (`internal/repository`): Manages data access and database interactions.

## Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/HappyR0b0t/go-crud-evo.git
cd go-crud-evo
```

### 2. Run with Docker Compose

Build and start the services:

```bash
docker-compose up --build
```

The application will be available at `http://localhost:8081`.

### 3. Stop the Services

```bash
docker-compose down
```

## API Usage

### Process a Number

**Endpoint:** `POST /`

**Request Body:**

```json
{
  "number": 42
}
```

**Response:**

Returns an array of all stored numbers, sorted in ascending order.

```json
[1, 5, 10, 42]
```

### Example using cURL

```bash
curl -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{"number": 42}'
```

### Example using PowerShell

```powershell
Invoke-RestMethod -Method Post -Uri "http://localhost:8081/" `
  -ContentType "application/json" `
  -Body '{"number": 42}'
```

## Project Structure

```
.
├── docker-compose.yml    # Docker Compose configuration
├── Dockerfile           # Docker build instructions
├── go.mod              # Go module definition
├── main.go             # Application entry point
└── internal
    ├── handler         # HTTP handlers
    ├── repository      # Database access
    └── service         # Business logic
```
