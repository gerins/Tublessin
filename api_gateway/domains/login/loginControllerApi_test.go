package login

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"tublessin/common/model"
)

func TestNewLoginControllerApi(t *testing.T) {
	type args struct {
		loginService  model.LoginClient
		montirService model.MontirClient
		userService   model.UserClient
	}
	tests := []struct {
		name string
		args args
		want *LoginControllerApi
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLoginControllerApi(tt.args.loginService, tt.args.montirService, tt.args.userService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoginControllerApi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoginControllerApi_HandleLoginMontir(t *testing.T) {
	req, err := http.NewRequest("POST", "/account/login/montir", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	LoginApi := NewLoginControllerApi(connectToServiceLogin(), connectToServiceMontir(), connectToServiceUser())
	handler := http.HandlerFunc(LoginApi.HandleLoginMontir())

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"username": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestLoginControllerApi_HandleLoginUser(t *testing.T) {
	tests := []struct {
		name string
		c    LoginControllerApi
		want func(w http.ResponseWriter, r *http.Request)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.HandleLoginUser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginControllerApi.HandleLoginUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoginControllerApi_HandleRegisterNewMontir(t *testing.T) {
	tests := []struct {
		name string
		c    LoginControllerApi
		want func(w http.ResponseWriter, r *http.Request)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.HandleRegisterNewMontir(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginControllerApi.HandleRegisterNewMontir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoginControllerApi_HandleRegisterNewUser(t *testing.T) {
	tests := []struct {
		name string
		c    LoginControllerApi
		want func(w http.ResponseWriter, r *http.Request)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.HandleRegisterNewUser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginControllerApi.HandleRegisterNewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
