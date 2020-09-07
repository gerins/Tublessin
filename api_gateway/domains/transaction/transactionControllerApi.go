package transaction

import (
	"encoding/json"
	"net/http"
	"tublessin/api_gateway/utils/message"
	"tublessin/common/model"

	"github.com/gorilla/mux"
)

type TransactionControllerApi struct {
	TransactionUsecaseApi TransactionUsecaseApiInterface
}

func NewLoginControllerApi(TransactionService model.TransactionClient) *TransactionControllerApi {
	return &TransactionControllerApi{TransactionUsecaseApi: NewTransactionUsecaseApi(TransactionService)}
}

func (c TransactionControllerApi) HandleGetAllTransactionHistory() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		montirId := mux.Vars(r)["montirid"]
		userId := mux.Vars(r)["userid"]

		result, err := c.TransactionUsecaseApi.HandleGetAllTransactionHistory(montirId, userId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Respone("Get All Transaction Failed", http.StatusBadRequest, err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Respone("Get All Transaction Success", http.StatusOK, result))
	}
}

func (c TransactionControllerApi) HandlePostNewTransaction() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var transaction model.TransactionHistory
		json.NewDecoder(r.Body).Decode(&transaction)

		result, err := c.TransactionUsecaseApi.HandlePostNewTransaction(&transaction)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Respone("Post Transaction Failed", http.StatusBadRequest, err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Respone("Post Transaction Success", http.StatusOK, result))
	}
}

func (c TransactionControllerApi) HandleUpdateTransactionByID() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var transaction model.TransactionHistory
		transId := mux.Vars(r)["id"]
		json.NewDecoder(r.Body).Decode(&transaction)
		transaction.Id = transId

		result, err := c.TransactionUsecaseApi.HandleUpdateTransactionByID(&transaction)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Respone("Update Transaction Failed", http.StatusBadRequest, err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Respone("Update Transaction Success", http.StatusOK, result))
	}
}
