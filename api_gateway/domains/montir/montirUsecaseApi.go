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
}

func NewMontirUsecaseApi(montirService model.MontirClient) MontirUsecaseApiInterface {
	return MontirUsecaseApi{MontirService: montirService}
}

func (s MontirUsecaseApi) HandleGetMontirProfileByID(montirId string) (*model.MontirResponeMessage, error) {
	id, _ := strconv.Atoi(montirId)
	montirAccountWithId := &model.MontirAccount{Id: int32(id)}

	montirResponeMessage, err := s.MontirService.GetMontirProfileByID(context.Background(), montirAccountWithId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return montirResponeMessage, nil
}
