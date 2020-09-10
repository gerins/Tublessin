package login

import (
	"tublessin/api_gateway/config"
	"tublessin/common/model"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func InitLoginRoute(mainRoute string, r *mux.Router) {
	subRouter := r.PathPrefix(mainRoute).Subrouter()
	loginControllerApi := NewLoginControllerApi(connectToServiceLogin(), connectToServiceMontir(), connectToServiceUser())
	subRouter.HandleFunc("/login/montir", loginControllerApi.HandleLoginMontir()).Methods("POST")
	subRouter.HandleFunc("/login/user", loginControllerApi.HandleLoginUser()).Methods("POST")
	subRouter.HandleFunc("/register/montir", loginControllerApi.HandleRegisterNewMontir()).Methods("POST")
	subRouter.HandleFunc("/register/user", loginControllerApi.HandleRegisterNewUser()).Methods("POST")
}

// Untuk Connect ke Service-Login
func connectToServiceLogin() model.LoginClient {
	host := config.SERVICE_LOGIN_HOST
	port := config.SERVICE_LOGIN_PORT
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not Connect to Login-Service", port, err)
	}

	return model.NewLoginClient(conn)
}

// Untuk Connect ke Service-Montir
func connectToServiceMontir() model.MontirClient {
	host := config.SERVICE_MONTIR_HOST
	port := config.SERVICE_MONTIR_PORT
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not Connect to Montir-Service", port, err)
	}

	return model.NewMontirClient(conn)
}

// Untuk Connect ke Service-User
func connectToServiceUser() model.UserClient {
	host := config.SERVICE_USER_HOST
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not Connect to User-Service", port, err)
	}

	return model.NewUserClient(conn)
}
