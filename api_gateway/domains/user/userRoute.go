package user

import (
	"log"
	"tublessin/common/config"
	"tublessin/common/model"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func InitUserRoute(mainRoute string, r *mux.Router) {
	subRouter := r.PathPrefix(mainRoute).Subrouter()
	userControllerApi := NewLoginControllerApi(connectToServiceUser())
	subRouter.HandleFunc("/profile/detail/{id}", userControllerApi.HandleGetUserProfileByID()).Methods("GET")
	subRouter.HandleFunc("/profile/image/upload/{id}", userControllerApi.HandleUpdateUserProfilePicture()).Methods("POST")
	subRouter.HandleFunc("/file/image/{namaFile}", userControllerApi.HandleServeUserFile()).Methods("GET")
}

// Untuk Connect ke Service-User
func connectToServiceUser() model.UserClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not Connect to User-Service", port, err)
	}

	return model.NewUserClient(conn)
}
