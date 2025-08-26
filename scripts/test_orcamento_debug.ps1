# Script para testar o endpoint de orçamento com debug
# Execute este script para testar diferentes cenários

Write-Host "=== Testando endpoint de orçamento com debug ===" -ForegroundColor Green

# URL do seu backend (ajuste se necessário)
$baseUrl = "http://localhost:8080"

Write-Host "`n1. Testando com JSON válido:" -ForegroundColor Yellow

$validJson = @{
    nome = "João Silva"
    email = "joao@email.com"
    telefone = "(11) 99999-9999"
    quantidade_pecas = 5
    descricao = "Peças para projeto X"
    anexo = "arquivo.pdf"
} | ConvertTo-Json

Write-Host "JSON enviado:" -ForegroundColor Cyan
Write-Host $validJson

try {
    $response = Invoke-RestMethod -Uri "$baseUrl/orcament" -Method POST -Body $validJson -ContentType "application/json"
    Write-Host "Resposta:" -ForegroundColor Green
    $response | ConvertTo-Json -Depth 3
} catch {
    Write-Host "Erro:" -ForegroundColor Red
    Write-Host $_.Exception.Message
    if ($_.Exception.Response) {
        $reader = New-Object System.IO.StreamReader($_.Exception.Response.GetResponseStream())
        $responseBody = $reader.ReadToEnd()
        Write-Host "Response Body:" -ForegroundColor Red
        Write-Host $responseBody
    }
}

Write-Host "`n2. Testando com JSON inválido (sem campo obrigatório):" -ForegroundColor Yellow

$invalidJson = @{
    nome = "João Silva"
    # email ausente (obrigatório)
    telefone = "(11) 99999-9999"
    quantidade_pecas = 5
} | ConvertTo-Json

Write-Host "JSON enviado:" -ForegroundColor Cyan
Write-Host $invalidJson

try {
    $response = Invoke-RestMethod -Uri "$baseUrl/orcament" -Method POST -Body $invalidJson -ContentType "application/json"
    Write-Host "Resposta:" -ForegroundColor Green
    $response | ConvertTo-Json -Depth 3
} catch {
    Write-Host "Erro esperado:" -ForegroundColor Red
    Write-Host $_.Exception.Message
    if ($_.Exception.Response) {
        $reader = New-Object System.IO.StreamReader($_.Exception.Response.GetResponseStream())
        $responseBody = $reader.ReadToEnd()
        Write-Host "Response Body:" -ForegroundColor Red
        Write-Host $responseBody
    }
}

Write-Host "`n3. Testando com JSON malformado:" -ForegroundColor Yellow

$malformedJson = '{"nome": "João", "email": "joao@email.com", "telefone": "(11) 99999-9999", "quantidade_pecas": 5'  # JSON incompleto

Write-Host "JSON enviado:" -ForegroundColor Cyan
Write-Host $malformedJson

try {
    $response = Invoke-RestMethod -Uri "$baseUrl/orcament" -Method POST -Body $malformedJson -ContentType "application/json"
    Write-Host "Resposta:" -ForegroundColor Green
    $response | ConvertTo-Json -Depth 3
} catch {
    Write-Host "Erro esperado:" -ForegroundColor Red
    Write-Host $_.Exception.Message
    if ($_.Exception.Response) {
        $reader = New-Object System.IO.StreamReader($_.Exception.Response.GetResponseStream())
        $responseBody = $reader.ReadToEnd()
        Write-Host "Response Body:" -ForegroundColor Red
        Write-Host $responseBody
    }
}

Write-Host "`n=== Teste concluído ===" -ForegroundColor Green
Write-Host "Verifique os logs do seu backend para ver os detalhes de debug!" -ForegroundColor Cyan