# Arquitetura do Projeto Turning Back

## Visão Geral da Clean Architecture

Este documento detalha a implementação da Clean Architecture no projeto Turning Back, explicando cada camada e suas responsabilidades.

## Diagrama da Arquitetura

```
┌─────────────────────────────────────────────────────────────┐
│                    External Interfaces                      │
│  (HTTP, CLI, gRPC, Message Queues, etc.)                   │
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│                Interface Adapters                           │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────────────────┐│
│  │  Controllers│ │ Presenters  │ │    Gateways             ││
│  │  (Handlers) │ │             │ │  (Repositories)         ││
│  └─────────────┘ └─────────────┘ └─────────────────────────┘│
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│                 Application Business Rules                  │
│  ┌─────────────────────────────────────────────────────────┐│
│  │              Use Cases (Services)                       ││
│  └─────────────────────────────────────────────────────────┘│
└─────────────────────┬───────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────┐
│              Enterprise Business Rules                      │
│  ┌─────────────────────────────────────────────────────────┐│
│  │                   Entities                              ││
│  └─────────────────────────────────────────────────────────┘│
└─────────────────────────────────────────────────────────────┘
```

## Detalhamento das Camadas

### 1. Entities (Entidades) - `internal/domain/entities/`

**Responsabilidade**: Contém as regras de negócio mais fundamentais e estáveis.

- Estruturas de dados principais do domínio
- Regras de negócio que raramente mudam
- Independentes de qualquer framework ou tecnologia

**Exemplo**:
```go
type User struct {
    ID       uint
    Email    string
    Username string
    // ... outros campos
}

func (u *User) FullName() string {
    return u.FirstName + " " + u.LastName
}
```

### 2. Use Cases (Casos de Uso) - `internal/usecases/`

**Responsabilidade**: Contém as regras de negócio específicas da aplicação.

- Orquestra o fluxo de dados entre entidades
- Implementa casos de uso específicos do sistema
- Define interfaces para repositórios e serviços externos

**Estrutura**:
```
internal/usecases/
├── interfaces/          # Contratos para repositórios e serviços
├── services/           # Implementação dos casos de uso
└── dto/               # Data Transfer Objects
```

### 3. Interface Adapters (Adaptadores) - `internal/adapters/`

**Responsabilidade**: Converte dados entre casos de uso e agentes externos.

- **Handlers**: Controladores HTTP que recebem requisições
- **Repositories**: Implementações de acesso a dados
- **External Services**: Integrações com APIs externas

**Estrutura**:
```
internal/adapters/
├── handlers/           # HTTP handlers (controllers)
├── repositories/       # Implementações de repositórios
└── external/          # Serviços externos (APIs, etc.)
```

### 4. Frameworks & Drivers (Infraestrutura) - `internal/infrastructure/`

**Responsabilidade**: Detalhes de implementação e frameworks.

- Configurações de banco de dados
- Middlewares HTTP
- Configurações da aplicação
- Logging e monitoramento

## Fluxo de Dados

### Requisição HTTP Típica

1. **HTTP Request** → Handler (Adapter)
2. **Handler** → Use Case Service
3. **Use Case** → Repository Interface
4. **Repository** → Database/External Service
5. **Response** ← Reverse path

### Exemplo Prático: Criar Usuário

```go
// 1. Handler recebe requisição HTTP
func (h *UserHandler) CreateUser(c *gin.Context) {
    var req CreateUserRequest
    c.ShouldBindJSON(&req)
    
    // 2. Chama o caso de uso
    user, err := h.userService.CreateUser(c.Request.Context(), req.ToDTO())
    
    // 3. Retorna resposta
    c.JSON(200, user)
}

// 4. Use Case processa a lógica de negócio
func (s *UserService) CreateUser(ctx context.Context, dto CreateUserDTO) (*entities.User, error) {
    // Validações de negócio
    if s.userRepo.ExistsByEmail(ctx, dto.Email) {
        return nil, ErrEmailAlreadyExists
    }
    
    // Cria entidade
    user := &entities.User{
        Email:    dto.Email,
        Username: dto.Username,
    }
    
    // 5. Persiste via repositório
    return s.userRepo.Create(ctx, user)
}
```

## Princípios Seguidos

### 1. Dependency Inversion Principle (DIP)
- Camadas internas não dependem de camadas externas
- Dependências apontam para dentro (direção das abstrações)

### 2. Single Responsibility Principle (SRP)
- Cada camada tem uma responsabilidade específica
- Separação clara entre lógica de negócio e infraestrutura

### 3. Open/Closed Principle (OCP)
- Extensível para novos recursos
- Fechado para modificações em código existente

### 4. Interface Segregation Principle (ISP)
- Interfaces pequenas e específicas
- Clientes não dependem de métodos que não usam

## Vantagens da Implementação

### 1. Testabilidade
- Cada camada pode ser testada independentemente
- Fácil criação de mocks e stubs
- Testes unitários focados em lógica de negócio

### 2. Manutenibilidade
- Mudanças em uma camada não afetam outras
- Código organizado e previsível
- Fácil localização de bugs

### 3. Flexibilidade
- Troca de frameworks sem afetar lógica de negócio
- Múltiplas interfaces (HTTP, CLI, gRPC)
- Diferentes implementações de repositórios

### 4. Escalabilidade
- Estrutura que cresce de forma organizada
- Fácil adição de novos recursos
- Separação clara de responsabilidades

## Padrões Utilizados

### 1. Repository Pattern
```go
type UserRepository interface {
    Create(ctx context.Context, user *entities.User) error
    GetByID(ctx context.Context, id uint) (*entities.User, error)
    // ... outros métodos
}
```

### 2. Dependency Injection
```go
type UserService struct {
    userRepo UserRepository
    logger   Logger
}

func NewUserService(userRepo UserRepository, logger Logger) *UserService {
    return &UserService{
        userRepo: userRepo,
        logger:   logger,
    }
}
```

### 3. DTO (Data Transfer Object)
```go
type CreateUserDTO struct {
    Email     string `json:"email" validate:"required,email"`
    Username  string `json:"username" validate:"required,min=3"`
    FirstName string `json:"first_name" validate:"required"`
    LastName  string `json:"last_name" validate:"required"`
}
```

## Considerações de Performance

### 1. Database Queries
- Use de índices apropriados
- Lazy loading quando necessário
- Paginação em listagens

### 2. Caching
- Cache em nível de aplicação
- Redis para sessões e dados temporários
- Cache de queries frequentes

### 3. Monitoring
- Logs estruturados
- Métricas de performance
- Tracing de requisições

## Próximos Passos

1. **Implementar casos de uso básicos** (CRUD de usuários)
2. **Adicionar autenticação e autorização**
3. **Implementar testes automatizados**
4. **Configurar CI/CD pipeline**
5. **Adicionar documentação da API (Swagger)**
6. **Implementar monitoramento e observabilidade**