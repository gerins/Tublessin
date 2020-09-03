package domain

import (
	"context"
	"database/sql"
	"strconv"
	"tublessin/common/model"
)

type UserServer struct {
	UserUsecase UserUsecaseInterface
}

func NewUserController(db *sql.DB) *UserServer {
	return &UserServer{NewUserUsecase(db)}
}

// Disini adalah pusat Method2 dari User-Service
func (c UserServer) Login(ctx context.Context, param *model.UserAccount) (*model.UserAccount, error) {
	result, err := c.UserUsecase.Login(param)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c UserServer) RegisterNewUser(ctx context.Context, param *model.UserAccount) (*model.UserResponeMessage, error) {
	userResponeMessage, err := c.UserUsecase.RegisterNewUser(param)
	if err != nil {
		return nil, err
	}

	return userResponeMessage, nil
}

func (c UserServer) GetUserProfileById(ctx context.Context, param *model.UserAccount) (*model.UserResponeMessage, error) {
	userId := strconv.Itoa(int(param.Id))
	userResponeMessage, err := c.UserUsecase.GetUserProfileById(userId)
	if err != nil {
		return nil, err
	}

	return userResponeMessage, nil
}

func (c UserServer) UpdateUserProfileById(ctx context.Context, param *model.UserProfile) (*model.UserResponeMessage, error) {
	return &model.UserResponeMessage{Response: ""}, nil

}

func (c UserServer) UpdateUserProfilePicture(ctx context.Context, param *model.UserProfile) (*model.UserResponeMessage, error) {
	userResponeMessage, err := c.UserUsecase.UpdateUserProfilePicture(param)
	if err != nil {
		return nil, err
	}

	return userResponeMessage, nil
}
