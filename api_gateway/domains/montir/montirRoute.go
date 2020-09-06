package montir

import (
	"log"
	"tublessin/common/config"
	"tublessin/common/model"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func InitMontirRoute(mainRoute string, r *mux.Router) {
	subRouter := r.PathPrefix(mainRoute).Subrouter()
	montirControllerApi := NewMontirControllerApi(connectToServiceMontir())
	subRouter.HandleFunc("/all", montirControllerApi.HandleGetAllMontirSummary()).Queries("keyword", "{keyword}", "page", "{page}", "limit", "{limit}", "status", "{status}", "orderBy", "{orderBy}", "sort", "{sort}").Methods("GET")
	subRouter.HandleFunc("/profile/detail/{id}", montirControllerApi.HandleGetMontirProfileByID()).Methods("GET")
	subRouter.HandleFunc("/profile/update/{id}", montirControllerApi.HandleUpdateMontirProfileByID()).Methods("POST")
	subRouter.HandleFunc("/profile/delete/{id}", montirControllerApi.HandleDeleteMontirByID()).Methods("DELETE")
	subRouter.HandleFunc("/profile/update/location/{id}", montirControllerApi.HandleUpdateMontirLocation()).Methods("POST")
	subRouter.HandleFunc("/profile/image/upload/{id}", montirControllerApi.HandleUpdateMontirProfilePicture()).Methods("POST")
	subRouter.HandleFunc("/file/image/{namaFile}", montirControllerApi.HandleServeMontirFile()).Methods("GET")
	subRouter.HandleFunc("/find/nearby", montirControllerApi.HandleGetAllActiveMontirWithLocation()).Queries("lat", "{lat}", "long", "{long}").Methods("GET")
}

// Untuk Connect ke Service-Montir
func connectToServiceMontir() model.MontirClient {
	port := config.SERVICE_MONTIR_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not Connect to Montir-Service", port, err)
	}

	return model.NewMontirClient(conn)
}
