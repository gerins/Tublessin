package domain

import (
	"database/sql"
	"errors"
	"tublessin/common/model"
)

type MontirUsecase struct {
	MontirRepository MontirRepositoryInterface
}

type MontirUsecaseInterface interface {
	Login(montirAccount *model.MontirAccount) (*model.MontirAccount, error)
	RegisterNewMontir(montirAccount *model.MontirAccount) (*model.MontirResponeMessage, error)
	GetMontirProfileByID(montirId string) (*model.MontirResponeMessage, error)
}

func NewMontirUsecase(db *sql.DB) MontirUsecaseInterface {
	return &MontirUsecase{NewMontirRepository(db)}
}

// Ini Adalah Layer Service dari Montir-Service, untuk menangani bussiness logic
func (s MontirUsecase) Login(montirAccount *model.MontirAccount) (*model.MontirAccount, error) {
	montirDetail, err := s.MontirRepository.Login(montirAccount.Username, "A")
	if err != nil {
		return nil, err
	}

	return montirDetail, nil
}

func (s MontirUsecase) RegisterNewMontir(montirAccount *model.MontirAccount) (*model.MontirResponeMessage, error) {
	if montirAccount == nil || montirAccount.Profile == nil {
		return nil, errors.New("Body Cannot Empty")
	}

	result, err := s.MontirRepository.RegisterNewMontir(montirAccount)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s MontirUsecase) GetMontirProfileByID(montirId string) (*model.MontirResponeMessage, error) {
	montirResponeMessage, err := s.MontirRepository.GetMontirProfileByID(montirId, "A")
	if err != nil {
		return nil, err
	}
	return montirResponeMessage, nil
}
