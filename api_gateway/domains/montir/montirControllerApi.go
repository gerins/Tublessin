package montir

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"tublessin/api_gateway/utils/storage"
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
			json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: err.Error(), Code: "400"})
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
		fileName, err := storage.SaveFileToStorage(r, getId, "montir")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: err.Error(), Code: "500"})
			return
		}

		result, err := c.MontirUsecaseApi.HandleUpdateMontirProfilePicture(getId, fileName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: err.Error(), Code: "500"})
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

		montirId := mux.Vars(r)["id"]
		var montirProfile model.MontirProfile
		json.NewDecoder(r.Body).Decode(&montirProfile)

		result, err := c.MontirUsecaseApi.HandleUpdateMontirProfileByID(montirId, &montirProfile)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: err.Error(), Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (c MontirControllerApi) HandleUpdateMontirLocation() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		montirId := mux.Vars(r)["id"]
		var montirLocation *model.MontirLocation
		json.NewDecoder(r.Body).Decode(&montirLocation)

		result, err := c.MontirUsecaseApi.HandleUpdateMontirLocation(montirId, &model.MontirProfile{Location: montirLocation})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: err.Error(), Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (c MontirControllerApi) HandleUpdateMontirStatusByID() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		montirId := mux.Vars(r)["id"]
		var montirStatus *model.MontirStatus
		json.NewDecoder(r.Body).Decode(&montirStatus)

		result, err := c.MontirUsecaseApi.HandleUpdateMontirStatusByID(montirId, montirStatus)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: err.Error(), Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (c MontirControllerApi) HandleGetAllActiveMontirWithLocation() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		getLatitude := mux.Vars(r)["lat"]
		getLongitude := mux.Vars(r)["long"]

		result, err := c.MontirUsecaseApi.HandleGetAllActiveMontirWithLocation(getLatitude, getLongitude)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: err.Error(), Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (c MontirControllerApi) HandleDeleteMontirByID() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		montirId := mux.Vars(r)["id"]
		result, err := c.MontirUsecaseApi.HandleDeleteMontirByID(montirId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: err.Error(), Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (c MontirControllerApi) HandleGetAllMontirSummary() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		keyword := mux.Vars(r)["keyword"]
		page := mux.Vars(r)["page"]
		limit := mux.Vars(r)["limit"]
		status := mux.Vars(r)["status"]
		orderBy := mux.Vars(r)["orderBy"]
		sort := mux.Vars(r)["sort"]

		result, err := c.MontirUsecaseApi.HandleGetAllMontirSummary(&model.MontirPagination{Keyword: keyword, Page: page, Limit: limit, Status: status, OrderBy: orderBy, Sort: sort})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: err.Error(), Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}
