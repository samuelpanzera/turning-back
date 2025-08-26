-- Migração para tornar o campo email opcional
-- Execute este script no seu banco de dados SQLite

-- Para SQLite, precisamos recriar a tabela pois ALTER COLUMN não é suportado
-- Primeiro, criar uma tabela temporária
CREATE TABLE orcamentos_temp (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    quantidade_pecas INTEGER NOT NULL,
    descricao TEXT,
    nome TEXT NOT NULL,
    anexo TEXT,
    email TEXT,  -- Removido NOT NULL
    telefone TEXT NOT NULL
);

-- Copiar dados existentes (com email padrão para registros sem email)
INSERT INTO orcamentos_temp (
    id, created_at, updated_at, deleted_at, 
    quantidade_pecas, descricao, nome, anexo, email, telefone
)
SELECT 
    id, created_at, updated_at, deleted_at,
    quantidade_pecas, descricao, nome, anexo,
    CASE 
        WHEN email = 'nao-informado@temp.com' THEN NULL 
        ELSE email 
    END as email,
    telefone
FROM orcamentos;

-- Remover tabela original
DROP TABLE orcamentos;

-- Renomear tabela temporária
ALTER TABLE orcamentos_temp RENAME TO orcamentos;

-- Recriar índices se necessário
CREATE INDEX idx_orcamentos_deleted_at ON orcamentos(deleted_at);