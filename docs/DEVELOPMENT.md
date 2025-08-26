# Guia de Desenvolvimento - Turning Back

## 🚀 Configuração do Ambiente de Desenvolvimento

### Pré-requisitos

1. **Go 1.21+** instalado e configurado
2. **Docker** e **Docker Compose**
3. **Git** para controle de versão
4. **IDE/Editor** recomendado: VS Code com extensão Go

### Configuração Inicial

```bash
# 1. Clone o repositório
git clone https://github.com/samuelpanzera/turning-back.git
cd turning-back

# 2. Copie as variáveis de ambiente
cp .env.example .env

# 3. Instale dependências
go mod download
go mod tidy

# 4. Instale ferramentas de desenvolvimento
make install-tools
```

### Ferramentas de Desenvolvimento

#### Air (Hot Reload)
```bash
# Instalar
go install github.com/cosmtrek/air@latest

# Usar
make dev
# ou
air
```

#### Swagger (Documentação da API)
```bash
# Instalar
go install github.com/swaggo/swag/cmd/swag@latest

# Gerar documentação
make docs
# ou
swag init -g cmd/api/main.go -o ./docs
```

#### Migrate (Migrações de Banco)
```bash
# Instalar
go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Criar migração
make db-migrate-create name=create_users_table

# Executar migrações
make db-migrate-up
```

## 🏗️ Estrutura de Desenvolvimento

### Adicionando uma Nova Feature

#### 1. Criar a Entidade (Domain)
```go
// internal/domain/entities/product.go
package entities

type Product struct {
    ID          uint      `json:"id" gorm:"primarykey"`
    Name        string    `json:"name" gorm:"not null"`
    Description string    `json:"description"`
    Price       float64   `json:"price" gorm:"not null"`
    UserID      uint      `json:"user_id" gorm:"not null"`
    User        User      `json:"user" gorm:"foreignKey:UserID"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

#### 2. Definir Interface do Repositório
```go
// internal/domain/interfaces/product_repository.go
package interfaces

type ProductRepository interface {
    Create(ctx context.Context, product *entities.Product) error
    GetByID(ctx context.Context, id uint) (*entities.Product, error)
    GetByUserID(ctx context.Context, userID uint) ([]*entities.Product, error)
    Update(ctx context.Context, product *entities.Product) error
    Delete(ctx context.Context, id uint) error
    List(ctx context.Context, offset, limit int) ([]*entities.Product, int64, error)
}
```

#### 3. Criar DTOs
```go
// internal/usecases/dto/product_dto.go
package dto

type CreateProductDTO struct {
    Name        string  `json:"name" validate:"required,min=3"`
    Description string  `json:"description"`
    Price       float64 `json:"price" validate:"required,gt=0"`
    UserID      uint    `json:"user_id" validate:"required"`
}

type UpdateProductDTO struct {
    ID          uint    `json:"id" validate:"required"`
    Name        string  `json:"name,omitempty"`
    Description string  `json:"description,omitempty"`
    Price       float64 `json:"price,omitempty"`
}

type ProductResponseDTO struct {
    ID          uint    `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    UserID      uint    `json:"user_id"`
    CreatedAt   string  `json:"created_at"`
    UpdatedAt   string  `json:"updated_at"`
}
```

#### 4. Implementar Use Case (Service)
```go
// internal/usecases/services/product_service.go
package services

type ProductService struct {
    productRepo interfaces.ProductRepository
    logger      *logger.Logger
}

func NewProductService(productRepo interfaces.ProductRepository, logger *logger.Logger) *ProductService {
    return &ProductService{
        productRepo: productRepo,
        logger:      logger,
    }
}

func (s *ProductService) CreateProduct(ctx context.Context, dto dto.CreateProductDTO) (*dto.ProductResponseDTO, error) {
    // Validações de negócio
    if dto.Price <= 0 {
        return nil, errors.New("price must be greater than zero")
    }

    // Criar entidade
    product := &entities.Product{
        Name:        dto.Name,
        Description: dto.Description,
        Price:       dto.Price,
        UserID:      dto.UserID,
    }

    // Persistir
    if err := s.productRepo.Create(ctx, product); err != nil {
        s.logger.Error("failed to create product", "error", err)
        return nil, err
    }

    // Retornar DTO de resposta
    return &dto.ProductResponseDTO{
        ID:          product.ID,
        Name:        product.Name,
        Description: product.Description,
        Price:       product.Price,
        UserID:      product.UserID,
        CreatedAt:   product.CreatedAt.Format(time.RFC3339),
        UpdatedAt:   product.UpdatedAt.Format(time.RFC3339),
    }, nil
}
```

#### 5. Implementar Repositório
```go
// internal/adapters/repositories/product_repository.go
package repositories

type ProductRepository struct {
    db *database.DB
}

func NewProductRepository(db *database.DB) *ProductRepository {
    return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(ctx context.Context, product *entities.Product) error {
    return r.db.WithContext(ctx).Create(product).Error
}

func (r *ProductRepository) GetByID(ctx context.Context, id uint) (*entities.Product, error) {
    var product entities.Product
    err := r.db.WithContext(ctx).Preload("User").First(&product, id).Error
    if err != nil {
        return nil, err
    }
    return &product, nil
}
```

#### 6. Criar Handler
```go
// internal/adapters/handlers/product_handler.go
package handlers

type ProductHandler struct {
    productService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
    return &ProductHandler{
        productService: productService,
    }
}

// @Summary Create a new product
// @Description Create a new product with the provided information
// @Tags products
// @Accept json
// @Produce json
// @Param product body dto.CreateProductDTO true "Product information"
// @Success 201 {object} dto.ProductResponseDTO
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /products [post]
// @Security BearerAuth
func (h *ProductHandler) CreateProduct(c *gin.Context) {
    var req dto.CreateProductDTO
    
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    product, err := h.productService.CreateProduct(c.Request.Context(), req)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(201, product)
}
```

#### 7. Registrar Rotas
```go
// cmd/api/main.go (ou arquivo de rotas separado)
func setupRoutes(router *gin.Engine, handlers *Handlers) {
    v1 := router.Group("/api/v1")
    {
        products := v1.Group("/products")
        {
            products.POST("", handlers.Product.CreateProduct)
            products.GET("/:id", handlers.Product.GetProduct)
            products.PUT("/:id", handlers.Product.UpdateProduct)
            products.DELETE("/:id", handlers.Product.DeleteProduct)
            products.GET("", handlers.Product.ListProducts)
        }
    }
}
```

## 🧪 Testes

### Estrutura de Testes

```
tests/
├── unit/                   # Testes unitários
│   ├── entities/
│   ├── services/
│   └── repositories/
├── integration/            # Testes de integração
│   ├── handlers/
│   └── database/
└── e2e/                   # Testes end-to-end
    └── api/
```

### Exemplo de Teste Unitário

```go
// tests/unit/services/product_service_test.go
package services_test

func TestProductService_CreateProduct(t *testing.T) {
    // Arrange
    mockRepo := &mocks.ProductRepository{}
    logger := logger.New("debug", "console")
    service := services.NewProductService(mockRepo, logger)
    
    dto := dto.CreateProductDTO{
        Name:   "Test Product",
        Price:  99.99,
        UserID: 1,
    }
    
    mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*entities.Product")).Return(nil)
    
    // Act
    result, err := service.CreateProduct(context.Background(), dto)
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, "Test Product", result.Name)
    mockRepo.AssertExpectations(t)
}
```

### Executar Testes

```bash
# Todos os testes
make test

# Testes com cobertura
make test-coverage

# Testes específicos
go test ./internal/usecases/services/...
```

## 📝 Convenções de Código

### Nomenclatura

- **Packages**: lowercase, sem underscores
- **Files**: snake_case
- **Types**: PascalCase
- **Functions/Methods**: PascalCase (públicos), camelCase (privados)
- **Variables**: camelCase
- **Constants**: UPPER_CASE ou PascalCase

### Estrutura de Arquivos

```go
// Ordem recomendada dentro de um arquivo .go
package packagename

import (
    // Standard library
    "context"
    "fmt"
    
    // Third party
    "github.com/gin-gonic/gin"
    
    // Local
    "github.com/samuelpanzera/turning-back/internal/domain/entities"
)

// Constants
const (
    DefaultTimeout = 30 * time.Second
)

// Types
type Service struct {
    // ...
}

// Constructor
func NewService() *Service {
    // ...
}

// Methods
func (s *Service) Method() error {
    // ...
}
```

### Comentários e Documentação

```go
// Package comment
// Package services provides business logic implementations.
package services

// Type comment
// UserService handles user-related business operations.
type UserService struct {
    userRepo interfaces.UserRepository
}

// Constructor comment
// NewUserService creates a new instance of UserService.
func NewUserService(userRepo interfaces.UserRepository) *UserService {
    return &UserService{
        userRepo: userRepo,
    }
}

// Method comment
// CreateUser creates a new user with the provided information.
// It validates the input and ensures email uniqueness.
func (s *UserService) CreateUser(ctx context.Context, dto CreateUserDTO) (*entities.User, error) {
    // Implementation
}
```

## 🔧 Ferramentas Úteis

### VS Code Extensions

- **Go** (oficial)
- **Go Test Explorer**
- **REST Client** (para testar APIs)
- **Docker**
- **GitLens**

### Configuração do VS Code

```json
// .vscode/settings.json
{
    "go.useLanguageServer": true,
    "go.formatTool": "goimports",
    "go.lintTool": "golangci-lint",
    "go.testFlags": ["-v"],
    "go.coverOnSave": true,
    "go.coverageDecorator": {
        "type": "gutter"
    }
}
```

### Makefile Commands

```bash
# Desenvolvimento
make dev          # Hot reload
make build        # Build da aplicação
make run          # Executar aplicação

# Testes
make test         # Todos os testes
make test-unit    # Testes unitários
make test-integration # Testes de integração

# Docker
make docker-build # Build da imagem
make docker-run   # Executar com Docker
make docker-stop  # Parar containers

# Banco de dados
make db-migrate-up    # Executar migrações
make db-migrate-down  # Reverter migrações
make db-migrate-create name=migration_name # Criar migração

# Qualidade de código
make lint         # Linter
make format       # Formatação
make docs         # Gerar documentação
```

## 🚀 Deploy e Produção

### Build para Produção

```bash
# Build otimizado
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

# Com Docker
docker build -t turning-back:latest .
```

### Variáveis de Ambiente para Produção

```bash
ENV=production
PORT=8080
DB_HOST=your-db-host
DB_PORT=5432
DB_USER=your-db-user
DB_PASSWORD=your-secure-password
DB_NAME=turning_back_prod
DB_SSL_MODE=require
JWT_SECRET=your-super-secure-jwt-secret
LOG_LEVEL=info
LOG_FORMAT=json
```

## 📚 Recursos Adicionais

- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Gin Framework Documentation](https://gin-gonic.com/docs/)
- [GORM Documentation](https://gorm.io/docs/)