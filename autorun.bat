cd api_gateway
go build api_gateway.go
START api_gateway.exe

cd ..
cd services
cd login_service
go build login_service.go
START login_service.exe

cd ..
cd montir_service
go build montir_service.go
START montir_service.exe

cd ..
cd user_service
go build user_service.go
START user_service.exe
