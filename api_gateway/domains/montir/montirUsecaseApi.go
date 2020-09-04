package montir

import (
	"context"
	"log"
	"strconv"
	"tublessin/common/model"
)

type MontirUsecaseApi struct {
	MontirService model.MontirClient
}

type MontirUsecaseApiInterface interface {
	HandleGetMontirProfileByID(montirId string) (*model.MontirResponeMessage, error)
	HandleUpdateMontirProfilePicture(MontirId, fileName string) (*model.MontirResponeMessage, error)
	HandleUpdateMontirProfileByID(montirId string, profile *model.MontirProfile) (*model.MontirResponeMessage, error)
	HandleUpdateMontirLocation(montirId string, montirProfile *model.MontirProfile) (*model.MontirResponeMessage, error)
	HandleGetAllActiveMontirWithLocation(latitude, longitude string) (*model.ListActiveMontirWithLocation, error)
}

func NewMontirUsecaseApi(montirService model.MontirClient) MontirUsecaseApiInterface {
	return MontirUsecaseApi{MontirService: montirService}
}

func (s MontirUsecaseApi) HandleGetMontirProfileByID(montirId string) (*model.MontirResponeMessage, error) {
	id, err := strconv.Atoi(montirId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	montirAccountWithId := &model.MontirAccount{Id: int32(id)}
	montirResponeMessage, err := s.MontirService.GetMontirProfileByID(context.Background(), montirAccountWithId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return montirResponeMessage, nil
}

func (s MontirUsecaseApi) HandleUpdateMontirProfilePicture(MontirId, fileName string) (*model.MontirResponeMessage, error) {
	convertIdToInt, err := strconv.Atoi(MontirId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	MontirResponeMessage, err := s.MontirService.UpdateMontirProfilePicture(context.Background(), &model.MontirProfile{Id: int32(convertIdToInt), ImageURL: fileName})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return MontirResponeMessage, nil
}

func (s MontirUsecaseApi) HandleUpdateMontirProfileByID(montirId string, profile *model.MontirProfile) (*model.MontirResponeMessage, error) {
	Id, err := strconv.Atoi(montirId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	profile.Id = int32(Id)
	MontirResponeMessage, err := s.MontirService.UpdateMontirProfileByID(context.Background(), profile)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return MontirResponeMessage, nil
}

func (s MontirUsecaseApi) HandleUpdateMontirLocation(montirId string, montirProfile *model.MontirProfile) (*model.MontirResponeMessage, error) {
	Id, err := strconv.Atoi(montirId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	montirProfile.Id = int32(Id)
	MontirResponeMessage, err := s.MontirService.UpdateMontirLocation(context.Background(), montirProfile)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return MontirResponeMessage, nil
}

func (s MontirUsecaseApi) HandleGetAllActiveMontirWithLocation(latitude, longitude string) (*model.ListActiveMontirWithLocation, error) {
	var userLocation model.RequestActiveMontir

	doubleLatitude, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	doubleLongitude, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	userLocation.Latitude = doubleLatitude
	userLocation.Longitude = doubleLongitude

	listActiveMontir, err := s.MontirService.GetAllActiveMontirWithLocation(context.Background(), &userLocation)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return listActiveMontir, nil
}
