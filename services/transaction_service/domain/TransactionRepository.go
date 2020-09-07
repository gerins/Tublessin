package domain

import (
	"database/sql"
	"strconv"
	"tublessin/common/model"
)

type TransactionRepository struct {
	db *sql.DB
}

type TransactionRepositoryInterface interface {
	GetAllTransactionHistory(trans *model.TransactionHistory) (*model.ListTransactionHistory, error)
	PostNewTransaction(trans *model.TransactionHistory) (*model.TransactionHistory, error)
	UpdateTransactionByID(trans *model.TransactionHistory) (*model.TransactionHistory, error)
}

func NewTransactionRepository(db *sql.DB) TransactionRepositoryInterface {
	return &TransactionRepository{db}
}

func (t TransactionRepository) GetAllTransactionHistory(trans *model.TransactionHistory) (*model.ListTransactionHistory, error) {
	var listTransaction model.ListTransactionHistory

	result, err := t.db.Query(`SELECT * FROM transaction_history_view  WHERE id_montir = ? AND id_user = ?  ORDER BY date_created DESC`, trans.IdMontir, trans.IdUser)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		var t model.TransactionHistory
		var location model.TransactionLocation

		err := result.Scan(&t.Id, &t.IdMontir, &t.IdUser, &t.MontirFirstname, &t.UserFirstname, &t.Status, &t.DateCreated, &location.Latitude, &location.Longitude)
		if err != nil {
			return nil, err
		}
		t.Location = &location

		listTransaction.Results = append(listTransaction.Results, &t)
	}

	return &listTransaction, nil
}

func (t TransactionRepository) PostNewTransaction(trans *model.TransactionHistory) (*model.TransactionHistory, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	stmnt1, _ := tx.Prepare("INSERT INTO transaction_history(id_montir,id_user,montir_firstname,user_firstname) VALUE (?,?,?,?)")
	result, err := stmnt1.Exec(trans.IdMontir, trans.IdUser, trans.MontirFirstname, trans.UserFirstname)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	lastInsertID, _ := result.LastInsertId()

	stmnt2, _ := tx.Prepare("INSERT INTO transaction_location(transaction_history_id,latitude,longitude) VALUE (?,?,?)")
	_, err = stmnt2.Exec(lastInsertID, trans.Location.Latitude, trans.Location.Longitude)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	trans.Id = strconv.Itoa(int(lastInsertID))
	return trans, nil
}

func (t TransactionRepository) UpdateTransactionByID(trans *model.TransactionHistory) (*model.TransactionHistory, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return nil, err
	}

	stmnt1, _ := tx.Prepare("UPDATE transaction_history SET master_status_transaction_id = ? WHERE id = ?")
	_, err = stmnt1.Exec(trans.Status, trans.Id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return trans, nil
}
