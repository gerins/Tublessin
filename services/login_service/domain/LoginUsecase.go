package domain

import (
	"context"
	"log"
	"tublessin/common/model"
)

type LoginUsecase struct {
	MontirService model.MontirClient
	UserService   model.UserClient
}

type LoginUsecaseInterface interface {
	MontirLogin(montirAccount *model.MontirAccount) (*model.MontirAccount, error)
	UserLogin(userAccount *model.UserAccount) (*model.UserAccount, error)
}

func NewLoginUsecase(clientMontir model.MontirClient, clientUser model.UserClient) LoginUsecaseInterface {
	return &LoginUsecase{MontirService: clientMontir, UserService: clientUser}
}

// Karna Login-Service tidak bisa akses langsung ke Database Montir, jadi harus dioper ke Montir-Service
func (s LoginUsecase) MontirLogin(montirAccount *model.MontirAccount) (*model.MontirAccount, error) {
	result, err := s.MontirService.Login(context.Background(), montirAccount)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return result, nil
}

func (s LoginUsecase) UserLogin(userAccount *model.UserAccount) (*model.UserAccount, error) {
	result, err := s.UserService.Login(context.Background(), userAccount)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return result, nil
}
