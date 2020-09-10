package montir

import (
	"tublessin/api_gateway/config"
	"tublessin/api_gateway/middleware"
	"tublessin/common/model"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func InitMontirRoute(mainRoute string, r *mux.Router) {
	subRouter := r.PathPrefix(mainRoute).Subrouter()
	subRouter.Use(middleware.TokenValidation)
	montirControllerApi := NewMontirControllerApi(connectToServiceMontir())
	subRouter.HandleFunc("/all", montirControllerApi.HandleGetAllMontirSummary()).Queries("keyword", "{keyword}", "page", "{page}", "limit", "{limit}", "status", "{status}", "orderBy", "{orderBy}", "sort", "{sort}").Methods("GET")
	subRouter.HandleFunc("/profile/detail/{id}", montirControllerApi.HandleGetMontirProfileByID()).Methods("GET")
	subRouter.HandleFunc("/profile/update/{id}", montirControllerApi.HandleUpdateMontirProfileByID()).Methods("POST")
	subRouter.HandleFunc("/profile/update/status/{id}", montirControllerApi.HandleUpdateMontirStatusByID()).Methods("POST")
	subRouter.HandleFunc("/profile/update/location/{id}", montirControllerApi.HandleUpdateMontirLocation()).Methods("POST")
	subRouter.HandleFunc("/profile/image/upload/{id}", montirControllerApi.HandleUpdateMontirProfilePicture()).Methods("POST")
	subRouter.HandleFunc("/location/{id}", montirControllerApi.HandleGetMontirLocation()).Methods("GET")
	subRouter.HandleFunc("/profile/delete/{id}", montirControllerApi.HandleDeleteMontirByID()).Methods("DELETE")
	subRouter.HandleFunc("/file/image/{namaFile}", montirControllerApi.HandleServeMontirFile()).Methods("GET")
	subRouter.HandleFunc("/find/nearby", montirControllerApi.HandleGetAllActiveMontirWithLocation()).Queries("lat", "{lat}", "long", "{long}").Methods("GET")
	subRouter.HandleFunc("/rating/add/{id}", montirControllerApi.HandleInsertNewMontirRating()).Methods("POST")
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
