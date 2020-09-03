package montir

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"tublessin/api_gateway/utils"
	"tublessin/common/model"

	"github.com/gorilla/mux"
)

type MontirControllerApi struct {
	MontirUsecaseApi MontirUsecaseApiInterface
}

func NewMontirControllerApi(montirService model.MontirClient) *MontirControllerApi {
	return &MontirControllerApi{MontirUsecaseApi: NewMontirUsecaseApi(montirService)}
}

func (c MontirControllerApi) HandleGetMontirProfileByID() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		montirId := mux.Vars(r)["id"]
		result, err := c.MontirUsecaseApi.HandleGetMontirProfileByID(montirId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: "Montir Id Not Found", Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (c MontirControllerApi) HandleUpdateMontirProfilePicture() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		getId := mux.Vars(r)["id"]
		fileName, err := utils.SaveFileToStorage(r, getId, "montir")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: "Uploading Image Failed", Code: "500"})
			return
		}

		result, err := c.MontirUsecaseApi.HandleUpdateMontirProfilePicture(getId, fileName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: "Uploading Image Failed", Code: "500"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

// Ini function untuk serve yang ada di harddisk
func (s MontirControllerApi) HandleServeMontirFile() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		dir, err := os.Getwd()
		if err != nil {
			log.Println(err)
			return
		}

		fileId := mux.Vars(r)["namaFile"]
		fileLocation := filepath.Join(dir, "fileServer", "montir", fileId)

		w.Header().Set("Content-Type", "image/jpeg")
		http.ServeFile(w, r, fileLocation)
	}
}
