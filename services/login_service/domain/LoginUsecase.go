package domain

import (
	"context"
	"errors"
	"log"
	"strconv"
	"tublessin/common/model"
	"tublessin/services/login_service/token"

	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase struct {
	MontirService model.MontirClient
	UserService   model.UserClient
}

type LoginUsecaseInterface interface {
	MontirLogin(montirAccount *model.MontirAccount) (*model.MontirAccount, string, error)
	UserLogin(userAccount *model.UserAccount) (*model.UserAccount, string, error)
}

func NewLoginUsecase(clientMontir model.MontirClient, clientUser model.UserClient) LoginUsecaseInterface {
	return &LoginUsecase{MontirService: clientMontir, UserService: clientUser}
}

// Karna Login-Service tidak bisa akses langsung ke Database Montir, jadi harus dioper ke Montir-Service
func (s LoginUsecase) MontirLogin(montirAccount *model.MontirAccount) (*model.MontirAccount, string, error) {
	result, err := s.MontirService.Login(context.Background(), montirAccount)
	if err != nil {
		log.Println(err.Error())
		return nil, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(montirAccount.Password))
	if err != nil {
		log.Println(err.Error())
		return nil, "", errors.New("Username atau Password salah")
	}

	generatedToken := token.GenerateToken(montirAccount.Username, strconv.Itoa(int(result.Id)), 3600*24)
	return result, generatedToken, nil
}

func (s LoginUsecase) UserLogin(userAccount *model.UserAccount) (*model.UserAccount, string, error) {
	result, err := s.UserService.Login(context.Background(), userAccount)
	if err != nil {
		log.Println(err.Error())
		return nil, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(userAccount.Password))
	if err != nil {
		log.Println(err.Error())
		return nil, "", errors.New("Username atau Password salah")
	}

	generatedToken := token.GenerateToken(userAccount.Username, strconv.Itoa(int(result.Id)), 3600*24)
	return result, generatedToken, nil
}
