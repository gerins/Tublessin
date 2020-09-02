@ECHO OFF
TITLE Tublessin Microservice

ECHO =============================
ECHO    Tublessin Microservice 
ECHO =============================
ECHO 1. Menjalankan Server Api Gateway
cd api_gateway
go build api_gateway.go
START /min api_gateway.exe

ECHO 2. Menjalankan Server Login Service
cd ..
cd services
cd login_service
go build login_service.go
START /min login_service.exe

ECHO 3. Menjalankan Server Montir Service
cd ..
cd montir_service
go build montir_service.go
START /min montir_service.exe

ECHO 4. Menjalankan Server User Service
cd ..
cd user_service
go build user_service.go
START /min user_service.exe

ECHO 5. Hacking NASA Server
ECHO 6. Hacking FBI Server
echo.
ECHO Press any key for terminate all server...
PAUSE >nul
taskkill /im api_gateway.exe /f
taskkill /im login_service.exe /f
taskkill /im montir_service.exe /f
taskkill /im user_service.exe /f
echo.
ECHO Press any key for Restarting all server...
PAUSE >nul
cd ..
cd ..
START autoRunServer.bat