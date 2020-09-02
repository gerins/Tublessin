package router

import (
	"fmt"
	"net/http"
	"tublessin/api_gateway/domains/login"
	"tublessin/api_gateway/domains/montir"

	"github.com/gorilla/mux"
)

type ConfigRouter struct {
	Router *mux.Router
}

// Disini tempat inisialisasi API yang akan di publish keluar
func (ar *ConfigRouter) InitRouter() {
	login.InitLoginRoute(LOGIN_MAIN_ROUTE, ar.Router)
	montir.InitMontirRoute(MONTIR_MAIN_ROUTE, ar.Router)
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
