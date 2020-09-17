package Chat

import (
	"log"
	"tublessin/api_gateway/config"
	"tublessin/api_gateway/middleware"
	"tublessin/common/model"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func InitChatRoute(mainRoute string, r *mux.Router) {
	subRouter := r.PathPrefix(mainRoute).Subrouter()
	subRouter.Use(middleware.TokenValidation)
	ChatControllerApi := NewChatControllerApi(connectToServiceChat())
	subRouter.HandleFunc("/get", ChatControllerApi.GetConversation()).Queries("senderid", "{senderid}", "receiverid", "{receiverid}").Methods("GET")
	subRouter.HandleFunc("/new", ChatControllerApi.PostNewConversation()).Methods("POST")
}

func connectToServiceChat() model.ChatClient {
	host := config.SERVICE_CHAT_HOST
	port := config.SERVICE_CHAT_PORT
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not Connect to Chat-Service", port, err)
	}

	return model.NewChatClient(conn)
}
