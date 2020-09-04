package montir

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
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

func (c MontirControllerApi) HandleUpdateMontirProfileByID() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var montirProfile model.MontirProfile
		json.NewDecoder(r.Body).Decode(&montirProfile)
		convertId, _ := strconv.Atoi(mux.Vars(r)["id"])
		montirProfile.Id = int32(convertId)

		result, err := c.MontirUsecaseApi.HandleUpdateMontirProfileByID(&montirProfile)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: "Updating Montir Profile Failed", Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (c MontirControllerApi) HandleUpdateMontirLocation() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var montirProfile model.MontirProfile
		var montirLocation *model.MontirLocation
		json.NewDecoder(r.Body).Decode(&montirLocation)

		convertId, _ := strconv.Atoi(mux.Vars(r)["id"])
		montirProfile.Id = int32(convertId)
		montirProfile.Location = montirLocation

		result, err := c.MontirUsecaseApi.HandleUpdateMontirLocation(&montirProfile)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: "Updating Montir Location Failed", Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (c MontirControllerApi) HandleGetAllActiveMontirWithLocation() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var userLocation model.RequestActiveMontir
		doubleLatitude, _ := strconv.Atoi(mux.Vars(r)["lat"])
		doubleLongitude, _ := strconv.Atoi(mux.Vars(r)["long"])
		userLocation.Latitude = float64(doubleLatitude)
		userLocation.Longitude = float64(doubleLongitude)

		result, err := c.MontirUsecaseApi.HandleGetAllActiveMontirWithLocation(&userLocation)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: "Search Nearby Montir Failed", Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}
