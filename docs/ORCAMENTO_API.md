# API de Orçamento

## Endpoint POST /orcament

Este endpoint permite criar uma solicitação de orçamento.

### Campos Obrigatórios
- `quantidade_pecas` (int): Quantidade de peças (mínimo 1)
- `nome` (string): Nome do solicitante
- `email` (string): Email válido do solicitante
- `telefone` (string): Telefone do solicitante

### Campos Opcionais
- `descricao` (string): Descrição adicional do orçamento
- `anexo` (string): URL ou caminho para anexo

### Exemplo de Requisição

```bash
curl -X POST http://localhost:8080/orcament \
  -H "Content-Type: application/json" \
  -d '{
    "quantidade_pecas": 10,
    "descricao": "Peças para projeto de automação industrial",
    "nome": "João Silva",
    "anexo": "https://example.com/projeto.pdf",
    "email": "joao.silva@empresa.com",
    "telefone": "(11) 99999-9999"
  }'
```

### Exemplo de Resposta (201 Created)

```json
{
  "message": "Orçamento criado com sucesso",
  "orcamento": {
    "id": 1,
    "created_at": "2025-01-27T10:30:00Z",
    "updated_at": "2025-01-27T10:30:00Z",
    "quantidade_pecas": 10,
    "descricao": "Peças para projeto de automação industrial",
    "nome": "João Silva",
    "anexo": "https://example.com/projeto.pdf",
    "email": "joao.silva@empresa.com",
    "telefone": "(11) 99999-9999"
  }
}
```

### Exemplo de Erro (400 Bad Request)

```json
{
  "error": "Invalid request data",
  "details": "Key: 'Orcamento.Email' Error:Tag: 'email' Error:Field validation for 'Email' failed on the 'email' tag"
}
```

## Outros Endpoints Disponíveis

### GET /orcament
Lista todos os orçamentos

### GET /orcament/:id
Busca um orçamento específico por ID

### Exemplo de uso com JavaScript

```javascript
const orcamentoData = {
  quantidade_pecas: 5,
  descricao: "Peças de reposição",
  nome: "Maria Santos",
  email: "maria@example.com",
  telefone: "(11) 88888-8888"
};

fetch('http://localhost:8080/orcament', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify(orcamentoData)
})
.then(response => response.json())
.then(data => console.log('Sucesso:', data))
.catch(error => console.error('Erro:', error));
```