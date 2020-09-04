package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"tublessin/common/config"
	"tublessin/common/model"
	"tublessin/services/montir_service/domain"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

// Ini gaboleh di hardcode, nanti diganti jadi environment variabel kalo semua udah siap
const (
	dbDriver = "mysql"
	dbUser   = "root"
	dbPass   = "admin"
	dbName   = "tublessin_montir"
	dbHost   = "localhost"
	dbPort   = "3306"
)

func main() {
	srv := grpc.NewServer()
	montirServer := domain.NewMontirController(connectToDatabase())
	model.RegisterMontirServer(srv, montirServer)

	result, _ := montirServer.GetAllActiveMontirWithLocation(context.Background(), &model.RequestActiveMontir{})
	fmt.Println(result)

	log.Println("Starting Montir-Service server at port", config.SERVICE_MONTIR_PORT)
	l, err := net.Listen("tcp", config.SERVICE_MONTIR_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_MONTIR_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}

func connectToDatabase() *sql.DB {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
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
