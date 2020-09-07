package domain

import (
	"context"
	"database/sql"
	"tublessin/common/model"
)

type TransactionServer struct {
	TransactionUsecase TransactionUsecaseInterface
}

func NewTransactionController(db *sql.DB) *TransactionServer {
	return &TransactionServer{NewTransactionUsecase(db)}
}

func (c TransactionServer) PostNewTransaction(ctx context.Context, param *model.TransactionHistory) (*model.TransactionHistory, error) {
	result, err := c.TransactionUsecase.PostNewTransaction(param)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c TransactionServer) UpdateTransactionByID(ctx context.Context, param *model.TransactionHistory) (*model.TransactionHistory, error) {
	result, err := c.TransactionUsecase.UpdateTransactionByID(param)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c TransactionServer) GetAllTransactionHistory(ctx context.Context, param *model.TransactionHistory) (*model.ListTransactionHistory, error) {
	result, err := c.TransactionUsecase.GetAllTransactionHistory(param)
	if err != nil {
		return nil, err
	}

	return result, nil
}
