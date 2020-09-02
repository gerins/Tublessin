package user

import (
	"encoding/json"
	"net/http"
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
