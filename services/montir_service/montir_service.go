package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"tublessin/common/model"
	"tublessin/services/montir_service/config"
	"tublessin/services/montir_service/domain"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {
	config.SetEnvironmentVariables()
	srv := grpc.NewServer()
	montirServer := domain.NewMontirController(connectToDatabase())
	model.RegisterMontirServer(srv, montirServer)

	log.Println("Starting Montir-Service server at port", config.GRPC_SERVICE_MONTIR_PORT)
	l, err := net.Listen(config.GRPC_SERVICE_MONTIR_HOST, ":"+config.GRPC_SERVICE_MONTIR_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.GRPC_SERVICE_MONTIR_PORT, err)
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
