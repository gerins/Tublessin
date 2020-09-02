package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var serverHost, serverPort string

// CreateRouter for creating new Route
func CreateRouter() *mux.Router {
	return mux.NewRouter()
}

// StartServer routing
func StartServer(r *mux.Router) {
	configServer()
	log.Println("Server Start at http://" + serverHost + ":" + serverPort)
	printRouteList()
	http.ListenAndServe(serverHost+":"+serverPort, r)
}

// Config server Host and Port
func configServer() {
	serverHost = "localhost"
	serverPort = "8080"
}

func printRouteList() {
	log.Println("======== ROUTE LIST ========")
	log.Println("http://"+serverHost+":"+serverPort+LOGIN_MAIN_ROUTE+"/login/montir", "(POST)")
	log.Println("http://"+serverHost+":"+serverPort+LOGIN_MAIN_ROUTE+"/login/user", "(POST)")
	log.Println("http://"+serverHost+":"+serverPort+LOGIN_MAIN_ROUTE+"/register/montir", "(POST)")
	log.Println("http://"+serverHost+":"+serverPort+LOGIN_MAIN_ROUTE+"/register/user", "(POST)")
	log.Println("http://"+serverHost+":"+serverPort+MONTIR_MAIN_ROUTE+"/profile/detail/{id}", "(GET)")
	log.Println("http://"+serverHost+":"+serverPort+USER_MAIN_ROUTE+"/profile/detail/{id}", "(GET)")
}
