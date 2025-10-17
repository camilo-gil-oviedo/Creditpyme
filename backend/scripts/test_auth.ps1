# Prueba automática: register -> exchange custom token -> call /api/me

param(
    [string]$ApiKey = "REPLACE_WITH_YOUR_API_KEY",
    [string]$BaseUrl = "http://localhost:8080"
)

Write-Host "Registrando usuario de prueba..."
$body = @{ email = "test@example.com"; password = "password123" } | ConvertTo-Json
$resp = Invoke-RestMethod -Method Post -Uri "$BaseUrl/register" -Body $body -ContentType 'application/json'

if (-not $resp.token) {
    Write-Error "No se recibió token del endpoint /register"
    exit 1
}

$custom = $resp.token
Write-Host "Custom token recibido: $custom"

Write-Host "Intercambiando custom token por idToken (API_KEY must be set)..."
if ($ApiKey -eq "REPLACE_WITH_YOUR_API_KEY") {
    Write-Error "Por favor coloca tu API key en el parámetro ApiKey o edita el script."
    exit 1
}

$body2 = @{ token = $custom; returnSecureToken = $true } | ConvertTo-Json
$res = Invoke-RestMethod -Method Post -Uri "https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=$ApiKey" -Body $body2 -ContentType 'application/json'

$idToken = $res.idToken
if (-not $idToken) {
    Write-Error "No se recibió idToken al intercambiar custom token"
    exit 1
}

Write-Host "idToken recibido: $idToken"

Write-Host "Llamando a /api/me con el idToken..."
$headers = @{ Authorization = "Bearer $idToken" }
$result = Invoke-RestMethod -Method Get -Uri "$BaseUrl/api/me" -Headers $headers

Write-Host "Resultado de /api/me:"
$result | ConvertTo-Json -Depth 5 | Write-Host
