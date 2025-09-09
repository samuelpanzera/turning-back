# Technology Stack

## Core Technologies
- **Language**: Go 1.23+
- **Web Framework**: Gin (HTTP router and middleware)
- **ORM**: GORM (Object-Relational Mapping)
- **Databases**: 
  - SQLite (development) via modernc.org/sqlite
  - PostgreSQL (production) via gorm.io/driver/postgres
- **Logging**: Uber Zap
- **Validation**: go-playground/validator
- **Configuration**: godotenv for environment variables

## Key Dependencies
- `github.com/gin-gonic/gin` - HTTP web framework
- `gorm.io/gorm` - ORM for database operations
- `go.uber.org/zap` - Structured logging
- `github.com/go-playground/validator/v10` - Request validation
- `github.com/joho/godotenv` - Environment configuration

## Build System & Commands

### Makefile Commands
```bash
# Development
make setup          # Setup development environment
make run-local      # Run with SQLite (development)
make run-docker     # Run with Docker Compose (PostgreSQL)
make dev           # Run with hot reload (requires air)

# Building & Testing
make build         # Build the application
make test          # Run tests
make test-coverage # Run tests with coverage report
make lint          # Run linter (requires golangci-lint)
make format        # Format code with go fmt and goimports

# Database Management
make db-reset      # Reset SQLite database
make db-reset-docker # Reset Docker PostgreSQL database

# Docker Operations
make docker-build  # Build Docker image
make docker-run    # Start with Docker Compose
make docker-stop   # Stop Docker Compose

# API Testing
make test-api      # Test API endpoints with curl
```

### Manual Commands
```bash
# Basic operations
go mod tidy        # Install/update dependencies
go run main.go     # Run application
go build -o bin/app main.go  # Build binary

# Testing
go test -v ./...   # Run all tests
go test -coverprofile=coverage.out ./...  # Test with coverage

# Code quality
go fmt ./...       # Format code
goimports -w .     # Organize imports
```

## Environment Configuration
- `.env.local` - SQLite development setup
- `.env.docker` - Docker PostgreSQL setup  
- `.env.production` - Production PostgreSQL setup
- Environment auto-detection based on `DB_TYPE` and `ENV` variables

## Development Tools
- **Air** - Hot reload for development (`go install github.com/cosmtrek/air@latest`)
- **golangci-lint** - Comprehensive linter
- **goimports** - Import organization (`go install golang.org/x/tools/cmd/goimports@latest`)

## Docker Support
- Multi-stage Dockerfile for optimized production builds
- Docker Compose with PostgreSQL for local development
- Health checks and proper signal handling