package main

import (
	"net"
	"tublessin/common/model"
	"tublessin/services/user_service/config"
	"tublessin/services/user_service/config/mysql"
	"tublessin/services/user_service/config/redis"
	"tublessin/services/user_service/domain"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

func main() {
	// logging.LoggingToFile()
	config.SetEnvironmentVariables()
	srv := grpc.NewServer()
	userServer := domain.NewUserController(mysql.ConnectToDatabase(), redis.NewRedisConnection())
	model.RegisterUserServer(srv, userServer)

	log.Println("Starting User-Service server at port", config.GRPC_SERVICE_USER_PORT)
	l, err := net.Listen(config.GRPC_SERVICE_USER_HOST, ":"+config.GRPC_SERVICE_USER_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.GRPC_SERVICE_USER_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}
