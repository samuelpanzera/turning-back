# Project Structure & Architecture

## Hexagonal Architecture Implementation

This project follows the Hexagonal Architecture (Ports & Adapters) pattern with clear separation between core business logic and external concerns.

## Directory Structure

```
├── adapter/                    # External adapters
│   ├── input/                  # Inbound adapters (HTTP, CLI, etc.)
│   │   ├── controller/         # HTTP controllers
│   │   ├── converter/          # Domain → Response converters
│   │   ├── model/              # Request/Response models
│   │   │   ├── request/        # HTTP request DTOs
│   │   │   └── response/       # HTTP response DTOs
│   │   └── routes/             # Route configuration
│   └── output/                 # Outbound adapters (DB, APIs, etc.)
│       ├── converter/          # Domain ↔ Entity converters
│       ├── model/entity/       # Database entities
│       └── repository/         # Repository implementations
├── application/                # Core business logic (hexagon center)
│   ├── domain/                 # Business domain objects
│   ├── port/                   # Interface definitions
│   │   ├── input/              # Inbound port interfaces (use cases)
│   │   └── output/             # Outbound port interfaces (repositories)
│   └── services/               # Business logic implementations
├── configuration/              # Cross-cutting concerns
│   ├── database/               # Database connection setup
│   ├── logger/                 # Logging configuration
│   ├── rest_errors/            # Error handling
│   └── validation/             # Input validation
└── data/                       # Local database files (SQLite)
```

## Architecture Layers

### Core (Application Layer)
- **Domain**: Pure business objects with no external dependencies
- **Ports**: Interface contracts for input (use cases) and output (repositories)
- **Services**: Business logic implementation, orchestrates domain objects

### Adapters Layer
- **Input Adapters**: Handle external requests (HTTP controllers, CLI commands)
- **Output Adapters**: Handle external resources (database repositories, API clients)

### Configuration Layer
- Cross-cutting concerns like logging, validation, error handling
- Database connection management
- Environment-specific configurations

## Naming Conventions

### Files & Packages
- Use snake_case for file names: `create_orcamento_controller.go`
- Package names should be lowercase, single word when possible
- Interface files end with the domain concept: `orcamento_use_case.go`

### Go Code Conventions
- Interfaces end with `Interface`: `OrcamentoControllerInterface`
- Domain objects end with `Domain`: `OrcamentoDomain`
- Entities end with `Entity`: `OrcamentoEntity`
- Requests end with `Request`: `OrcamentoRequest`
- Responses end with `Response`: `OrcamentoResponse`

### Dependency Flow
```
HTTP Request → Controller → Use Case → Repository → Database
     ↓              ↓           ↓           ↓
  Input Adapter → Port → Domain Logic → Port → Output Adapter
```

## Key Principles
1. **Dependency Inversion**: Core depends on abstractions, not implementations
2. **Single Responsibility**: Each layer has a clear, focused purpose
3. **Interface Segregation**: Small, focused interfaces for each port
4. **Separation of Concerns**: Business logic isolated from infrastructure
5. **Testability**: Easy mocking through interface-based design

## File Organization Rules
- Controllers handle HTTP concerns only (parsing, validation, responses)
- Services contain pure business logic
- Repositories handle data persistence only
- Converters transform between layer-specific models
- Domain objects should have no external dependencies