package domain

import (
	"context"
	"database/sql"
	"encoding/json"
	"sort"
	"strconv"
	"time"
	"tublessin/common/model"
	"tublessin/services/montir_service/utils"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"

	"golang.org/x/crypto/bcrypt"
)

type MontirUsecase struct {
	MontirRepository MontirRepositoryInterface
	RedisDatabase    *redis.Client
}

type MontirUsecaseInterface interface {
	Login(montirAccount *model.MontirAccount) (*model.MontirAccount, error)
	RegisterNewMontir(montirAccount *model.MontirAccount) (*model.MontirResponeMessage, error)
	GetMontirProfileByID(montirId int32) (*model.MontirResponeMessage, error)
	UpdateMontirProfilePicture(montirProfile *model.MontirProfile) (*model.MontirResponeMessage, error)
	UpdateMontirProfileByID(montirProfile *model.MontirProfile) (*model.MontirResponeMessage, error)
	UpdateMontirStatusByID(montirProfile *model.MontirProfile) (*model.MontirResponeMessage, error)
	UpdateMontirLocation(montirProfile *model.MontirProfile) (*model.MontirResponeMessage, error)
	GetAllActiveMontirWithLocation(userLocation *model.RequestActiveMontir) (*model.ListActiveMontirWithLocation, error)
	DeleteMontirByID(montirAccount *model.MontirAccount) (*model.MontirResponeMessage, error)
	GetAllMontirSummary(query *model.MontirPagination) (*model.ListActiveMontirWithLocation, error)
	InsertNewMontirRating(montirProfile *model.MontirProfile) (*model.MontirResponeMessage, error)
}

func NewMontirUsecase(db *sql.DB, rdb *redis.Client) MontirUsecaseInterface {
	return &MontirUsecase{MontirRepository: NewMontirRepository(db), RedisDatabase: rdb}
}

// Ini Adalah Layer Service dari Montir-Service, untuk menangani bussiness logic
func (s MontirUsecase) Login(montirAccount *model.MontirAccount) (*model.MontirAccount, error) {
	montirDetail, err := s.MontirRepository.Login(montirAccount.Username, "A", "S")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return montirDetail, nil
}

func (s MontirUsecase) RegisterNewMontir(montirAccount *model.MontirAccount) (*model.MontirResponeMessage, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(montirAccount.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	montirAccount.Password = string(hash)

	result, err := s.MontirRepository.RegisterNewMontir(montirAccount)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (s MontirUsecase) GetMontirProfileByID(montirId int32) (*model.MontirResponeMessage, error) {
	value, err := s.RedisDatabase.Get(context.Background(), strconv.Itoa(int(montirId))).Result()
	if err == nil {
		var montirRespone model.MontirResponeMessage
		json.Unmarshal([]byte(value), &montirRespone)
		if err != nil {
			log.Println("Something wrong when Unmarshal data to Montir Profile", err)
		}
		return &montirRespone, nil
	} else if err != nil && err != redis.Nil {
		log.Println("Something wrong when read data from Redis", err)
	}

	montirResponeMessage, err := s.MontirRepository.GetMontirProfileByID(montirId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result, err := json.Marshal(montirResponeMessage)
	err = s.RedisDatabase.Set(context.Background(), strconv.Itoa(int(montirId)), result, 30*time.Second).Err()
	if err != nil {
		log.Println("Cannot save montir profile to Redis", err)
	}

	return montirResponeMessage, nil
}

func (s MontirUsecase) UpdateMontirProfilePicture(montirProfile *model.MontirProfile) (*model.MontirResponeMessage, error) {
	montirResponeMessage, err := s.MontirRepository.UpdateMontirProfilePicture(montirProfile)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = s.RedisDatabase.Set(context.Background(), strconv.Itoa(int(montirProfile.Id)), "", 1*time.Second).Err()
	if err != nil {
		log.Println("Cannot Remove Montir profile From Redis", err)
	}

	return montirResponeMessage, nil
}

func (s MontirUsecase) UpdateMontirProfileByID(montirProfile *model.MontirProfile) (*model.MontirResponeMessage, error) {
	if montirProfile.VerifiedAccount != "" {
		result, err := s.MontirRepository.VerifiedMontirAccountByID(montirProfile)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		return result, nil
	} else {
		result, err := s.MontirRepository.UpdateMontirProfileByID(montirProfile)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		err = s.RedisDatabase.Set(context.Background(), strconv.Itoa(int(montirProfile.Id)), "", 1*time.Second).Err()
		if err != nil {
			log.Println("Cannot Remove Montir profile From Redis", err)
		}

		return result, nil
	}
}

func (s MontirUsecase) UpdateMontirStatusByID(montirProfile *model.MontirProfile) (*model.MontirResponeMessage, error) {
	montirResponeMessage, err := s.MontirRepository.UpdateMontirStatusByID(montirProfile)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = s.RedisDatabase.Set(context.Background(), strconv.Itoa(int(montirProfile.Id)), "", 1*time.Second).Err()
	if err != nil {
		log.Println("Cannot Remove Montir profile From Redis", err)
	}

	return montirResponeMessage, nil
}

func (s MontirUsecase) UpdateMontirLocation(montirProfile *model.MontirProfile) (*model.MontirResponeMessage, error) {
	montirResponeMessage, err := s.MontirRepository.UpdateMontirLocation(montirProfile)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return montirResponeMessage, nil
}

func (c MontirUsecase) GetAllActiveMontirWithLocation(userLocation *model.RequestActiveMontir) (*model.ListActiveMontirWithLocation, error) {
	result, err := c.MontirRepository.GetAllActiveMontirWithLocation("A")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, value := range result {
		montirLatitude := value.Location.Latitude
		montirLongitude := value.Location.Longitude
		value.Distance = int32(utils.CalculateDistance(userLocation.Latitude, userLocation.Longitude, montirLatitude, montirLongitude))
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Distance < result[j].Distance
	})

	if len(result) >= 6 {
		result = result[:6]
	}

	return &model.ListActiveMontirWithLocation{Response: "Search Nearby Montir Success", Code: "200", TotalMontir: strconv.Itoa(len(result)), List: result}, nil
}

func (s MontirUsecase) DeleteMontirByID(montirAccount *model.MontirAccount) (*model.MontirResponeMessage, error) {
	montirResponeMessage, err := s.MontirRepository.DeleteMontirByID(montirAccount)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return montirResponeMessage, nil
}

func (s MontirUsecase) GetAllMontirSummary(query *model.MontirPagination) (*model.ListActiveMontirWithLocation, error) {
	result, countItem, err := s.MontirRepository.GetAllMontirSummary(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &model.ListActiveMontirWithLocation{Response: "Success", Code: "200", TotalMontir: strconv.Itoa(countItem), List: result}, nil
}

func (s MontirUsecase) InsertNewMontirRating(montirProfile *model.MontirProfile) (*model.MontirResponeMessage, error) {
	montirResponeMessage, err := s.MontirRepository.InsertNewMontirRating(montirProfile)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return montirResponeMessage, nil
}
