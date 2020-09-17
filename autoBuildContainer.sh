export GOOS=linux  
export GOARCH=amd64 

read -p "Press enter to continue"
cd api_gateway
go build -o tublessin-gateway
docker build -t gerins/tublessin-gateway .
docker push gerins/tublessin-gateway &
rm tublessin-gateway

cd ..
cd database
cd mysql
docker build -t gerins/tublessin-db .
docker push gerins/tublessin-db &

cd ..
cd redis
docker build -t gerins/tublessin-redis .
docker push gerins/tublessin-redis &

cd ..
cd ..

cd services
cd login_service
go build -o tublessin-login
docker build -t gerins/tublessin-login .
docker push gerins/tublessin-login &
rm tublessin-login

cd ..
cd montir_service
go build -o tublessin-montir
docker build -t gerins/tublessin-montir .
docker push gerins/tublessin-montir &
rm tublessin-montir

cd ..
cd user_service
go build -o tublessin-user
docker build -t gerins/tublessin-user .
docker push gerins/tublessin-user &
rm tublessin-user

cd ..
cd transaction_service
go build -o tublessin-transaction
docker build -t gerins/tublessin-transaction .
docker push gerins/tublessin-transaction &
rm tublessin-transaction

cd ..
cd chat_service
go build -o tublessin-chat
docker build -t gerins/tublessin-chat .
docker push gerins/tublessin-chat &
rm tublessin-chat

read -p "Updating container complete, press enter to continue"