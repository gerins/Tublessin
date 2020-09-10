package main

import (
	"database/sql"
	"fmt"
	"net"
	"tublessin/common/model"
	"tublessin/services/user_service/config"
	"tublessin/services/user_service/domain"
	"tublessin/services/user_service/utils/logging"

	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {
	logging.LoggingToFile()
	config.SetEnvironmentVariables()
	srv := grpc.NewServer()
	userServer := domain.NewUserController(connectToDatabase())
	model.RegisterUserServer(srv, userServer)

	log.Println("Starting User-Service server at port", config.GRPC_SERVICE_USER_PORT)
	l, err := net.Listen(config.GRPC_SERVICE_USER_HOST, ":"+config.GRPC_SERVICE_USER_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.GRPC_SERVICE_USER_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}

func connectToDatabase() *sql.DB {
	db, err := sql.Open(config.DbDriver, config.DbUser+":"+config.DbPass+"@tcp("+config.DbHost+":"+config.DbPort+")/"+config.DbName)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Print(err)
		fmt.Scanln()
		log.Fatal(err)
	}
	log.Println("DataBase Successfully Connected")
	return db
}
