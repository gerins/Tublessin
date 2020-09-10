package login

import (
	"context"
	"errors"
	"tublessin/common/model"

	log "github.com/sirupsen/logrus"
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
	if montirAccount.Username == "" || montirAccount.Password == "" {
		return nil, errors.New("Username atau password cannot empty")
	}

	result, err := s.LoginService.MontirLogin(context.Background(), montirAccount)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return result, nil
}

func (s LoginUsecaseApi) HandleLoginUser(userAccount *model.UserLoginForm) (*model.LoginResponeMessage, error) {
	if userAccount.Username == "" || userAccount.Password == "" {
		return nil, errors.New("Username atau password cannot empty")
	}

	result, err := s.LoginService.UserLogin(context.Background(), userAccount)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return result, nil
}

func (s LoginUsecaseApi) HandleRegisterNewMontir(montirAccount *model.MontirAccount) (*model.MontirResponeMessage, error) {
	if montirAccount.Username == "" ||
		montirAccount.Password == "" ||
		montirAccount.Profile.Firstname == "" ||
		montirAccount.Profile.Lastname == "" ||
		montirAccount.Profile.Gender == "" ||
		montirAccount.Profile.City == "" ||
		montirAccount.Profile.Email == "" ||
		montirAccount.Profile.PhoneNumber == "" {
		return nil, errors.New("Form Body Cannot empty")
	}

	result, err := s.MontirService.RegisterNewMontir(context.Background(), montirAccount)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return result, nil
}

func (s LoginUsecaseApi) HandleRegisterNewUser(userAccount *model.UserAccount) (*model.UserResponeMessage, error) {
	if userAccount.Username == "" ||
		userAccount.Password == "" ||
		userAccount.Profile.Firstname == "" ||
		userAccount.Profile.Lastname == "" ||
		userAccount.Profile.Gender == "" ||
		userAccount.Profile.Email == "" ||
		userAccount.Profile.PhoneNumber == "" {
		return nil, errors.New("Form Body Cannot empty")
	}

	result, err := s.UserService.RegisterNewUser(context.Background(), userAccount)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return result, nil
}
