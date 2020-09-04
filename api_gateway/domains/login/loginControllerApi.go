package login

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"tublessin/common/model"
)

type LoginControllerApi struct {
	LoginUsecaseApi LoginUsecaseApiInterface
}

func NewLoginControllerApi(loginService model.LoginClient, montirService model.MontirClient, userService model.UserClient) *LoginControllerApi {
	return &LoginControllerApi{LoginUsecaseApi: NewLoginUsecaseApi(loginService, montirService, userService)}
}

// Nangkep request dari depan yang nanti nya akan di teruskan ke Login-Service
// Disini cuman nge parse data json yang masuk
func (c LoginControllerApi) HandleLoginMontir() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var montirAccount model.MontirLoginForm
		json.NewDecoder(r.Body).Decode(&montirAccount)

		log.Print(`LOGIN -> `, montirAccount.Username)

		result, err := c.LoginUsecaseApi.HandleLoginMontir(&montirAccount)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.LoginResponeMessage{
				Message: "Username atau Password Salah",
				Token:   "0",
				Account: nil,
			})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (c LoginControllerApi) HandleLoginUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var userAccount model.UserLoginForm
		json.NewDecoder(r.Body).Decode(&userAccount)

		log.Print(`LOGIN -> `, userAccount.Username)

		result, err := c.LoginUsecaseApi.HandleLoginUser(&userAccount)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&model.LoginResponeMessage{
				Message: "Username atau Password Salah",
				Token:   "0",
				Account: nil,
			})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (c LoginControllerApi) HandleRegisterNewMontir() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var montirAccount model.MontirAccount
		json.NewDecoder(r.Body).Decode(&montirAccount)

		result, err := c.LoginUsecaseApi.HandleRegisterNewMontir(&montirAccount)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if strings.Contains(err.Error(), "username_UNIQUE") {
				json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: "Username Sudah Digunakan", Code: "900"})
				return
			} else if strings.Contains(err.Error(), "phone_number_UNIQUE") {
				json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: "Nomor Telefon Sudah Digunakan", Code: "800"})
				return
			} else if strings.Contains(err.Error(), "email_UNIQUE") {
				json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: "Email Sudah Digunakan", Code: "700"})
				return
			}
			json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: err.Error(), Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func (c LoginControllerApi) HandleRegisterNewUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var userAccount model.UserAccount
		json.NewDecoder(r.Body).Decode(&userAccount)

		result, err := c.LoginUsecaseApi.HandleRegisterNewUser(&userAccount)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			if strings.Contains(err.Error(), "username_UNIQUE") {
				json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: "Username Sudah Digunakan", Code: "900"})
				return
			} else if strings.Contains(err.Error(), "phone_number_UNIQUE") {
				json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: "Nomor Telefon Sudah Digunakan", Code: "800"})
				return
			} else if strings.Contains(err.Error(), "email_UNIQUE") {
				json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: "Email Sudah Digunakan", Code: "700"})
				return
			}
			json.NewEncoder(w).Encode(&model.MontirResponeMessage{Response: err.Error(), Code: "400"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}
