package domain

import (
	"database/sql"
	"tublessin/common/model"

	log "github.com/sirupsen/logrus"
)

type TransactionUsecase struct {
	TransactionRepository TransactionRepositoryInterface
}

type TransactionUsecaseInterface interface {
	GetAllTransactionHistory(trans *model.TransactionHistory) (*model.ListTransactionHistory, error)
	PostNewTransaction(trans *model.TransactionHistory) (*model.TransactionHistory, error)
	UpdateTransactionByID(trans *model.TransactionHistory) (*model.TransactionHistory, error)
}

func NewTransactionUsecase(db *sql.DB) TransactionUsecaseInterface {
	return &TransactionUsecase{NewTransactionRepository(db)}
}

func (t TransactionUsecase) GetAllTransactionHistory(trans *model.TransactionHistory) (*model.ListTransactionHistory, error) {
	if trans.IdMontir == "" && trans.IdUser != "" {
		result, err := t.TransactionRepository.GetAllTransactionHistoryUser(trans)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return result, nil
	}

	if trans.IdMontir != "" && trans.IdUser == "" {
		result, err := t.TransactionRepository.GetAllTransactionHistoryMontir(trans)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return result, nil
	}

	result, err := t.TransactionRepository.GetAllTransactionHistory(trans)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

func (t TransactionUsecase) PostNewTransaction(trans *model.TransactionHistory) (*model.TransactionHistory, error) {
	result, err := t.TransactionRepository.PostNewTransaction(trans)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

func (t TransactionUsecase) UpdateTransactionByID(trans *model.TransactionHistory) (*model.TransactionHistory, error) {
	result, err := t.TransactionRepository.UpdateTransactionByID(trans)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}
