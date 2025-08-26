# Exemplos de teste com curl

## 1. Teste com JSON válido

```bash
curl -X POST http://localhost:8080/orcament \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "João Silva",
    "email": "joao@email.com", 
    "telefone": "(11) 99999-9999",
    "quantidade_pecas": 5,
    "descricao": "Peças para projeto X",
    "anexo": "arquivo.pdf"
  }'
```

## 2. Teste com JSON inválido (campo obrigatório ausente)

```bash
curl -X POST http://localhost:8080/orcament \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "João Silva",
    "telefone": "(11) 99999-9999",
    "quantidade_pecas": 5
  }'
```

## 3. Teste com JSON malformado

```bash
curl -X POST http://localhost:8080/orcament \
  -H "Content-Type: application/json" \
  -d '{"nome": "João", "email": "joao@email.com"'
```

## 4. Teste sem Content-Type

```bash
curl -X POST http://localhost:8080/orcament \
  -d '{
    "nome": "João Silva",
    "email": "joao@email.com",
    "telefone": "(11) 99999-9999", 
    "quantidade_pecas": 5
  }'
```

## 5. Teste com verbose para ver headers

```bash
curl -v -X POST http://localhost:8080/orcament \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "João Silva",
    "email": "joao@email.com",
    "telefone": "(11) 99999-9999",
    "quantidade_pecas": 5
  }'
```

## Campos obrigatórios da entidade Orcamento:

- `nome` (string, obrigatório)
- `email` (string, obrigatório, deve ser email válido)  
- `telefone` (string, obrigatório)
- `quantidade_pecas` (int, obrigatório, mínimo 1)
- `descricao` (string, opcional)
- `anexo` (string, opcional)