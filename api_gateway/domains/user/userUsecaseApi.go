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
	HandleUpdateUserProfileByID(userId string, UserProfile *model.UserProfile) (*model.UserResponeMessage, error)
	HandleUpdateUserLocation(userId string, userProfile *model.UserProfile) (*model.UserResponeMessage, error)
	HandleDeleteUserByID(userId string) (*model.UserResponeMessage, error)
	HandleGetAllUserSummary(query *model.UserPagination) (*model.UserResponeMessage, error)
}

func NewUserUsecaseApi(UserService model.UserClient) UserUsecaseApiInterface {
	return UserUsecaseApi{UserService: UserService}
}

func (s UserUsecaseApi) HandleGetUserProfileByID(userId string) (*model.UserResponeMessage, error) {
	id, err := strconv.Atoi(userId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	userResponeMessage, err := s.UserService.GetUserProfileById(context.Background(), &model.UserAccount{Id: int32(id)})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return userResponeMessage, nil
}

func (s UserUsecaseApi) HandleUpdateUserProfilePicture(userId, fileName string) (*model.UserResponeMessage, error) {
	convertIdToInt, err := strconv.Atoi(userId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	userResponeMessage, err := s.UserService.UpdateUserProfilePicture(context.Background(), &model.UserProfile{Id: int32(convertIdToInt), ImageURL: fileName})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return userResponeMessage, nil
}

func (s UserUsecaseApi) HandleUpdateUserProfileByID(userId string, UserProfile *model.UserProfile) (*model.UserResponeMessage, error) {
	Id, err := strconv.Atoi(userId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	UserProfile.Id = int32(Id)
	UserResponeMessage, err := s.UserService.UpdateUserProfileById(context.Background(), UserProfile)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return UserResponeMessage, nil
}

func (s UserUsecaseApi) HandleUpdateUserLocation(userId string, userProfile *model.UserProfile) (*model.UserResponeMessage, error) {
	Id, err := strconv.Atoi(userId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	userProfile.Id = int32(Id)
	UserResponeMessage, err := s.UserService.UpdateUserLocation(context.Background(), userProfile)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return UserResponeMessage, nil
}

func (s UserUsecaseApi) HandleDeleteUserByID(userId string) (*model.UserResponeMessage, error) {
	Id, err := strconv.Atoi(userId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	userResponeMessage, err := s.UserService.DeleteUserByID(context.Background(), &model.UserAccount{Id: int32(Id)})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return userResponeMessage, nil
}

func (s UserUsecaseApi) HandleGetAllUserSummary(query *model.UserPagination) (*model.UserResponeMessage, error) {
	page, err := strconv.Atoi(query.Page)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	limit, err := strconv.Atoi(query.Limit)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	query.Page = strconv.Itoa((page * limit) - limit)

	result, err := s.UserService.GetAllUserSummary(context.Background(), query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return result, nil
}
