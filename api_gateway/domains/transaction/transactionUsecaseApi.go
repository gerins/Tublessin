package transaction

import (
	"context"
	"errors"
	"log"
	"tublessin/common/model"
)

type TransactionUsecaseApi struct {
	TransactionService model.TransactionClient
}

type TransactionUsecaseApiInterface interface {
	HandleGetAllTransactionHistory(montirId, userId string) (*model.ListTransactionHistory, error)
	HandlePostNewTransaction(trans *model.TransactionHistory) (*model.TransactionHistory, error)
	HandleUpdateTransactionByID(trans *model.TransactionHistory) (*model.TransactionHistory, error)
}

func NewTransactionUsecaseApi(TransactionService model.TransactionClient) TransactionUsecaseApiInterface {
	return TransactionUsecaseApi{TransactionService: TransactionService}
}

func (s TransactionUsecaseApi) HandlePostNewTransaction(trans *model.TransactionHistory) (*model.TransactionHistory, error) {
	if trans.IdMontir == "" ||
		trans.IdUser == "" ||
		trans.MontirFirstname == "" ||
		trans.UserFirstname == "" ||
		trans.Location.Latitude == 0 ||
		trans.Location.Longitude == 0 {
		return nil, errors.New("Form Body Cannot empty")
	}

	result, err := s.TransactionService.PostNewTransaction(context.Background(), trans)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return result, nil
}

func (s TransactionUsecaseApi) HandleUpdateTransactionByID(trans *model.TransactionHistory) (*model.TransactionHistory, error) {
	if trans.Status == "" {
		return nil, errors.New("Form Body Cannot empty")
	}

	result, err := s.TransactionService.UpdateTransactionByID(context.Background(), trans)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return result, nil
}

func (s TransactionUsecaseApi) HandleGetAllTransactionHistory(montirId, userId string) (*model.ListTransactionHistory, error) {
	result, err := s.TransactionService.GetAllTransactionHistory(context.Background(), &model.TransactionHistory{IdMontir: montirId, IdUser: userId})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return result, nil
}
