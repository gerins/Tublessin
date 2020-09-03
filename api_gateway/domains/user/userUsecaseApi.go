package user

import (
	"context"
	"log"
	"strconv"
	"tublessin/common/model"
)

type UserUsecaseApi struct {
	UserService model.UserClient
}

type UserUsecaseApiInterface interface {
	HandleGetUserProfileByID(UserId string) (*model.UserResponeMessage, error)
	HandleUpdateUserProfilePicture(userId, fileName string) (*model.UserResponeMessage, error)
	HandleUpdateUserProfileByID(UserProfile *model.UserProfile) (*model.UserResponeMessage, error)
	HandleUpdateUserLocation(UserProfile *model.UserProfile) (*model.UserResponeMessage, error)
}

func NewUserUsecaseApi(UserService model.UserClient) UserUsecaseApiInterface {
	return UserUsecaseApi{UserService: UserService}
}

func (s UserUsecaseApi) HandleGetUserProfileByID(userId string) (*model.UserResponeMessage, error) {
	id, _ := strconv.Atoi(userId)
	UserAccountWithId := &model.UserAccount{Id: int32(id)}

	userResponeMessage, err := s.UserService.GetUserProfileById(context.Background(), UserAccountWithId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return userResponeMessage, nil
}

func (s UserUsecaseApi) HandleUpdateUserProfilePicture(userId, fileName string) (*model.UserResponeMessage, error) {
	convertIdToInt, _ := strconv.Atoi(userId)
	userResponeMessage, err := s.UserService.UpdateUserProfilePicture(context.Background(), &model.UserProfile{Id: int32(convertIdToInt), ImageURL: fileName})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return userResponeMessage, nil
}

func (s UserUsecaseApi) HandleUpdateUserProfileByID(UserProfile *model.UserProfile) (*model.UserResponeMessage, error) {
	UserResponeMessage, err := s.UserService.UpdateUserProfileById(context.Background(), UserProfile)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return UserResponeMessage, nil
}

func (s UserUsecaseApi) HandleUpdateUserLocation(UserProfile *model.UserProfile) (*model.UserResponeMessage, error) {
	UserResponeMessage, err := s.UserService.UpdateUserLocation(context.Background(), UserProfile)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return UserResponeMessage, nil
}
