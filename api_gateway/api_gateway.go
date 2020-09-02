package main

import (
	"tublessin/api_gateway/config/router"
)

func main() {
	muxRouter := router.CreateRouter()
	// muxRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))
	router.NewAppRouter(muxRouter).InitRouter()
	router.StartServer(muxRouter)
}
