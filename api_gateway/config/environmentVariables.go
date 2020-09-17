package config

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

var API_GATEWAY_SERVER_HOST,
	API_GATEWAY_SERVER_PORT,

	SERVICE_LOGIN_HOST,
	SERVICE_LOGIN_PORT,

	SERVICE_TRANSACTION_HOST,
	SERVICE_TRANSACTION_PORT,

	SERVICE_MONTIR_HOST,
	SERVICE_MONTIR_PORT,

	SERVICE_USER_HOST,
	SERVICE_USER_PORT,

	SERVICE_CHAT_HOST,
	SERVICE_CHAT_PORT string

// Config server Host and Port
func SetEnvironmentVariables() {
	API_GATEWAY_SERVER_HOST = os.Getenv("API_GATEWAY_SERVER_HOST")
	API_GATEWAY_SERVER_PORT = os.Getenv("API_GATEWAY_SERVER_PORT")

	SERVICE_LOGIN_HOST = os.Getenv("SERVICE_LOGIN_HOST")
	SERVICE_LOGIN_PORT = os.Getenv("SERVICE_LOGIN_PORT")

	SERVICE_TRANSACTION_HOST = os.Getenv("SERVICE_TRANSACTION_HOST")
	SERVICE_TRANSACTION_PORT = os.Getenv("SERVICE_TRANSACTION_PORT")

	SERVICE_MONTIR_HOST = os.Getenv("SERVICE_MONTIR_HOST")
	SERVICE_MONTIR_PORT = os.Getenv("SERVICE_MONTIR_PORT")

	SERVICE_USER_HOST = os.Getenv("SERVICE_USER_HOST")
	SERVICE_USER_PORT = os.Getenv("SERVICE_USER_PORT")

	SERVICE_CHAT_HOST = os.Getenv("SERVICE_CHAT_HOST")
	SERVICE_CHAT_PORT = os.Getenv("SERVICE_CHAT_PORT")

	log.Print(`ENVIRONMENT VARIABLE`)
	fmt.Println()

	log.Println(`API_GATEWAY_SERVER_HOST=`, API_GATEWAY_SERVER_HOST)
	log.Println(`API_GATEWAY_SERVER_PORT=`, API_GATEWAY_SERVER_PORT)
	fmt.Println()

	log.Println(`SERVICE_LOGIN_HOST=`, SERVICE_LOGIN_HOST)
	log.Println(`SERVICE_LOGIN_PORT=`, SERVICE_LOGIN_PORT)
	fmt.Println()

	log.Println(`SERVICE_TRANSACTION_HOST=`, SERVICE_TRANSACTION_HOST)
	log.Println(`SERVICE_TRANSACTION_PORT=`, SERVICE_TRANSACTION_PORT)
	fmt.Println()

	log.Println(`SERVICE_MONTIR_HOST=`, SERVICE_MONTIR_HOST)
	log.Println(`SERVICE_MONTIR_PORT=`, SERVICE_MONTIR_PORT)
	fmt.Println()

	log.Println(`SERVICE_USER_HOST=`, SERVICE_USER_HOST)
	log.Println(`SERVICE_USER_PORT=`, SERVICE_USER_PORT)
	fmt.Println()

	log.Println(`SERVICE_CHAT_HOST=`, SERVICE_CHAT_HOST)
	log.Println(`SERVICE_CHAT_PORT=`, SERVICE_CHAT_PORT)
	fmt.Println()
}
