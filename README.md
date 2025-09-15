# Shop DDD Example

This project is an example implementation of Domain-Driven Design (DDD) principles in Go.  
It demonstrates the usage of CQRS (Command Query Responsibility Segregation), layered architecture, and integration with PostgreSQL.

## Features
- Domain model for `Product`
- CQRS pattern with separate commands and queries
- HTTP API with Swagger documentation
- PostgreSQL integration via GORM
- Docker Compose setup for local development


## Requirements
- Go 1.22+
- Docker and Docker Compose.

## Setup

### Run locally
```bash
go mod tidy
go run ./cmd/shop

docker-compose up --build

