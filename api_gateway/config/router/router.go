package router

import (
	"log"
	"net/http"
	"tublessin/api_gateway/config"

	"github.com/gorilla/mux"
)

// CreateRouter for creating new Route
func CreateRouter() *mux.Router {
	return mux.NewRouter()
}

// StartServer routing
func StartServer(r *mux.Router) {
	log.Println("Server Start at http://" + config.API_GATEWAY_SERVER_HOST + ":" + config.API_GATEWAY_SERVER_PORT)
	http.ListenAndServe(config.API_GATEWAY_SERVER_HOST+":"+config.API_GATEWAY_SERVER_PORT, r)
}
