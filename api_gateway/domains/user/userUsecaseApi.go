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
