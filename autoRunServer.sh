cd gateway_service 
export API_GATEWAY_SERVER_HOST=0.0.0.0
export API_GATEWAY_SERVER_PORT=8084
export SERVICE_LOGIN_HOST=localhost
export SERVICE_LOGIN_PORT=9000
export SERVICE_TRANSACTION_HOST=localhost
export SERVICE_TRANSACTION_PORT=6000
export SERVICE_MONTIR_HOST=localhost
export SERVICE_MONTIR_PORT=8000
export SERVICE_USER_HOST=localhost
export SERVICE_USER_PORT=7000
export SERVICE_CHAT_HOST=localhost
export SERVICE_CHAT_PORT=5000
nohup ./tublessin-gateway &

cd ..
cd services
cd login_service
export GRPC_SERVICE_LOGIN_HOST=tcp
export GRPC_SERVICE_LOGIN_PORT=9000
export SERVICE_MONTIR_HOST=localhost
export SERVICE_MONTIR_PORT=8000
export SERVICE_USER_HOST=localhost
export SERVICE_USER_PORT=7000
nohup ./tublessin-login &

cd ..
cd montir_service
export GRPC_SERVICE_MONTIR_HOST=tcp
export GRPC_SERVICE_MONTIR_PORT=8000
export REDIS_DATABASE_HOST=127.0.0.1
export REDIS_DATABASE_PORT=6379
export REDIS_DATABASE_USERNAME=admin
export REDIS_DATABASE_PASSWORD=
export REDIS_DATABASE_SELECT=1
export MYSQL_DB_DRIVER=mysql
export MYSQL_DB_USER=squad4
export MYSQL_DB_PASSWORD=password
export MYSQL_DB_NAME=squad4
export MYSQL_DB_HOST=34.101.198.49
export MYSQL_DB_PORT=3306
nohup ./tublessin-montir &

cd ..
cd user_service
export GRPC_SERVICE_USER_HOST=tcp
export GRPC_SERVICE_USER_PORT=7000
export REDIS_DATABASE_HOST=127.0.0.1
export REDIS_DATABASE_PORT=6379
export REDIS_DATABASE_USERNAME=admin
export REDIS_DATABASE_PASSWORD=
export REDIS_DATABASE_SELECT=2
export MYSQL_DB_DRIVER=mysql
export MYSQL_DB_USER=squad4
export MYSQL_DB_PASSWORD=password
export MYSQL_DB_NAME=squad4
export MYSQL_DB_HOST=34.101.198.49
export MYSQL_DB_PORT=3306
nohup ./tublessin-user &

cd ..
cd chat_service
export GRPC_SERVICE_CHAT_HOST=tcp
export GRPC_SERVICE_CHAT_PORT=5000
export REDIS_DATABASE_HOST=127.0.0.1
export REDIS_DATABASE_PORT=6379
export REDIS_DATABASE_USERNAME=admin
export REDIS_DATABASE_PASSWORD=
export REDIS_DATABASE_SELECT=3
export MYSQL_DB_DRIVER=mysql
export MYSQL_DB_USER=squad4
export MYSQL_DB_PASSWORD=password
export MYSQL_DB_NAME=squad4
export MYSQL_DB_HOST=34.101.198.49
export MYSQL_DB_PORT=3306
nohup ./tublessin-chat &

cd ..
cd transaction_service
export GRPC_SERVICE_TRANSACTION_HOST=tcp
export GRPC_SERVICE_TRANSACTION_PORT=6000
export MYSQL_DB_DRIVER=mysql
export MYSQL_DB_USER=squad4
export MYSQL_DB_PASSWORD=password
export MYSQL_DB_NAME=squad4
export MYSQL_DB_HOST=34.101.198.49
export MYSQL_DB_PORT=3306
nohup ./tublessin-transaction &