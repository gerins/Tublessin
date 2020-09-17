package main

import (
	"log"
	"net"
	"tublessin/common/model"
	"tublessin/services/chat_service/config"
	"tublessin/services/chat_service/config/mysql"
	"tublessin/services/chat_service/config/redis"
	"tublessin/services/chat_service/domain"

	"google.golang.org/grpc"
)

func main() {
	config.SetEnvironmentVariables()
	srv := grpc.NewServer()

	chatServer := domain.NewChatController(mysql.ConnectToDatabase(), redis.NewRedisConnection())
	model.RegisterChatServer(srv, chatServer)

	log.Println("Starting CHAT-Service server at port", config.GRPC_SERVICE_CHAT_PORT)
	l, err := net.Listen(config.GRPC_SERVICE_CHAT_HOST, ":"+config.GRPC_SERVICE_CHAT_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.GRPC_SERVICE_CHAT_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}
