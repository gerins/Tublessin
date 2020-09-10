package user

import (
	"tublessin/api_gateway/config"
	"tublessin/api_gateway/middleware"
	"tublessin/common/model"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func InitUserRoute(mainRoute string, r *mux.Router) {
	subRouter := r.PathPrefix(mainRoute).Subrouter()
	subRouter.Use(middleware.TokenValidation)
	userControllerApi := NewLoginControllerApi(connectToServiceUser())
	subRouter.HandleFunc("/all", userControllerApi.HandleGetAllUserSummary()).Queries("keyword", "{keyword}", "page", "{page}", "limit", "{limit}", "status", "{status}", "orderBy", "{orderBy}", "sort", "{sort}").Methods("GET")
	subRouter.HandleFunc("/profile/detail/{id}", userControllerApi.HandleGetUserProfileByID()).Methods("GET")
	subRouter.HandleFunc("/profile/update/{id}", userControllerApi.HandleUpdateUserProfileByID()).Methods("POST")
	subRouter.HandleFunc("/profile/update/location/{id}", userControllerApi.HandleUpdateUserLocation()).Methods("POST")
	subRouter.HandleFunc("/profile/delete/{id}", userControllerApi.HandleDeleteUserByID()).Methods("DELETE")
	subRouter.HandleFunc("/profile/image/upload/{id}", userControllerApi.HandleUpdateUserProfilePicture()).Methods("POST")
	subRouter.HandleFunc("/file/image/{namaFile}", userControllerApi.HandleServeUserFile()).Methods("GET")
}

// Untuk Connect ke Service-User
func connectToServiceUser() model.UserClient {
	host := config.SERVICE_LOGIN_HOST
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not Connect to User-Service", port, err)
	}

	return model.NewUserClient(conn)
}
