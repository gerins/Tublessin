package config

import (
	"fmt"
	"log"
	"os"
)

var DbDriver,
	DbUser,
	DbPass,
	DbName,
	DbHost,
	DbPort,
	GRPC_SERVICE_USER_HOST,
	GRPC_SERVICE_USER_PORT string

// Config server Host and Port
func SetEnvironmentVariables() {
	GRPC_SERVICE_USER_HOST = os.Getenv("GRPC_SERVICE_USER_HOST")
	GRPC_SERVICE_USER_PORT = os.Getenv("GRPC_SERVICE_USER_PORT")

	DbDriver = os.Getenv("MYSQL_DB_DRIVER")
	DbUser = os.Getenv("MYSQL_DB_USER")
	DbPass = os.Getenv("MYSQL_DB_PASSWORD")
	DbName = os.Getenv("MYSQL_DB_NAME")
	DbHost = os.Getenv("MYSQL_DB_HOST")
	DbPort = os.Getenv("MYSQL_DB_PORT")

	log.Print(`ENVIRONMENT VARIABLE`)
	fmt.Println()

	log.Println(`GRPC_SERVICE_USER_HOST=`, GRPC_SERVICE_USER_HOST)
	log.Println(`GRPC_SERVICE_USER_PORT=`, GRPC_SERVICE_USER_PORT)
	fmt.Println()

	log.Println(`MYSQL_DB_DRIVER=`, DbDriver)
	log.Println(`MYSQL_DB_USER=`, DbUser)
	log.Println(`MYSQL_DB_PASSWORD=`, DbPass)
	log.Println(`MYSQL_DB_NAME=`, DbName)
	log.Println(`MYSQL_DB_HOST=`, DbHost)
	log.Println(`MYSQL_DB_PORT=`, DbPort)
	fmt.Println()

}
