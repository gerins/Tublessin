package domain

import (
	"database/sql"
	"errors"
	"sort"
	"strconv"
	"tublessin/common/model"
	"tublessin/services/montir_service/utils"

	"golang.org/x/crypto/bcrypt"
)

type MontirUsecase struct {
	MontirRepository MontirRepositoryInterface
}

type MontirUsecaseInterface interface {
	Login(montirAccount *model.MontirAccount) (*model.MontirAccount, error)
	RegisterNewMontir(montirAccount *model.MontirAccount) (*model.MontirResponeMessage, error)
	GetMontirProfileByID(montirId string) (*model.MontirResponeMessage, error)
	UpdateMontirProfilePicture(montirProfile *model.MontirProfile) (*model.MontirResponeMessage, error)
	UpdateMontirProfileByID(montirProfile *model.MontirProfile) (*model.MontirResponeMessage, error)
	UpdateMontirLocation(montirProfile *model.MontirProfile) (*model.MontirResponeMessage, error)
	GetAllActiveMontirWithLocation(userLocation *model.RequestActiveMontir) (*model.ListActiveMontirWithLocation, error)
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

	hash, err := bcrypt.GenerateFromPassword([]byte(montirAccount.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	montirAccount.Password = string(hash)

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

func (s MontirUsecase) UpdateMontirProfilePicture(montirProfile *model.MontirProfile) (*model.MontirResponeMessage, error) {
	montirResponeMessage, err := s.MontirRepository.UpdateMontirProfilePicture(montirProfile)
	if err != nil {
		return nil, err
	}
	return montirResponeMessage, nil
}

func (s MontirUsecase) UpdateMontirProfileByID(montirProfile *model.MontirProfile) (*model.MontirResponeMessage, error) {
	montirResponeMessage, err := s.MontirRepository.UpdateMontirProfileByID(montirProfile)
	if err != nil {
		return nil, err
	}
	return montirResponeMessage, nil
}

func (s MontirUsecase) UpdateMontirLocation(montirProfile *model.MontirProfile) (*model.MontirResponeMessage, error) {
	if montirProfile.Location == nil || montirProfile.Location.Latitude == 0 || montirProfile.Location.Longitude == 0 {
		return nil, errors.New("Body cannot empty")
	}

	montirResponeMessage, err := s.MontirRepository.UpdateMontirLocation(montirProfile)
	if err != nil {
		return nil, err
	}
	return montirResponeMessage, nil
}

func (c MontirUsecase) GetAllActiveMontirWithLocation(userLocation *model.RequestActiveMontir) (*model.ListActiveMontirWithLocation, error) {
	mockLatitude := -6.174277
	mockLongitude := 106.829723

	result, err := c.MontirRepository.GetAllActiveMontirWithLocation("A")
	if err != nil {
		return nil, err
	}

	for _, value := range result {
		montirLatitude := value.Location.Latitude
		montirLongitude := value.Location.Longitude
		value.Distance = int32(utils.CalculateDistance(mockLatitude, mockLongitude, montirLatitude, montirLongitude))
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Distance < result[j].Distance
	})

	return &model.ListActiveMontirWithLocation{Response: "Search Nearby Montir Success", Code: "200", TotalMontir: strconv.Itoa(len(result)), List: result}, nil
}
