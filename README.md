# Go CRUD Practice

A simple Go application demonstrating basic CRUD operations with a PostgreSQL database, containerized using Docker.

## Features

- **REST API** built with standard library `net/http`.
- **PostgreSQL** integration for data persistence.
- **Docker & Docker Compose** support for easy deployment.

## Prerequisites

- [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://golang.org/) (optional, for local development)

## Getting Started

The easiest way to run the application is using Docker Compose.

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd go-crud-evo
   ```

2. **Start the services:**
   ```bash
   docker-compose up --build
   ```
   This will start both the Go application (on port 8080) and the PostgreSQL database (on port 5432).

3. **Stop the services:**
   ```bash
   docker-compose down
   ```

## API Endpoints

### 1. Hello World
Checks if the server is running.

- **URL:** `/hello`
- **Method:** `GET`
- **Response:** `hello world!`

### 2. Number Operations
Stores a number and retrieves all stored numbers sorted in ascending order.

- **URL:** `/number`
- **Method:** `POST`
- **Content-Type:** `application/json`
- **Body:**
  ```json
  {
    "number": 42
  }
  ```
- **Response:** JSON array of all numbers.
  ```json
  [1, 5, 42, 100]
  ```

## Project Structure

- `main.go`: Entry point, database connection, and routing.
- `handlers/`: Contains the request handlers and business logic.
- `Dockerfile`: Multi-stage build for the Go application.
- `docker-compose.yml`: Orchestration for the app and database.
