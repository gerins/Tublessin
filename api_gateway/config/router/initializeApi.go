package router

import (
	"fmt"
	"net/http"
	Chat "tublessin/api_gateway/domains/chat"
	"tublessin/api_gateway/domains/login"
	"tublessin/api_gateway/domains/montir"
	"tublessin/api_gateway/domains/transaction"
	"tublessin/api_gateway/domains/user"

	"github.com/gorilla/mux"
)

type ConfigRouter struct {
	Router *mux.Router
}

// Disini tempat inisialisasi API yang akan di publish keluar
func (ar *ConfigRouter) InitRouter() {
	login.InitLoginRoute(LOGIN_MAIN_ROUTE, ar.Router)
	montir.InitMontirRoute(MONTIR_MAIN_ROUTE, ar.Router)
	user.InitUserRoute(USER_MAIN_ROUTE, ar.Router)
	transaction.InitTransactionRoute(TRANSACTION_MAIN_ROUTE, ar.Router)
	Chat.InitChatRoute(CHAT_MAIN_ROUTE, ar.Router)
	ar.Router.NotFoundHandler = http.HandlerFunc(notFound)
}

// NotFound Handler biasa
func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, `<h1>404 Status Not Found</h1>`)
}

// NewAppRouter for creating new Config Router
func NewAppRouter(r *mux.Router) *ConfigRouter {
	return &ConfigRouter{Router: r}
}
