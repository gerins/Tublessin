package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"tublessin/api_gateway/utils"
	"tublessin/common/model"

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
			json.NewEncoder(w).Encode(&model.UserResponeMessage{Response: "User Id Not Found", Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (c UserControllerApi) HandleUpdateUserProfilePicture() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		fileName, err := utils.SaveFileToStorage(r, mux.Vars(r)["id"], "user")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&model.UserResponeMessage{Response: "Uploading Image Failed", Code: "500"})
			return
		}

		fmt.Println(fileName)

		w.WriteHeader(http.StatusOK)
		coverUserId, _ := strconv.Atoi(mux.Vars(r)["id"])
		json.NewEncoder(w).Encode(&model.UserResponeMessage{Response: "Uploading Image Success", Code: "200", Result: &model.UserAccount{Id: int32(coverUserId)}})
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
