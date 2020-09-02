package montir

import (
	"encoding/json"
	"net/http"
	"tublessin/common/model"

	"github.com/gorilla/mux"
)

type MontirControllerApi struct {
	MontirUsecaseApi MontirUsecaseApiInterface
}

func NewLoginControllerApi(montirService model.MontirClient) *MontirControllerApi {
	return &MontirControllerApi{MontirUsecaseApi: NewMontirUsecaseApi(montirService)}
}

func (c MontirControllerApi) HandleGetMontirProfileByID() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		montirId := mux.Vars(r)["id"]
		result, err := c.MontirUsecaseApi.HandleGetMontirProfileByID(montirId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: "Id Not Found", Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}
