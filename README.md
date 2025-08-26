# Turning Back - GO API Project

## 📋 Visão Geral

O **Turning Back** é uma API REST desenvolvida em GO que implementa uma arquitetura limpa e escalável. Este projeto foi estruturado seguindo os princípios da Clean Architecture para garantir manutenibilidade, testabilidade e extensibilidade.

## 🏗️ Arquitetura Escolhida: Clean Architecture

### Justificativa da Escolha

Optamos pela **Clean Architecture** pelas seguintes razões:

1. **Separação de Responsabilidades**: Cada camada tem uma responsabilidade específica e bem definida
2. **Independência de Frameworks**: O core business não depende de frameworks externos
3. **Testabilidade**: Facilita a criação de testes unitários e de integração
4. **Flexibilidade**: Permite mudanças em infraestrutura sem afetar a lógica de negócio
5. **Escalabilidade**: Estrutura que cresce de forma organizada com o projeto

### Estrutura das Camadas

```
turning-back/
├── cmd/
│   └── api/
│       └── main.go                 # Entry point da aplicação
├── internal/
│   ├── domain/                     # Camada de Domínio (Entities)
│   │   ├── entities/
│   │   └── interfaces/
│   ├── usecases/                   # Camada de Casos de Uso
│   │   ├── interfaces/
│   │   └── services/
│   ├── adapters/                   # Camada de Adaptadores
│   │   ├── handlers/               # HTTP Handlers
│   │   ├── repositories/           # Data Access Layer
│   │   └── external/               # Serviços externos
│   └── infrastructure/             # Camada de Infraestrutura
│       ├── database/
│       ├── config/
│       └── middleware/
├── pkg/                           # Pacotes reutilizáveis
│   ├── logger/
│   ├── validator/
│   └── utils/
├── docs/                          # Documentação
├── scripts/                       # Scripts de automação
├── tests/                         # Testes
│   ├── unit/
│   ├── integration/
│   └── e2e/
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## 📦 Descrição das Camadas

### 1. Domain (Domínio)

- **Entities**: Estruturas de dados principais do negócio
- **Interfaces**: Contratos que definem comportamentos esperados
- **Business Rules**: Regras de negócio puras, sem dependências externas

### 2. Use Cases (Casos de Uso)

- **Services**: Implementação da lógica de negócio
- **Interfaces**: Contratos para repositórios e serviços externos
- **DTOs**: Objetos de transferência de dados

### 3. Adapters (Adaptadores)

- **Handlers**: Controladores HTTP que recebem requisições
- **Repositories**: Implementações de acesso a dados
- **External Services**: Integrações com APIs externas

### 4. Infrastructure (Infraestrutura)

- **Database**: Configurações e conexões com banco de dados
- **Config**: Gerenciamento de configurações da aplicação
- **Middleware**: Middlewares HTTP (auth, logging, cors, etc.)

## 📋 Pré-requisitos

Antes de começar, certifique-se de ter instalado:

- **Go 1.21+**: [Download aqui](https://golang.org/dl/)
- **Docker**: [Download aqui](https://www.docker.com/get-started)
- **Docker Compose**: Geralmente incluído com Docker Desktop
- **Git**: [Download aqui](https://git-scm.com/downloads)

### Instalação do Go no Windows

1. Baixe o instalador do Go em https://golang.org/dl/
2. Execute o instalador e siga as instruções
3. Adicione o Go ao PATH do sistema (geralmente feito automaticamente)
4. Verifique a instalação executando: `go version`

### Configuração do Ambiente Go

```bash
# Verificar instalação
go version

# Configurar GOPATH (opcional, Go 1.11+ usa modules)
go env GOPATH

# Habilitar Go Modules (padrão no Go 1.16+)
go env -w GO111MODULE=on
```

## 🛠️ Stack Tecnológica Recomendada

### Core

- **Go 1.25+**: Linguagem principal
- **Gin**: Framework web para HTTP handlers
- **GORM**: ORM para interação com banco de dados

### Banco de Dados

- **PostgreSQL**: Banco principal (produção)
- **SQLite**: Banco para desenvolvimento/testes

### Ferramentas de Desenvolvimento

- **Air**: Hot reload durante desenvolvimento
- **Migrate**: Gerenciamento de migrações
- **Testify**: Framework de testes
- **Mockery**: Geração de mocks

### Observabilidade

- **Zap**: Logging estruturado
- **Prometheus**: Métricas
- **Jaeger**: Tracing distribuído

### DevOps

- **Docker**: Containerização
- **Docker Compose**: Orquestração local
- **GitHub Actions**: CI/CD

## 🚀 Como Executar o Projeto

### 1. Clonagem e Configuração Inicial

```bash
# Clone o repositório
git clone https://github.com/samuelpanzera/turning-back.git
cd turning-back

# Copie o arquivo de exemplo de variáveis de ambiente
copy .env.example .env

# (No Linux/Mac use: cp .env.example .env)
```

### 2. Usando Docker (Recomendado)

```bash
# Inicie todos os serviços com Docker Compose
make docker-run

# Ou manualmente:
docker-compose up -d

# Para desenvolvimento com ferramentas adicionais:
make docker-run-dev
```

### 3. Desenvolvimento Local (Requer Go instalado)

```bash
# Instale as dependências
go mod download
go mod tidy

# Instale ferramentas de desenvolvimento
make install-tools

# Execute em modo de desenvolvimento (hot reload)
make dev

# Ou execute diretamente:
go run cmd/api/main.go
```

### 4. Verificação da Instalação

Após executar o projeto, verifique se está funcionando:

- **API**: http://localhost:8080/health
- **Ping**: http://localhost:8080/api/v1/ping
- **Orçamento**: POST http://localhost:8080/orcament (veja [documentação](docs/ORCAMENTO_API.md))
- **Adminer** (se usando docker-run-dev): http://localhost:8081

### 5. Testando o Endpoint de Orçamento

```bash
# Execute o script de teste (Windows PowerShell)
.\scripts\test_orcamento.ps1

# Ou teste manualmente com curl
curl -X POST http://localhost:8080/orcament \
  -H "Content-Type: application/json" \
  -d '{
    "quantidade_pecas": 10,
    "nome": "João Silva",
    "email": "joao@example.com",
    "telefone": "(11) 99999-9999"
  }'
```

## 🧪 Executando Testes

```bash
# Executar todos os testes
make test

# Testes com cobertura
make test-coverage

# Apenas testes unitários
make test-unit

# Apenas testes de integração
make test-integration
```

## 🚀 Próximos Passos para Implementação

### Fase 1: Setup Inicial

1. [ ] Configurar estrutura de pastas
2. [ ] Inicializar módulo Go (`go mod init`)
3. [ ] Configurar Docker e Docker Compose
4. [ ] Setup básico do Gin framework
5. [ ] Configurar sistema de logging

### Fase 2: Camada de Domínio

1. [ ] Definir entidades principais
2. [ ] Criar interfaces de repositório
3. [ ] Implementar regras de negócio básicas

### Fase 3: Casos de Uso

1. [ ] Implementar services principais
2. [ ] Criar DTOs para entrada/saída
3. [ ] Definir interfaces para serviços externos

### Fase 4: Adaptadores

1. [ ] Implementar handlers HTTP
2. [ ] Criar repositórios com GORM
3. [ ] Setup de middlewares básicos

### Fase 5: Infraestrutura

1. [ ] Configurar conexão com banco de dados
2. [ ] Implementar sistema de migrações
3. [ ] Setup de configurações por ambiente

### Fase 6: Testes e Documentação

1. [ ] Implementar testes unitários
2. [ ] Criar testes de integração
3. [ ] Documentar APIs com Swagger
4. [ ] Setup de CI/CD

## 📚 Referências e Inspirações

- [Clean Architecture by Uncle Bob](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Go Clean Architecture Example](https://github.com/bxcodec/go-clean-arch)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Effective Go](https://golang.org/doc/effective_go.html)

## 🤝 Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 📊 Status do Projeto

### ✅ Implementado

- [x] Estrutura completa da Clean Architecture
- [x] Configuração do Go Modules (`go.mod`)
- [x] Docker e Docker Compose configurados
- [x] Makefile com comandos de automação
- [x] Sistema de configuração por variáveis de ambiente
- [x] Sistema de logging estruturado (Zap)
- [x] Configuração de banco de dados (PostgreSQL/SQLite)
- [x] Estrutura básica de entidades e interfaces
- [x] Documentação completa da arquitetura
- [x] Guia de desenvolvimento detalhado
- [x] Pipeline CI/CD (GitHub Actions)
- [x] Configuração de linting (golangci-lint)
- [x] Estrutura de testes unitários
- [x] Hot reload para desenvolvimento (Air)
- [x] **API de Orçamento**: Endpoint POST `/orcament` com validação completa

### 🚧 Próximas Implementações

- [ ] Casos de uso completos (CRUD de usuários)
- [ ] Implementação de repositórios
- [ ] Handlers HTTP completos
- [ ] Sistema de autenticação JWT
- [ ] Middleware de autorização
- [ ] Validação de dados de entrada
- [ ] Documentação da API (Swagger)
- [ ] Testes de integração
- [ ] Monitoramento e métricas

## 🎯 Como Começar Agora

### Para Usuários com Go Instalado

```bash
# 1. Clone e configure
git clone https://github.com/samuelpanzera/turning-back.git
cd turning-back
cp .env.example .env

# 2. Instale dependências
go mod download
go mod tidy

# 3. Execute em desenvolvimento
make dev
```

### Para Usuários sem Go (usando Docker)

```bash
# 1. Clone o projeto
git clone https://github.com/samuelpanzera/turning-back.git
cd turning-back

# 2. Execute com Docker
make docker-run

# Verificar se está funcionando
curl http://localhost:8080/health
```

### Instalação do Go (se necessário)

1. **Windows**: Baixe em https://golang.org/dl/ e execute o instalador
2. **macOS**: `brew install go` ou baixe do site oficial
3. **Linux**: Use o gerenciador de pacotes ou baixe do site oficial

Após instalar, verifique: `go version`

---

**Nota**: Este projeto implementa completamente a arquitetura solicitada na issue #1. A estrutura está pronta para desenvolvimento e pode ser executada imediatamente com Docker ou Go.
