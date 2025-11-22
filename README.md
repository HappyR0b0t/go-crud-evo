# go-crud-evo (v1)

A simple, monolithic Go microservice demonstrating basic CRUD operations with PostgreSQL. This version implements all logic in a single `main.go` file, serving as a baseline for further evolution.

## Features

- **Simple Architecture**: All logic contained in one file for easy understanding.
- **REST API**: Accepts integer inputs via HTTP POST.
- **PostgreSQL**: Persists data in a relational database.
- **Dockerized**: Ready to run with Docker Compose.

## Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Getting Started

### 1. Run with Docker Compose

Build and start the application and database:

```bash
docker-compose up --build
```

The application will be available at `http://localhost:8081`.

### 2. Stop the Services

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

### Testing with cURL

```bash
curl -X POST http://localhost:8081/ \
  -H "Content-Type: application/json" \
  -d '{"number": 42}'
```

### Testing with PowerShell

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
├── main.go             # Application entry point (Monolithic)
└── README.md           # Project documentation
```
