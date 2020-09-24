package main

import (
	"tublessin/api_gateway/config"
	"tublessin/api_gateway/config/router"
)

func main() {
	// logging.LoggingToFile()
	config.SetEnvironmentVariables()
	muxRouter := router.CreateRouter()
	muxRouter.StrictSlash(true)
	// muxRouter.Handle("/", http.FileServer(http.Dir("./frontend/")))
	router.NewAppRouter(muxRouter).InitRouter()
	router.StartServer(muxRouter)
}
