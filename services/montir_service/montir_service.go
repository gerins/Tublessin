package main

import (
	"net"
	"tublessin/common/model"
	"tublessin/services/montir_service/config"
	"tublessin/services/montir_service/config/mysql"
	"tublessin/services/montir_service/config/redis"
	"tublessin/services/montir_service/domain"
	"tublessin/services/montir_service/utils/logging"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

func main() {
	logging.LoggingToFile()
	config.SetEnvironmentVariables()

	srv := grpc.NewServer()
	montirServer := domain.NewMontirController(mysql.ConnectToDatabase(), redis.NewRedisConnection())
	model.RegisterMontirServer(srv, montirServer)

	log.Println("Starting Montir-Service server at port", config.GRPC_SERVICE_MONTIR_PORT)
	l, err := net.Listen(config.GRPC_SERVICE_MONTIR_HOST, ":"+config.GRPC_SERVICE_MONTIR_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.GRPC_SERVICE_MONTIR_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}
