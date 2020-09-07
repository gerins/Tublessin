package transaction

import (
	"log"
	"tublessin/api_gateway/config"
	"tublessin/api_gateway/middleware"
	"tublessin/common/model"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func InitTransactionRoute(mainRoute string, r *mux.Router) {
	subRouter := r.PathPrefix(mainRoute).Subrouter()
	subRouter.Use(middleware.TokenValidation)
	TransactionControllerApi := NewLoginControllerApi(connectToServiceTransaction())
	subRouter.HandleFunc("/history/get", TransactionControllerApi.HandleGetAllTransactionHistory()).Queries("montirid", "{montirid}", "userid", "{userid}").Methods("GET")
	subRouter.HandleFunc("/add", TransactionControllerApi.HandlePostNewTransaction()).Methods("POST")
	subRouter.HandleFunc("/update/status/{id}", TransactionControllerApi.HandleUpdateTransactionByID()).Methods("POST")
}

// Untuk Connect ke Service-Transaction
func connectToServiceTransaction() model.TransactionClient {
	host := config.SERVICE_TRANSACTION_HOST
	port := config.SERVICE_TRANSACTION_PORT
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not Connect to Transaction-Service", port, err)
	}

	return model.NewTransactionClient(conn)
}
