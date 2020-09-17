package Chat

import (
	"encoding/json"
	"net/http"
	"tublessin/api_gateway/utils/message"
	"tublessin/common/model"

	"github.com/gorilla/mux"
)

type ChatControllerApi struct {
	ChatUsecaseApi ChatUsecaseApiInterface
}

func NewChatControllerApi(ChatService model.ChatClient) *ChatControllerApi {
	return &ChatControllerApi{ChatUsecaseApi: NewChatUsecaseApi(ChatService)}
}

// Nangkep request dari depan yang nanti nya akan di teruskan ke Chat-Service
// Disini cuman nge parse data json yang masuk
func (c ChatControllerApi) GetConversation() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var chatPayload model.ChatPayload
		chatPayload.SenderId = mux.Vars(r)["senderid"]
		chatPayload.ReceiverId = mux.Vars(r)["receiverid"]

		result, err := c.ChatUsecaseApi.GetConversation(&chatPayload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Respone("Get Conversation Failed", http.StatusBadRequest, err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Respone("Get Conversation Success", http.StatusOK, result))
	}
}

func (c ChatControllerApi) PostNewConversation() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var chatPayload model.ChatPayload
		err := json.NewDecoder(r.Body).Decode(&chatPayload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Respone("Post Conversation Failed", http.StatusBadRequest, err.Error()))
			return
		}

		result, err := c.ChatUsecaseApi.PostNewConversation(&chatPayload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(message.Respone("Post Conversation Failed", http.StatusBadRequest, err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message.Respone("Post Conversation Success", http.StatusOK, result))
	}
}
