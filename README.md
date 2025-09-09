# Turning Back - Hexagonal Architecture

## 🏗️ Arquitetura Hexagonal (Ports & Adapters)

Este projeto implementa a **Arquitetura Hexagonal** (também conhecida como Ports & Adapters), criada por Alistair Cockburn. Esta arquitetura permite que a aplicação seja igualmente dirigida por usuários, programas, testes automatizados ou scripts em lote, e seja desenvolvida e testada isoladamente de seus dispositivos e bancos de dados de tempo de execução.

## 📁 Estrutura do Projeto

```
hexagonal/
├── adapter/
│   ├── input/                          # Adaptadores de Entrada
│   │   ├── controller/                 # Controllers HTTP
│   │   │   ├── create_orcamento_controller.go
│   │   ├── converter/                  # Conversores Domain → Response
│   │   ├── model/                      # Modelos de Request/Response
│   │   │   ├── request/
│   │   │   └── response/
│   │   └── routes/                     # Configuração de Rotas
│   └── output/                         # Adaptadores de Saída
│       ├── converter/                  # Conversores Domain ↔ Entity
│       ├── model/entity/               # Entidades de Banco
│       │   └── orcamento_entity.go
│       └── repository/                 # Implementações de Repository
│           ├── create_orcamento_repository.go
├── application/                        # Núcleo da Aplicação
│   ├── domain/                         # Objetos de Domínio
│   │   └── orcamento_domain.go
│   ├── port/                           # Portas (Interfaces)
│   │   ├── input/                      # Portas de Entrada
│   │   │   └── orcamento_use_case.go
│   │   └── output/                     # Portas de Saída
│   │       └── orcamento_port.go
│   └── services/                       # Casos de Uso (Business Logic)
│       └── create_orcamento.go
├── configuration/                      # Configurações da Aplicação
│   ├── database/                       # Conexão com Banco
│   │   └── database_connection.go
│   ├── logger/                         # Configuração de Logs
│   ├── rest_errors/                    # Tratamento de Erros
│   └── validation/                     # Validações
│       └── validate.go
├── .env.example
├── go.mod
├── main.go
└── README.md
```

## 🎯 Conceitos da Arquitetura Hexagonal

### Núcleo (Application)
- **Domain**: Objetos de negócio puros, sem dependências externas
- **Ports**: Interfaces que definem contratos de entrada e saída
- **Services**: Implementação da lógica de negócio (casos de uso)

### Adaptadores (Adapters)
- **Input Adapters**: Recebem requisições externas (HTTP, CLI, etc.)
- **Output Adapters**: Implementam comunicação com recursos externos (DB, APIs, etc.)

### Fluxo de Dados
```
HTTP Request → Controller → Use Case → Repository → Database
     ↓              ↓           ↓           ↓
  Input Adapter → Port → Domain Logic → Port → Output Adapter
```

## 🚀 Como Executar

### 1. Pré-requisitos
- Go 1.23+
- Docker & Docker Compose (para ambiente PostgreSQL)

### 2. Configuração
```bash
# Clone e configure
git clone <repository>
cd turning-back

# Instale dependências
go mod tidy
```

### 3. Ambientes Disponíveis

#### 🔧 Desenvolvimento Local (SQLite)
Ideal para desenvolvimento rápido e testes locais.
```bash
# Usando Makefile
make run-local

# Ou manualmente
cp .env.local .env
go run main.go
```

#### 🐳 Desenvolvimento Docker (PostgreSQL)
Simula o ambiente de produção localmente.
```bash
# Usando Makefile
make run-docker

# Ou manualmente
docker-compose up --build
```

#### 🚀 Produção (PostgreSQL - Heroku)
Configuração para deploy no Heroku.
```bash
# Usando Makefile
make run-production

# Ou manualmente
cp .env.production .env
go run main.go
```

### 4. Comandos Úteis

```bash
# Ver todos os comandos disponíveis
make help

# Resetar banco SQLite
make db-reset

# Resetar banco Docker PostgreSQL
make db-reset-docker

# Testar API
make test-api

# Executar testes
make test
```

## 📋 Endpoints Disponíveis

### Orçamentos
- `POST /api/v1/orcamentos` - Criar orçamento
- `GET /api/v1/orcamentos` - Listar todos os orçamentos
- `GET /api/v1/orcamentos/:id` - Buscar orçamento por ID
- `PUT /api/v1/orcamentos/:id` - Atualizar orçamento
- `DELETE /api/v1/orcamentos/:id` - Deletar orçamento

### Compatibilidade
- `POST /orcament` - Endpoint legado (compatibilidade)

### Utilitários
- `GET /health` - Health check
- `GET /api/v1/ping` - Ping/Pong

## 🧪 Exemplo de Uso

### Criar Orçamento
```bash
curl -X POST http://localhost:8080/api/v1/orcamentos \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "João Silva",
    "email": "joao@example.com",
    "telefone": "(11) 99999-9999",
    "quantidade_pecas": 10,
    "descricao": "Peças para projeto X"
  }'
```

### Listar Orçamentos
```bash
curl http://localhost:8080/api/v1/orcamentos
```

## 🔧 Configuração de Banco de Dados

O projeto suporta múltiplos ambientes com diferentes bancos de dados:

### 🗃️ SQLite (Desenvolvimento Local)
```env
# .env.local
DB_TYPE=sqlite
DB_PATH=./data/turning_back.db
```

### 🐘 PostgreSQL (Docker/Produção)
```env
# .env.docker ou .env.production
DB_TYPE=postgres
DB_HOST=localhost
DB_USER=turning_back_user
DB_PASSWORD=turning_back_pass
DB_NAME=turning_back_db
DB_PORT=5432
DB_SSLMODE=disable
DB_TIMEZONE=UTC
```

### 🔄 Detecção Automática
O sistema detecta automaticamente o tipo de banco baseado em:
1. Variável `DB_TYPE` (sqlite/postgres)
2. Variável `ENV` (development/production)
3. Presença de configurações PostgreSQL

## 🏆 Vantagens da Arquitetura Hexagonal

1. **Testabilidade**: Fácil criação de mocks e testes unitários
2. **Flexibilidade**: Troca de adaptadores sem afetar o core
3. **Independência**: Core business isolado de frameworks
4. **Manutenibilidade**: Código organizado e bem estruturado
5. **Escalabilidade**: Fácil adição de novos adaptadores


## 📚 Referências

- [Hexagonal Architecture by Alistair Cockburn](https://alistair.cockburn.us/hexagonal-architecture/)
- [Ports & Adapters Pattern](https://jmgarridopaz.github.io/content/hexagonalarchitecture.html)
- [Go Clean Architecture Examples](https://github.com/bxcodec/go-clean-arch)

---

**Nota**: Esta implementação mantém compatibilidade com o endpoint legado `/orcament` enquanto oferece uma API REST completa em `/api/v1/orcamentos`.