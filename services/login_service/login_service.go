package main

import (
	"net"
	"tublessin/common/model"
	"tublessin/services/login_service/config"
	"tublessin/services/login_service/domain"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

func main() {
	// logging.LoggingToFile()
	config.SetEnvironmentVariables()
	srv := grpc.NewServer()
	loginServer := domain.NewLoginController(connectToServiceMontir(), connectToServiceUser())
	model.RegisterLoginServer(srv, loginServer)

	log.Println("Starting Login-Service server at port", config.GRPC_SERVICE_LOGIN_PORT)
	l, err := net.Listen(config.GRPC_SERVICE_LOGIN_HOST, ":"+config.GRPC_SERVICE_LOGIN_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.GRPC_SERVICE_LOGIN_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}

func connectToServiceMontir() model.MontirClient {
	host := config.SERVICE_MONTIR_HOST
	port := config.SERVICE_MONTIR_PORT
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not Connect to Montir-Service", port, err)
	}

	return model.NewMontirClient(conn)
}

func connectToServiceUser() model.UserClient {
	host := config.SERVICE_USER_HOST
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not Connect to User-Service", port, err)
	}

	return model.NewUserClient(conn)
}
