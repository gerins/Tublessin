package main

import (
	"database/sql"
	"fmt"
	"net"
	"tublessin/common/model"
	"tublessin/services/transaction_service/config"
	"tublessin/services/transaction_service/domain"

	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

func main() {
	// logging.LoggingToFile()
	config.SetEnvironmentVariables()
	srv := grpc.NewServer()
	transactionServer := domain.NewTransactionController(connectToDatabase())
	model.RegisterTransactionServer(srv, transactionServer)

	log.Println("Starting Transaction-Service server at port", config.GRPC_SERVICE_TRANSACTION_PORT)
	l, err := net.Listen(config.GRPC_SERVICE_TRANSACTION_HOST, ":"+config.GRPC_SERVICE_TRANSACTION_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.GRPC_SERVICE_TRANSACTION_PORT, err)
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
