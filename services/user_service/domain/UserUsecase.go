package domain

import (
	"database/sql"
	"errors"
	"tublessin/common/model"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	UserRepository UserRepositoryInterface
}

type UserUsecaseInterface interface {
	Login(UserAccount *model.UserAccount) (*model.UserAccount, error)
	RegisterNewUser(UserAccount *model.UserAccount) (*model.UserResponeMessage, error)
	GetUserProfileById(userId string) (*model.UserResponeMessage, error)
}

func NewUserUsecase(db *sql.DB) UserUsecaseInterface {
	return &UserUsecase{NewUserRepository(db)}
}

// Ini Adalah Layer Service dari User-Service, untuk menangani bussiness logic
func (s UserUsecase) Login(UserAccount *model.UserAccount) (*model.UserAccount, error) {
	userDetail, err := s.UserRepository.Login(UserAccount.Username)
	if err != nil {
		return nil, err
	}

	return userDetail, nil
}

func (s UserUsecase) RegisterNewUser(UserAccount *model.UserAccount) (*model.UserResponeMessage, error) {
	if UserAccount == nil || UserAccount.Profile == nil {
		return nil, errors.New("Body Cannot Empty")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(UserAccount.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	UserAccount.Password = string(hash)

	userResponeMessage, err := s.UserRepository.RegisterNewUser(UserAccount)
	if err != nil {
		return nil, err
	}

	return userResponeMessage, nil
}

func (s UserUsecase) GetUserProfileById(userId string) (*model.UserResponeMessage, error) {
	userResponeMessage, err := s.UserRepository.GetUserProfileById(userId)
	if err != nil {
		return nil, err
	}
	return userResponeMessage, nil
}
