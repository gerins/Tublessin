@ECHO OFF
TITLE Tublessin Microservice

ECHO =============================
ECHO    Tublessin Microservice 
ECHO =============================

ECHO 1. Menjalankan Server Api Gateway
cd api_gateway
go build api_gateway.go
SET API_GATEWAY_SERVER_HOST=localhost
SET API_GATEWAY_SERVER_PORT=8080
SET SERVICE_LOGIN_HOST=localhost
SET SERVICE_LOGIN_PORT=9000
SET SERVICE_TRANSACTION_HOST=localhost
SET SERVICE_TRANSACTION_PORT=6000
SET SERVICE_MONTIR_HOST=localhost
SET SERVICE_MONTIR_PORT=8000
SET SERVICE_USER_HOST=localhost
SET SERVICE_USER_PORT=7000
SET SERVICE_CHAT_HOST=localhost
SET SERVICE_CHAT_PORT=5000
START /min api_gateway.exe


ECHO 2. Menjalankan Server Login Service
cd ..
cd services
cd login_service
go build login_service.go
SET GRPC_SERVICE_LOGIN_HOST=tcp
SET GRPC_SERVICE_LOGIN_PORT=9000
SET SERVICE_MONTIR_HOST=localhost
SET SERVICE_MONTIR_PORT=8000
SET SERVICE_USER_HOST=localhost
SET SERVICE_USER_PORT=7000
START /min login_service.exe


ECHO 3. Menjalankan Server Montir Service
cd ..
cd montir_service
go build montir_service.go
SET GRPC_SERVICE_MONTIR_HOST=tcp
SET GRPC_SERVICE_MONTIR_PORT=8000
SET REDIS_DATABASE_HOST=127.0.0.1
SET REDIS_DATABASE_PORT=6379
SET REDIS_DATABASE_USERNAME=admin
SET REDIS_DATABASE_PASSWORD=redisadmin
SET REDIS_DATABASE_SELECT=1
SET MYSQL_DB_DRIVER=mysql
SET MYSQL_DB_USER=root
SET MYSQL_DB_PASSWORD=admin
SET MYSQL_DB_NAME=tublessin_montir
SET MYSQL_DB_HOST=localhost
SET MYSQL_DB_PORT=3306
START /min montir_service.exe


ECHO 4. Menjalankan Server User Service
cd ..
cd user_service
go build user_service.go
SET GRPC_SERVICE_USER_HOST=tcp
SET GRPC_SERVICE_USER_PORT=7000
SET REDIS_DATABASE_HOST=127.0.0.1
SET REDIS_DATABASE_PORT=6379
SET REDIS_DATABASE_USERNAME=admin
SET REDIS_DATABASE_PASSWORD=redisadmin
SET REDIS_DATABASE_SELECT=2
SET MYSQL_DB_DRIVER=mysql
SET MYSQL_DB_USER=root
SET MYSQL_DB_PASSWORD=admin
SET MYSQL_DB_NAME=tublessin_user
SET MYSQL_DB_HOST=localhost
SET MYSQL_DB_PORT=3306
START /min user_service.exe

ECHO 5. Menjalankan Server Chat Service
cd ..
cd chat_service
go build chat_service.go
SET GRPC_SERVICE_CHAT_HOST=tcp
SET GRPC_SERVICE_CHAT_PORT=5000
SET REDIS_DATABASE_HOST=127.0.0.1
SET REDIS_DATABASE_PORT=6379
SET REDIS_DATABASE_USERNAME=admin
SET REDIS_DATABASE_PASSWORD=redisadmin
SET REDIS_DATABASE_SELECT=3
SET MYSQL_DB_DRIVER=mysql
SET MYSQL_DB_USER=root
SET MYSQL_DB_PASSWORD=admin
SET MYSQL_DB_NAME=tublessin_chat
SET MYSQL_DB_HOST=localhost
SET MYSQL_DB_PORT=3306
START /min chat_service.exe


ECHO 6. Menjalankan Server Transaction Service
cd ..
cd transaction_service
go build transaction_service.go
SET GRPC_SERVICE_TRANSACTION_HOST=tcp
SET GRPC_SERVICE_TRANSACTION_PORT=6000
SET MYSQL_DB_DRIVER=mysql
SET MYSQL_DB_USER=root
SET MYSQL_DB_PASSWORD=admin
SET MYSQL_DB_NAME=tublessin_transaction
SET MYSQL_DB_HOST=localhost
SET MYSQL_DB_PORT=3306
START /min transaction_service.exe


echo.
ECHO Press any key for terminate all server...
PAUSE >nul
taskkill /im api_gateway.exe /f
taskkill /im login_service.exe /f
taskkill /im montir_service.exe /f
taskkill /im user_service.exe /f
taskkill /im transaction_service.exe /f
taskkill /im chat_service.exe /f
del transaction_service.exe /q
cd ..
cd chat_service
del chat_service.exe /q
cd ..
cd user_service
del user_service.exe /q
cd ..
cd montir_service
del montir_service.exe /q
cd ..
cd login_service
del login_service.exe /q
cd ..
cd ..
cd api_gateway
del api_gateway.exe /q
echo.
ECHO Press any key for Restarting all server...
PAUSE >nul
cd ..
START autoRunServer.bat
EXIT