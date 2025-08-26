# Variables
APP_NAME=turning-back
DOCKER_IMAGE=$(APP_NAME):latest
GO_VERSION=1.21

# Colors for output
RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[1;33m
BLUE=\033[0;34m
NC=\033[0m # No Color

.PHONY: help build run test clean docker-build docker-run docker-stop setup dev

# Default target
help: ## Show this help message
	@echo "$(BLUE)$(APP_NAME) - Available commands:$(NC)"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  $(GREEN)%-15s$(NC) %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Setup and Installation
setup: ## Install dependencies and setup development environment
	@echo "$(YELLOW)Setting up development environment...$(NC)"
	go mod download
	go mod tidy
	@echo "$(GREEN)Setup completed!$(NC)"

install-tools: ## Install development tools
	@echo "$(YELLOW)Installing development tools...$(NC)"
	go install github.com/cosmtrek/air@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@echo "$(GREEN)Tools installed!$(NC)"

# Development
dev: ## Run the application in development mode with hot reload
	@echo "$(YELLOW)Starting development server...$(NC)"
	air

run: ## Run the application
	@echo "$(YELLOW)Running application...$(NC)"
	go run cmd/api/main.go

build: ## Build the application
	@echo "$(YELLOW)Building application...$(NC)"
	go build -o bin/$(APP_NAME) cmd/api/main.go
	@echo "$(GREEN)Build completed! Binary: bin/$(APP_NAME)$(NC)"

# Testing
test: ## Run all tests
	@echo "$(YELLOW)Running tests...$(NC)"
	go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "$(YELLOW)Running tests with coverage...$(NC)"
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)Coverage report generated: coverage.html$(NC)"

test-unit: ## Run unit tests only
	@echo "$(YELLOW)Running unit tests...$(NC)"
	go test -v ./tests/unit/...

test-integration: ## Run integration tests only
	@echo "$(YELLOW)Running integration tests...$(NC)"
	go test -v ./tests/integration/...

# Database
db-migrate-up: ## Run database migrations up
	@echo "$(YELLOW)Running database migrations up...$(NC)"
	migrate -path ./migrations -database "postgres://turning_back_user:turning_back_pass@localhost:5432/turning_back_db?sslmode=disable" up

db-migrate-down: ## Run database migrations down
	@echo "$(YELLOW)Running database migrations down...$(NC)"
	migrate -path ./migrations -database "postgres://turning_back_user:turning_back_pass@localhost:5432/turning_back_db?sslmode=disable" down

db-migrate-create: ## Create a new migration file (usage: make db-migrate-create name=migration_name)
	@echo "$(YELLOW)Creating migration: $(name)$(NC)"
	migrate create -ext sql -dir ./migrations $(name)

# Docker
docker-build: ## Build Docker image
	@echo "$(YELLOW)Building Docker image...$(NC)"
	docker build -t $(DOCKER_IMAGE) .
	@echo "$(GREEN)Docker image built: $(DOCKER_IMAGE)$(NC)"

docker-run: ## Run application with Docker Compose
	@echo "$(YELLOW)Starting services with Docker Compose...$(NC)"
	docker-compose up -d
	@echo "$(GREEN)Services started!$(NC)"

docker-run-dev: ## Run application with Docker Compose (development profile)
	@echo "$(YELLOW)Starting services with Docker Compose (dev profile)...$(NC)"
	docker-compose --profile dev up -d
	@echo "$(GREEN)Services started with development tools!$(NC)"

docker-stop: ## Stop Docker Compose services
	@echo "$(YELLOW)Stopping Docker Compose services...$(NC)"
	docker-compose down
	@echo "$(GREEN)Services stopped!$(NC)"

docker-logs: ## Show Docker Compose logs
	docker-compose logs -f

docker-clean: ## Clean Docker images and containers
	@echo "$(YELLOW)Cleaning Docker resources...$(NC)"
	docker-compose down -v --remove-orphans
	docker system prune -f
	@echo "$(GREEN)Docker cleanup completed!$(NC)"

# Code Quality
lint: ## Run linter
	@echo "$(YELLOW)Running linter...$(NC)"
	golangci-lint run

format: ## Format code
	@echo "$(YELLOW)Formatting code...$(NC)"
	go fmt ./...
	goimports -w .

# Documentation
docs: ## Generate API documentation
	@echo "$(YELLOW)Generating API documentation...$(NC)"
	swag init -g cmd/api/main.go -o ./docs
	@echo "$(GREEN)Documentation generated in ./docs$(NC)"

# Cleanup
clean: ## Clean build artifacts
	@echo "$(YELLOW)Cleaning build artifacts...$(NC)"
	rm -rf bin/
	rm -f coverage.out coverage.html
	go clean
	@echo "$(GREEN)Cleanup completed!$(NC)"

# Environment
env-example: ## Create .env.example file
	@echo "$(YELLOW)Creating .env.example...$(NC)"
	@cat > .env.example << 'EOF'
# Application
ENV=development
PORT=8080
APP_NAME=turning-back

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=turning_back_user
DB_PASSWORD=turning_back_pass
DB_NAME=turning_back_db
DB_SSL_MODE=disable

# JWT
JWT_SECRET=your-super-secret-jwt-key
JWT_EXPIRY=24h

# Logging
LOG_LEVEL=info
LOG_FORMAT=json
EOF
	@echo "$(GREEN).env.example created!$(NC)"

# All-in-one commands
init: setup install-tools env-example ## Initialize the project (setup + tools + env)
	@echo "$(GREEN)Project initialized! You can now run 'make dev' to start development.$(NC)"

start: docker-run ## Start all services
	@echo "$(GREEN)All services are running!$(NC)"
	@echo "$(BLUE)API: http://localhost:8080$(NC)"
	@echo "$(BLUE)Adminer: http://localhost:8081$(NC)"

stop: docker-stop ## Stop all services

# Heroku deployment commands
heroku-login: ## Login to Heroku
	@echo "$(YELLOW)Logging into Heroku...$(NC)"
	heroku login

heroku-create: ## Create Heroku app (usage: make heroku-create app=your-app-name)
	@echo "$(YELLOW)Creating Heroku app: $(app)$(NC)"
	heroku create $(app)
	heroku addons:create heroku-postgresql:mini -a $(app)
	@echo "$(GREEN)Heroku app created: $(app)$(NC)"

heroku-config: ## Set Heroku environment variables (usage: make heroku-config app=your-app-name)
	@echo "$(YELLOW)Setting Heroku config vars...$(NC)"
	heroku config:set ENV=production -a $(app)
	heroku config:set LOG_LEVEL=info -a $(app)
	heroku config:set LOG_FORMAT=json -a $(app)
	@echo "$(RED)Don't forget to set JWT_SECRET manually:$(NC)"
	@echo "$(BLUE)heroku config:set JWT_SECRET=your-secret -a $(app)$(NC)"

heroku-deploy: ## Deploy to Heroku using buildpack
	@echo "$(YELLOW)Deploying to Heroku...$(NC)"
	git add .
	git commit -m "Deploy to Heroku" || true
	git push heroku main
	@echo "$(GREEN)Deployed to Heroku!$(NC)"

heroku-deploy-container: ## Deploy to Heroku using container (usage: make heroku-deploy-container app=your-app-name)
	@echo "$(YELLOW)Deploying container to Heroku...$(NC)"
	heroku container:login
	heroku container:push web -a $(app)
	heroku container:release web -a $(app)
	@echo "$(GREEN)Container deployed to Heroku!$(NC)"

heroku-logs: ## Show Heroku logs (usage: make heroku-logs app=your-app-name)
	heroku logs --tail -a $(app)

heroku-setup: heroku-create heroku-config ## Complete Heroku setup (usage: make heroku-setup app=your-app-name)