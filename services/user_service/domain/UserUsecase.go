package domain

import (
	"context"
	"database/sql"
	"encoding/json"
	"strconv"
	"time"
	"tublessin/common/model"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	UserRepository UserRepositoryInterface
	RedisDatabase  *redis.Client
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

func NewUserUsecase(db *sql.DB, rdb *redis.Client) UserUsecaseInterface {
	return &UserUsecase{UserRepository: NewUserRepository(db), RedisDatabase: rdb}
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
	value, err := s.RedisDatabase.Get(context.Background(), strconv.Itoa(int(userId))).Result()
	if err == nil {
		var userRespone model.UserResponeMessage
		json.Unmarshal([]byte(value), &userRespone)
		if err != nil {
			log.Println("Something wrong when Unmarshal data to User Profile", err)
		}
		return &userRespone, nil
	} else if err != nil && err != redis.Nil {
		log.Println("Something wrong when read data from Redis", err)
	}

	userResponeMessage, err := s.UserRepository.GetUserProfileById(userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result, err := json.Marshal(userResponeMessage)
	err = s.RedisDatabase.Set(context.Background(), strconv.Itoa(int(userId)), result, 30*time.Second).Err()
	if err != nil {
		log.Println("Cannot save User profile to Redis", err)
	}

	return userResponeMessage, nil
}

func (s UserUsecase) UpdateUserProfilePicture(userProfile *model.UserProfile) (*model.UserResponeMessage, error) {
	UserResponeMessage, err := s.UserRepository.UpdateUserProfilePicture(userProfile)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = s.RedisDatabase.Set(context.Background(), strconv.Itoa(int(userProfile.Id)), UserResponeMessage, 1*time.Second).Err()
	if err != nil {
		log.Println("Cannot Remove User profile From Redis", err)
	}

	return UserResponeMessage, nil
}

func (s UserUsecase) UpdateUserProfileByID(userProfile *model.UserProfile) (*model.UserResponeMessage, error) {
	UserResponeMessage, err := s.UserRepository.UpdateUserProfileByID(userProfile)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = s.RedisDatabase.Set(context.Background(), strconv.Itoa(int(userProfile.Id)), UserResponeMessage, 1*time.Second).Err()
	if err != nil {
		log.Println("Cannot Remove User profile From Redis", err)
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
