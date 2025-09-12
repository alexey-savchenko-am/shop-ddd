# Shop DDD Example

This project is an example implementation of Domain-Driven Design (DDD) principles in Go.  
It demonstrates the usage of CQRS (Command Query Responsibility Segregation), layered architecture, and integration with PostgreSQL.

## Features
- Domain model for `Product`
- CQRS pattern with separate commands and queries
- HTTP API with Swagger documentation
- PostgreSQL integration via GORM
- Docker Compose setup for local development

## Project Structure
├── cmd/
│ └── shop/ # Application entry point (main.go)
├── internal/
│ ├── application/ # Use cases (commands and queries)
│ │ └── product/
│ ├── domain/ # Domain layer
│ │ ├── common/ # Shared types (e.g., Money)
│ │ └── product/ # Product aggregate
│ ├── infrastructure/ # Database, repositories, external integrations
│ │ └── postgres/
│ └── interfaces/ # HTTP handlers, REST API
│ └── http/
│ └── product/
├── docs/ # Swagger auto-generated files
├── docker-compose.yml # Local development with PostgreSQL
├── go.mod
├── go.sum
└── README.md


## Requirements
- Go 1.22+
- Docker and Docker Compose

## Setup

### Run locally
```bash
go mod tidy
go run ./cmd/shop

docker-compose up --build

