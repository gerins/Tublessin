package domain

import (
	"database/sql"
	"log"
	"tublessin/common/model"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	UserRepository UserRepositoryInterface
}

type UserUsecaseInterface interface {
	Login(UserAccount *model.UserAccount) (*model.UserAccount, error)
	RegisterNewUser(UserAccount *model.UserAccount) (*model.UserResponeMessage, error)
	GetUserProfileById(userId int32) (*model.UserResponeMessage, error)
	UpdateUserProfilePicture(userProfile *model.UserProfile) (*model.UserResponeMessage, error)
	UpdateUserProfileByID(userProfile *model.UserProfile) (*model.UserResponeMessage, error)
	UpdateUserLocation(userProfile *model.UserProfile) (*model.UserResponeMessage, error)
}

func NewUserUsecase(db *sql.DB) UserUsecaseInterface {
	return &UserUsecase{NewUserRepository(db)}
}

// Ini Adalah Layer Service dari User-Service, untuk menangani bussiness logic
func (s UserUsecase) Login(UserAccount *model.UserAccount) (*model.UserAccount, error) {
	userDetail, err := s.UserRepository.Login(UserAccount.Username)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return userDetail, nil
}

func (s UserUsecase) RegisterNewUser(UserAccount *model.UserAccount) (*model.UserResponeMessage, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(UserAccount.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	UserAccount.Password = string(hash)

	userResponeMessage, err := s.UserRepository.RegisterNewUser(UserAccount)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return userResponeMessage, nil
}

func (s UserUsecase) GetUserProfileById(userId int32) (*model.UserResponeMessage, error) {
	userResponeMessage, err := s.UserRepository.GetUserProfileById(userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return userResponeMessage, nil
}

func (s UserUsecase) UpdateUserProfilePicture(userProfile *model.UserProfile) (*model.UserResponeMessage, error) {
	UserResponeMessage, err := s.UserRepository.UpdateUserProfilePicture(userProfile)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return UserResponeMessage, nil
}

func (s UserUsecase) UpdateUserProfileByID(userProfile *model.UserProfile) (*model.UserResponeMessage, error) {
	UserResponeMessage, err := s.UserRepository.UpdateUserProfileByID(userProfile)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return UserResponeMessage, nil
}

func (s UserUsecase) UpdateUserLocation(userProfile *model.UserProfile) (*model.UserResponeMessage, error) {
	UserResponeMessage, err := s.UserRepository.UpdateUserLocation(userProfile)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return UserResponeMessage, nil
}
