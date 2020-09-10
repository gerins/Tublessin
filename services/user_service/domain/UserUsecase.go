package domain

import (
	"database/sql"
	"strconv"
	"tublessin/common/model"

	log "github.com/sirupsen/logrus"

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
	DeleteUserByID(UserAccount *model.UserAccount) (*model.UserResponeMessage, error)
	GetAllUserSummary(query *model.UserPagination) (*model.UserResponeMessage, error)
}

func NewUserUsecase(db *sql.DB) UserUsecaseInterface {
	return &UserUsecase{NewUserRepository(db)}
}

// Ini Adalah Layer Service dari User-Service, untuk menangani bussiness logic
func (s UserUsecase) Login(UserAccount *model.UserAccount) (*model.UserAccount, error) {
	userDetail, err := s.UserRepository.Login(UserAccount.Username, "A")
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

func (s UserUsecase) DeleteUserByID(UserAccount *model.UserAccount) (*model.UserResponeMessage, error) {
	userResponeMessage, err := s.UserRepository.DeleteUserByID(UserAccount)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return userResponeMessage, nil
}

func (s UserUsecase) GetAllUserSummary(query *model.UserPagination) (*model.UserResponeMessage, error) {
	result, countItem, err := s.UserRepository.GetAllUserSummary(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &model.UserResponeMessage{Response: "Success", Code: "200", TotalUser: strconv.Itoa(countItem), List: result}, nil
}
