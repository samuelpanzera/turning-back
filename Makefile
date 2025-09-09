# Turning Back - Hexagonal Architecture Makefile

# Variables
APP_NAME=turning-back-hexagonal
MAIN_PATH=./main.go
BUILD_DIR=./bin

# Colors for output
GREEN=\033[0;32m
YELLOW=\033[1;33m
RED=\033[0;31m
NC=\033[0m # No Color

.PHONY: help build run clean test deps dev docker-build docker-run

# Default target
help: ## Show this help message
	@echo "$(GREEN)Turning Back - Hexagonal Architecture$(NC)"
	@echo "Available commands:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  $(YELLOW)%-15s$(NC) %s\n", $$1, $$2}' $(MAKEFILE_LIST)

deps: ## Install dependencies
	@echo "$(GREEN)Installing dependencies...$(NC)"
	go mod download
	go mod tidy

build: ## Build the application
	@echo "$(GREEN)Building application...$(NC)"
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)
	@echo "$(GREEN)Build completed: $(BUILD_DIR)/$(APP_NAME)$(NC)"

run: ## Run the application
	@echo "$(GREEN)Running application...$(NC)"
	go run $(MAIN_PATH)

dev: ## Run in development mode with hot reload (requires air)
	@echo "$(GREEN)Starting development server...$(NC)"
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "$(RED)Air not found. Install with: go install github.com/cosmtrek/air@latest$(NC)"; \
		echo "$(YELLOW)Running without hot reload...$(NC)"; \
		go run $(MAIN_PATH); \
	fi

test: ## Run tests
	@echo "$(GREEN)Running tests...$(NC)"
	go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "$(GREEN)Running tests with coverage...$(NC)"
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)Coverage report generated: coverage.html$(NC)"

clean: ## Clean build artifacts
	@echo "$(GREEN)Cleaning build artifacts...$(NC)"
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html
	go clean

lint: ## Run linter (requires golangci-lint)
	@echo "$(GREEN)Running linter...$(NC)"
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "$(RED)golangci-lint not found. Install from: https://golangci-lint.run/usage/install/$(NC)"; \
	fi

format: ## Format code
	@echo "$(GREEN)Formatting code...$(NC)"
	go fmt ./...
	@if command -v goimports > /dev/null; then \
		goimports -w .; \
	else \
		echo "$(YELLOW)goimports not found. Install with: go install golang.org/x/tools/cmd/goimports@latest$(NC)"; \
	fi

docker-build: ## Build Docker image
	@echo "$(GREEN)Building Docker image...$(NC)"
	docker build -t $(APP_NAME):latest .

docker-run: ## Run with Docker Compose
	@echo "$(GREEN)Starting with Docker Compose...$(NC)"
	docker-compose up -d

docker-stop: ## Stop Docker Compose
	@echo "$(GREEN)Stopping Docker Compose...$(NC)"
	docker-compose down

docker-logs: ## Show Docker logs
	docker-compose logs -f

install-tools: ## Install development tools
	@echo "$(GREEN)Installing development tools...$(NC)"
	go install github.com/cosmtrek/air@latest
	go install golang.org/x/tools/cmd/goimports@latest
	@echo "$(GREEN)Tools installed successfully$(NC)"

setup: deps install-tools ## Setup development environment
	@echo "$(GREEN)Setting up development environment...$(NC)"
	@if [ ! -f .env ]; then \
		cp .env.example .env; \
		echo "$(YELLOW)Created .env file from .env.example$(NC)"; \
	fi
	@mkdir -p data
	@echo "$(GREEN)Development environment ready!$(NC)"

# Environment-specific commands
run-local: ## Run with SQLite (local development)
	@echo "$(GREEN)Running with SQLite (local development)...$(NC)"
	@cp .env.local .env
	go run $(MAIN_PATH)

run-docker: ## Run with Docker Compose (PostgreSQL)
	@echo "$(GREEN)Starting with Docker Compose (PostgreSQL)...$(NC)"
	docker-compose up --build

run-production: ## Run in production mode
	@echo "$(GREEN)Running in production mode...$(NC)"
	@cp .env.production .env
	go run $(MAIN_PATH)

# Database commands
db-migrate: ## Run database migrations (if implemented)
	@echo "$(GREEN)Running database migrations...$(NC)"
	go run $(MAIN_PATH) migrate

db-reset: ## Reset database (remove SQLite file)
	@echo "$(GREEN)Resetting database...$(NC)"
	rm -f ./data/turning_back.db
	@echo "$(GREEN)Database reset completed$(NC)"

db-reset-docker: ## Reset Docker PostgreSQL database
	@echo "$(GREEN)Resetting Docker PostgreSQL database...$(NC)"
	docker-compose down -v
	docker-compose up -d postgres
	@echo "$(GREEN)Docker database reset completed$(NC)"

# API testing
test-api: ## Test API endpoints
	@echo "$(GREEN)Testing API endpoints...$(NC)"
	@echo "$(YELLOW)Health check:$(NC)"
	@curl -s http://localhost:8080/health | jq . || echo "API not running or jq not installed"
	@echo "\n$(YELLOW)Ping:$(NC)"
	@curl -s http://localhost:8080/api/v1/ping | jq . || echo "API not running or jq not installed"

# Help with common tasks
info: ## Show project information
	@echo "$(GREEN)Project Information:$(NC)"
	@echo "Name: $(APP_NAME)"
	@echo "Go version: $(shell go version)"
	@echo "Main file: $(MAIN_PATH)"
	@echo "Build directory: $(BUILD_DIR)"
	@echo ""
	@echo "$(GREEN)Useful commands:$(NC)"
	@echo "  make setup     - Setup development environment"
	@echo "  make dev       - Start development server"
	@echo "  make test-api  - Test API endpoints"
	@echo "  make help      - Show all available commands"