package login

import (
	"context"
	"log"
	"tublessin/common/model"
)

type LoginUsecaseApi struct {
	LoginService  model.LoginClient
	MontirService model.MontirClient
	UserService   model.UserClient
}

type LoginUsecaseApiInterface interface {
	HandleLoginMontir(montirAccount *model.MontirLoginForm) (*model.LoginResponeMessage, error)
	HandleLoginUser(userAccount *model.UserLoginForm) (*model.LoginResponeMessage, error)
	HandleRegisterNewMontir(montirAccount *model.MontirAccount) (*model.MontirResponeMessage, error)
	HandleRegisterNewUser(userAccount *model.UserAccount) (*model.UserResponeMessage, error)
}

func NewLoginUsecaseApi(loginService model.LoginClient, montirService model.MontirClient, userService model.UserClient) LoginUsecaseApiInterface {
	return LoginUsecaseApi{LoginService: loginService, MontirService: montirService, UserService: userService}
}

// Dioper ke Login-Service untuk ditangani ada di folder services/login_services/domain/LoginController.go
func (s LoginUsecaseApi) HandleLoginMontir(montirAccount *model.MontirLoginForm) (*model.LoginResponeMessage, error) {
	result, err := s.LoginService.MontirLogin(context.Background(), montirAccount)
	if err != nil {
		log.Println(err.Error())
		return &model.LoginResponeMessage{
			Message: "Username atau Password Salah",
			Token:   "0",
			Account: nil,
		}, err
	}

	return result, nil
}

func (s LoginUsecaseApi) HandleLoginUser(userAccount *model.UserLoginForm) (*model.LoginResponeMessage, error) {
	result, err := s.LoginService.UserLogin(context.Background(), userAccount)
	if err != nil {
		log.Println(err.Error())
		return &model.LoginResponeMessage{
			Message: "Username atau Password Salah",
			Token:   "0",
			Account: nil,
		}, err
	}

	return result, nil
}

func (s LoginUsecaseApi) HandleRegisterNewMontir(montirAccount *model.MontirAccount) (*model.MontirResponeMessage, error) {
	result, err := s.MontirService.RegisterNewMontir(context.Background(), montirAccount)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return result, nil
}

func (s LoginUsecaseApi) HandleRegisterNewUser(userAccount *model.UserAccount) (*model.UserResponeMessage, error) {
	result, err := s.UserService.RegisterNewUser(context.Background(), userAccount)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return result, nil
}
