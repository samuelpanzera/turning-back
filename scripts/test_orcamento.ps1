# Script para testar o endpoint de orçamento
# Execute este script após iniciar o servidor com: go run cmd/api/main.go

$baseUrl = "http://localhost:8080"

# Teste 1: Criar um orçamento válido
Write-Host "Testando criação de orçamento..." -ForegroundColor Green

$orcamentoData = @{
    quantidade_pecas = 15
    descricao = "Peças para linha de produção"
    nome = "Carlos Oliveira"
    anexo = "https://example.com/especificacoes.pdf"
    email = "carlos@empresa.com"
    telefone = "(11) 98765-4321"
} | ConvertTo-Json

try {
    $response = Invoke-RestMethod -Uri "$baseUrl/orcament" -Method POST -Body $orcamentoData -ContentType "application/json"
    Write-Host "✅ Orçamento criado com sucesso!" -ForegroundColor Green
    Write-Host "ID: $($response.orcamento.id)" -ForegroundColor Yellow
    $orcamentoId = $response.orcamento.id
} catch {
    Write-Host "❌ Erro ao criar orçamento: $($_.Exception.Message)" -ForegroundColor Red
    exit 1
}

# Teste 2: Buscar o orçamento criado
Write-Host "`nBuscando orçamento por ID..." -ForegroundColor Green

try {
    $response = Invoke-RestMethod -Uri "$baseUrl/orcament/$orcamentoId" -Method GET
    Write-Host "✅ Orçamento encontrado!" -ForegroundColor Green
    Write-Host "Nome: $($response.nome)" -ForegroundColor Yellow
    Write-Host "Email: $($response.email)" -ForegroundColor Yellow
} catch {
    Write-Host "❌ Erro ao buscar orçamento: $($_.Exception.Message)" -ForegroundColor Red
}

# Teste 3: Listar todos os orçamentos
Write-Host "`nListando todos os orçamentos..." -ForegroundColor Green

try {
    $response = Invoke-RestMethod -Uri "$baseUrl/orcament" -Method GET
    Write-Host "✅ Lista obtida com sucesso!" -ForegroundColor Green
    Write-Host "Total de orçamentos: $($response.total)" -ForegroundColor Yellow
} catch {
    Write-Host "❌ Erro ao listar orçamentos: $($_.Exception.Message)" -ForegroundColor Red
}

# Teste 4: Tentar criar orçamento inválido (sem email)
Write-Host "`nTestando validação (orçamento sem email)..." -ForegroundColor Green

$orcamentoInvalido = @{
    quantidade_pecas = 5
    nome = "Teste Inválido"
    telefone = "(11) 99999-9999"
} | ConvertTo-Json

try {
    $response = Invoke-RestMethod -Uri "$baseUrl/orcament" -Method POST -Body $orcamentoInvalido -ContentType "application/json"
    Write-Host "❌ Deveria ter falhado!" -ForegroundColor Red
} catch {
    Write-Host "✅ Validação funcionando corretamente!" -ForegroundColor Green
    Write-Host "Erro esperado: $($_.Exception.Message)" -ForegroundColor Yellow
}

Write-Host "`n🎉 Testes concluídos!" -ForegroundColor Cyan