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

	REDIS_DATABASE_HOST,
	REDIS_DATABASE_PORT,
	REDIS_DATABASE_USERNAME,
	REDIS_DATABASE_PASSWORD,
	REDIS_DATABASE_SELECT,

	GRPC_SERVICE_CHAT_HOST,
	GRPC_SERVICE_CHAT_PORT string

// Config server Host and Port
func SetEnvironmentVariables() {
	GRPC_SERVICE_CHAT_HOST = os.Getenv("GRPC_SERVICE_CHAT_HOST")
	GRPC_SERVICE_CHAT_PORT = os.Getenv("GRPC_SERVICE_CHAT_PORT")

	REDIS_DATABASE_HOST = os.Getenv("REDIS_DATABASE_HOST")
	REDIS_DATABASE_PORT = os.Getenv("REDIS_DATABASE_PORT")
	REDIS_DATABASE_SELECT = os.Getenv("REDIS_DATABASE_SELECT")
	REDIS_DATABASE_USERNAME = os.Getenv("REDIS_DATABASE_USERNAME")
	REDIS_DATABASE_PASSWORD = os.Getenv("REDIS_DATABASE_PASSWORD")

	DbDriver = os.Getenv("MYSQL_DB_DRIVER")
	DbUser = os.Getenv("MYSQL_DB_USER")
	DbPass = os.Getenv("MYSQL_DB_PASSWORD")
	DbName = os.Getenv("MYSQL_DB_NAME")
	DbHost = os.Getenv("MYSQL_DB_HOST")
	DbPort = os.Getenv("MYSQL_DB_PORT")

	log.Print(`ENVIRONMENT VARIABLE`)
	fmt.Println()

	log.Println(`GRPC_SERVICE_CHAT_HOST=`, GRPC_SERVICE_CHAT_HOST)
	log.Println(`GRPC_SERVICE_CHAT_PORT=`, GRPC_SERVICE_CHAT_PORT)
	fmt.Println()

	log.Println(`REDIS_DATABASE_HOST=`, REDIS_DATABASE_HOST)
	log.Println(`REDIS_DATABASE_PORT=`, REDIS_DATABASE_PORT)
	log.Println(`REDIS_DATABASE_SELECT=`, REDIS_DATABASE_SELECT)
	log.Println(`REDIS_DATABASE_USERNAME=`, REDIS_DATABASE_USERNAME)
	log.Println(`REDIS_DATABASE_PASSWORD=`, REDIS_DATABASE_PASSWORD)
	fmt.Println()

	log.Println(`MYSQL_DB_DRIVER=`, DbDriver)
	log.Println(`MYSQL_DB_USER=`, DbUser)
	log.Println(`MYSQL_DB_PASSWORD=`, DbPass)
	log.Println(`MYSQL_DB_NAME=`, DbName)
	log.Println(`MYSQL_DB_HOST=`, DbHost)
	log.Println(`MYSQL_DB_PORT=`, DbPort)
	fmt.Println()
}
