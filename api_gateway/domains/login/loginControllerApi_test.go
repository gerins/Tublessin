package login

import (
	"net/http"
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
	tests := []struct {
		name string
		c    LoginControllerApi
		want func(w http.ResponseWriter, r *http.Request)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.HandleLoginMontir(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginControllerApi.HandleLoginMontir() = %v, want %v", got, tt.want)
			}
		})
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
