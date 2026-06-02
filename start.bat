@echo off
setlocal

set "ROOT_DIR=%~dp0"
pushd "%ROOT_DIR%" >nul

where go >nul 2>nul
if errorlevel 1 (
	echo [ERROR] Go is not installed or not available on PATH.
	popd >nul
	exit /b 1
)

if not exist go.mod (
	echo [INFO] Initializing Go module...
	go mod init goflare
	if errorlevel 1 goto :fail
)

echo [INFO] Installing and resolving Go dependencies...
go mod tidy
if errorlevel 1 goto :fail

echo [INFO] Starting development server on http://localhost:3000 ...
go run ./server/main.go
if errorlevel 1 goto :fail

popd >nul
exit /b 0

:fail
set "EXIT_CODE=%errorlevel%"
echo [ERROR] Startup failed with exit code %EXIT_CODE%.
popd >nul
exit /b %EXIT_CODE%
