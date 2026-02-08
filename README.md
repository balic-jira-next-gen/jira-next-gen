# Jira Next-Gen Clone

A high-performance work management platform built with Next.js and Go.

## Prerequisites

-   [Node.js](https://nodejs.org/) (v18+)
-   [Go](https://go.dev/) (v1.22+)
-   [Docker](https://www.docker.com/) & Docker Compose

## Getting Started

### 1. Start the Database
```bash
docker-compose up -d
```
This starts PostgreSQL on port 5432.

### 2. Start the Backend
```bash
cd backend
# Install dependencies
go mod tidy
# Run the server
go run main.go
```
The API will be available at `http://localhost:8080`.

### 3. Start the Frontend
```bash
cd frontend
npm run dev
```
The UI will be available at `http://localhost:3000`.

## Project Structure

-   `/backend`: Go API service (Chi router, pgx driver).
-   `/frontend`: Next.js App Router application.
-   `/docker-compose.yml`: Infrastructure definition.
