package config

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

var DbDriver,
	DbUser,
	DbPass,
	DbName,
	DbHost,
	DbPort,
	GRPC_SERVICE_TRANSACTION_HOST,
	GRPC_SERVICE_TRANSACTION_PORT string

// Config server Host and Port
func SetEnvironmentVariables() {
	GRPC_SERVICE_TRANSACTION_HOST = os.Getenv("GRPC_SERVICE_TRANSACTION_HOST")
	GRPC_SERVICE_TRANSACTION_PORT = os.Getenv("GRPC_SERVICE_TRANSACTION_PORT")

	DbDriver = os.Getenv("MYSQL_DB_DRIVER")
	DbUser = os.Getenv("MYSQL_DB_USER")
	DbPass = os.Getenv("MYSQL_DB_PASSWORD")
	DbName = os.Getenv("MYSQL_DB_NAME")
	DbHost = os.Getenv("MYSQL_DB_HOST")
	DbPort = os.Getenv("MYSQL_DB_PORT")

	log.Print(`ENVIRONMENT VARIABLE`)
	fmt.Println()

	log.Println(`GRPC_SERVICE_TRANSACTION_HOST=`, GRPC_SERVICE_TRANSACTION_HOST)
	log.Println(`GRPC_SERVICE_TRANSACTION_PORT=`, GRPC_SERVICE_TRANSACTION_PORT)
	fmt.Println()

	log.Println(`MYSQL_DB_DRIVER=`, DbDriver)
	log.Println(`MYSQL_DB_USER=`, DbUser)
	log.Println(`MYSQL_DB_PASSWORD=`, DbPass)
	log.Println(`MYSQL_DB_NAME=`, DbName)
	log.Println(`MYSQL_DB_HOST=`, DbHost)
	log.Println(`MYSQL_DB_PORT=`, DbPort)
	fmt.Println()

}
