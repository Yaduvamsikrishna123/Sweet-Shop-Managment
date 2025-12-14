Write-Host "Starting Sweet Shop Management System..." -ForegroundColor Green

$projectRoot = Get-Location

# Start Backend
Write-Host "Starting Backend Server..." -ForegroundColor Cyan
Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd '$projectRoot/Backend'; go run cmd/server/main.go"

# Start Frontend
Write-Host "Starting Frontend Server..." -ForegroundColor Cyan
Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd '$projectRoot/frontend'; npm run dev"

Write-Host "Both servers are starting in separate windows." -ForegroundColor Green
Write-Host "Backend URL: http://localhost:8080"
Write-Host "Frontend URL: http://localhost:5173"
