# Deploy no Heroku - Turning Back

## 🎯 Duas Opções de Deploy

### Opção 1: Buildpack Go (Recomendado - Mais Simples)

O Heroku detecta automaticamente projetos Go e faz o build nativo.

#### Pré-requisitos
- Conta no Heroku
- Heroku CLI instalado
- Git configurado

#### Passo a Passo

1. **Instalar Heroku CLI**:
   - Windows: https://devcenter.heroku.com/articles/heroku-cli
   - macOS: `brew tap heroku/brew && brew install heroku`
   - Linux: `curl https://cli-assets.heroku.com/install.sh | sh`

2. **Login no Heroku**:
```bash
heroku login
```

3. **Criar aplicação**:
```bash
heroku create turning-back-api
# ou com nome customizado:
# heroku create seu-nome-app
```

4. **Adicionar PostgreSQL**:
```bash
heroku addons:create heroku-postgresql:mini
```

5. **Configurar variáveis de ambiente**:
```bash
heroku config:set ENV=production
heroku config:set JWT_SECRET=your-super-secure-jwt-secret-here-change-this
heroku config:set LOG_LEVEL=info
heroku config:set LOG_FORMAT=json
```

6. **Deploy**:
```bash
git add .
git commit -m "Deploy to Heroku"
git push heroku main
```

7. **Verificar**:
```bash
heroku open
# ou
curl https://seu-app.herokuapp.com/health
```

#### Vantagens do Buildpack:
- ✅ Configuração mínima
- ✅ Build automático otimizado
- ✅ Detecção automática de dependências
- ✅ Menos complexidade

### Opção 2: Container Registry (Docker)

Use se precisar de controle total sobre o ambiente.

#### Passo a Passo

1. **Configurar stack container**:
```bash
heroku create turning-back-api
heroku stack:set container -a turning-back-api
```

2. **Adicionar PostgreSQL**:
```bash
heroku addons:create heroku-postgresql:mini -a turning-back-api
```

3. **Login no container registry**:
```bash
heroku container:login
```

4. **Build e deploy**:
```bash
heroku container:push web -a turning-back-api
heroku container:release web -a turning-back-api
```

5. **Configurar variáveis**:
```bash
heroku config:set ENV=production -a turning-back-api
heroku config:set JWT_SECRET=your-super-secure-jwt-secret -a turning-back-api
```

#### Vantagens do Container:
- ✅ Ambiente idêntico ao desenvolvimento
- ✅ Controle total das dependências
- ✅ Configurações customizadas

## 🗄️ Configuração do Banco de Dados

### PostgreSQL Automático

O Heroku configura automaticamente:
- `DATABASE_URL` - URL completa de conexão PostgreSQL

**Exemplo de DATABASE_URL**:
```
postgres://user:password@host:5432/database
```

### Verificar Configuração do Banco

```bash
# Ver todas as variáveis
heroku config

# Ver apenas DATABASE_URL
heroku config:get DATABASE_URL

# Conectar ao banco via CLI
heroku pg:psql
```

## 🔧 Configurações Importantes

### Variáveis de Ambiente Essenciais

```bash
# Obrigatórias
heroku config:set ENV=production
heroku config:set JWT_SECRET=sua-chave-super-secreta-aqui

# Opcionais (com valores padrão)
heroku config:set LOG_LEVEL=info
heroku config:set LOG_FORMAT=json
heroku config:set PORT=8080  # Heroku define automaticamente
```

### Procfile (Buildpack)

O arquivo `Procfile` já está configurado:
```
web: ./bin/turning-back
```

### heroku.yml (Container)

O arquivo `heroku.yml` já está configurado:
```yaml
build:
  docker:
    web: Dockerfile
run:
  web: /app/main
```

## 🚀 Comandos Úteis

### Logs e Monitoramento

```bash
# Ver logs em tempo real
heroku logs --tail

# Ver logs específicos
heroku logs --source app

# Informações da aplicação
heroku ps
heroku releases
```

### Manutenção

```bash
# Reiniciar aplicação
heroku restart

# Escalar dynos
heroku ps:scale web=1

# Modo manutenção
heroku maintenance:on
heroku maintenance:off
```

### Banco de Dados

```bash
# Backup do banco
heroku pg:backups:capture
heroku pg:backups:download

# Reset do banco (CUIDADO!)
heroku pg:reset DATABASE_URL --confirm seu-app-name
```

## 🔒 Segurança

### Variáveis Sensíveis

**NUNCA** commite no Git:
- JWT_SECRET
- Senhas de banco
- API keys

Use sempre `heroku config:set`:
```bash
heroku config:set JWT_SECRET=$(openssl rand -base64 32)
```

### SSL/TLS

O Heroku fornece SSL automático para domínios `.herokuapp.com`.

## 📊 Monitoramento

### Métricas Básicas

```bash
# Ver métricas
heroku logs --tail | grep "method="

# Adicionar New Relic (opcional)
heroku addons:create newrelic:wayne
```

### Health Check

A aplicação já tem endpoint de health check:
- `GET /health` - Status da aplicação

## 🐛 Troubleshooting

### Problemas Comuns

1. **Build falha**:
```bash
# Ver logs detalhados
heroku logs --tail

# Verificar go.mod
go mod tidy
```

2. **Aplicação não inicia**:
```bash
# Verificar PORT
heroku config:get PORT

# Verificar Procfile
cat Procfile
```

3. **Erro de banco**:
```bash
# Verificar DATABASE_URL
heroku config:get DATABASE_URL

# Testar conexão
heroku pg:psql
```

### Logs Úteis

```bash
# Filtrar por erro
heroku logs --tail | grep ERROR

# Ver apenas aplicação
heroku logs --source app --tail
```

## 💰 Custos

### Plano Gratuito (Eco Dynos)
- 1000 horas/mês grátis
- Aplicação "dorme" após 30min inativa
- PostgreSQL: 10k linhas grátis

### Upgrade Recomendado
```bash
# Dyno básico ($7/mês)
heroku ps:type basic

# PostgreSQL mini ($5/mês)
heroku addons:upgrade heroku-postgresql:mini
```

## 🎯 Resumo: Container vs Buildpack

### Use **Buildpack** se:
- ✅ Quer simplicidade máxima
- ✅ Projeto Go padrão
- ✅ Não precisa de dependências especiais
- ✅ Primeiro deploy no Heroku

### Use **Container** se:
- ✅ Precisa de controle total do ambiente
- ✅ Dependências específicas do sistema
- ✅ Já usa Docker no desenvolvimento
- ✅ Configurações avançadas

**Recomendação**: Comece com **Buildpack**. É mais simples e atende 90% dos casos. Migre para Container apenas se precisar de funcionalidades específicas.