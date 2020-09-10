package user

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"tublessin/api_gateway/utils/storage"
	"tublessin/common/model"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type UserControllerApi struct {
	UserUsecaseApi UserUsecaseApiInterface
}

func NewLoginControllerApi(userService model.UserClient) *UserControllerApi {
	return &UserControllerApi{UserUsecaseApi: NewUserUsecaseApi(userService)}
}

func (c UserControllerApi) HandleGetUserProfileByID() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		userId := mux.Vars(r)["id"]
		result, err := c.UserUsecaseApi.HandleGetUserProfileByID(userId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.UserResponeMessage{Response: err.Error(), Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (c UserControllerApi) HandleUpdateUserProfilePicture() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		getId := mux.Vars(r)["id"]

		fileName, err := storage.SaveFileToStorage(r, getId, "user")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&model.UserResponeMessage{Response: err.Error(), Code: "500"})
			return
		}

		result, err := c.UserUsecaseApi.HandleUpdateUserProfilePicture(getId, fileName)
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
func (s UserControllerApi) HandleServeUserFile() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		dir, err := os.Getwd()
		if err != nil {
			log.Println(err)
			return
		}

		fileId := mux.Vars(r)["namaFile"]
		fileLocation := filepath.Join(dir, "fileServer", "user", fileId)

		w.Header().Set("Content-Type", "image/jpeg")
		http.ServeFile(w, r, fileLocation)
	}
}

func (c UserControllerApi) HandleUpdateUserProfileByID() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		userId := mux.Vars(r)["id"]
		var userProfile model.UserProfile
		json.NewDecoder(r.Body).Decode(&userProfile)

		result, err := c.UserUsecaseApi.HandleUpdateUserProfileByID(userId, &userProfile)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.UserResponeMessage{Response: err.Error(), Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (c UserControllerApi) HandleUpdateUserLocation() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		userId := mux.Vars(r)["id"]
		var userLocation *model.UserLocation
		json.NewDecoder(r.Body).Decode(&userLocation)

		result, err := c.UserUsecaseApi.HandleUpdateUserLocation(userId, &model.UserProfile{Location: userLocation})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.UserResponeMessage{Response: err.Error(), Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (c UserControllerApi) HandleDeleteUserByID() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		userId := mux.Vars(r)["id"]
		result, err := c.UserUsecaseApi.HandleDeleteUserByID(userId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.UserResponeMessage{Response: err.Error(), Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (c UserControllerApi) HandleGetAllUserSummary() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		keyword := mux.Vars(r)["keyword"]
		page := mux.Vars(r)["page"]
		limit := mux.Vars(r)["limit"]
		status := mux.Vars(r)["status"]
		orderBy := mux.Vars(r)["orderBy"]
		sort := mux.Vars(r)["sort"]

		result, err := c.UserUsecaseApi.HandleGetAllUserSummary(&model.UserPagination{Keyword: keyword, Page: page, Limit: limit, Status: status, OrderBy: orderBy, Sort: sort})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.UserResponeMessage{Response: err.Error(), Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}
