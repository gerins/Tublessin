package domain

import (
	"database/sql"
	"errors"
	"log"
	"tublessin/common/model"
)

type MontirRepository struct {
	db *sql.DB
}

type MontirRepositoryInterface interface {
	Login(username, status string) (*model.MontirAccount, error)
	RegisterNewMontir(m *model.MontirAccount) (*model.MontirResponeMessage, error)
	GetMontirProfileByID(montirId, statusAccount string) (*model.MontirResponeMessage, error)
}

func NewMontirRepository(db *sql.DB) MontirRepositoryInterface {
	return &MontirRepository{db}
}

// Disini adalah layer Repository dari Montir-Service, untuk berkomunikasi dengan database
func (r MontirRepository) Login(username, status string) (*model.MontirAccount, error) {
	results := r.db.QueryRow("SELECT * FROM montir_account WHERE username=? AND status_account=?", username, status)
	var montirAccount model.MontirAccount

	err := results.Scan(&montirAccount.Id, &montirAccount.Username, &montirAccount.Password, &montirAccount.StatusAccount)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return &montirAccount, nil
}

func (r MontirRepository) RegisterNewMontir(m *model.MontirAccount) (*model.MontirResponeMessage, error) {
	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	stmnt1, _ := tx.Prepare("INSERT INTO montir_account(username, password) VALUE (?,?)")
	stmnt2, _ := tx.Prepare("INSERT INTO montir_profile(montir_account_id, firstname, lastname, gender, city, email, phone_number) VALUE (?,?,?,?,?,?,?)")
	stmnt3, _ := tx.Prepare("INSERT INTO montir_location(montir_account_id) VALUE(?)")
	stmnt4, _ := tx.Prepare("INSERT INTO montir_status(montir_account_id) VALUE(?)")

	result, err := stmnt1.Exec(m.Username, m.Password)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	lastInsertID, _ := result.LastInsertId()
	_, err = stmnt2.Exec(lastInsertID, m.Profile.Firstname, m.Profile.Lastname, m.Profile.Gender, m.Profile.City, m.Profile.Email, m.Profile.PhoneNumber)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	_, err = stmnt3.Exec(lastInsertID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	_, err = stmnt4.Exec(lastInsertID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &model.MontirResponeMessage{Response: "Inserting New Montir Success", Code: "200", Result: m}, nil
}

func (r MontirRepository) GetMontirProfileByID(montirId, statusAccount string) (*model.MontirResponeMessage, error) {
	var montirAccount model.MontirAccount

	result := r.db.QueryRow("SELECT * FROM montir_account WHERE id=? AND status_account=?", montirId, statusAccount)
	err := result.Scan(&montirAccount.Id, &montirAccount.Username, &montirAccount.Password, &montirAccount.StatusAccount)
	if err != nil {
		return nil, errors.New("Montir ID Not Found")
	}

	var mp model.MontirProfile
	result2 := r.db.QueryRow("SELECT * FROM montir_profile WHERE montir_account_id=?", montirId)
	err = result2.Scan(&mp.Id, &mp.Firstname, &mp.Lastname, &mp.BornDate, &mp.Gender, &mp.Ktp, &mp.Address, &mp.City, &mp.Email, &mp.PhoneNumber, &mp.ImageURL, &mp.VerifiedAccount, &mp.DateUpdated, &mp.DateCreated)
	if err != nil {
		log.Println(err)
	}
	montirAccount.Profile = &mp

	var ml model.MontirLocation
	result3 := r.db.QueryRow("SELECT * FROM montir_location WHERE montir_account_id=? ", montirId)
	err = result3.Scan(&mp.Id, &ml.Latitude, &ml.Longitude, &ml.DateUpdated)
	if err != nil {
		log.Println(err)
	}
	montirAccount.Profile.Location = &ml

	var ms model.MontirStatus
	result4 := r.db.QueryRow(`SELECT ms.status_operational, msa.status, ms.date_updated
	FROM montir_status ms JOIN master_status_activity msa ON ms.status_activity_id = msa.id
	WHERE ms.montir_account_id = ?`, montirId)
	err = result4.Scan(&ms.StatusOperational, &ms.StatusActivity, &ms.DateUpdated)
	if err != nil {
		log.Println(err)
	}
	montirAccount.Profile.Status = &ms

	result5, err := r.db.Query("SELECT * FROM montir_rating WHERE montir_account_id=?", montirId)
	if err != nil {
		log.Println(err)
	}

	for result5.Next() {
		var rating model.MontirRating
		err := result5.Scan(&mp.Id, &rating.Rating, &rating.RaterId, &rating.Review, &rating.DateCreated)
		if err != nil {
			log.Println(err)
			break
		}
		montirAccount.Profile.RatingList = append(montirAccount.Profile.RatingList, &rating)
	}

	return &model.MontirResponeMessage{Response: "Get Montir Profile Success", Code: "200", Result: &montirAccount}, nil
}
