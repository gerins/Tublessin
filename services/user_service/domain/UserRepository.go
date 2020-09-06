package domain

import (
	"database/sql"
	"tublessin/common/model"
)

type UserRepository struct {
	db *sql.DB
}

type UserRepositoryInterface interface {
	Login(username, status string) (*model.UserAccount, error)
	RegisterNewUser(m *model.UserAccount) (*model.UserResponeMessage, error)
	GetUserProfileById(userId int32) (*model.UserResponeMessage, error)
	UpdateUserProfilePicture(userProfile *model.UserProfile) (*model.UserResponeMessage, error)
	UpdateUserProfileByID(mp *model.UserProfile) (*model.UserResponeMessage, error)
	UpdateUserLocation(mp *model.UserProfile) (*model.UserResponeMessage, error)
	DeleteUserByID(userAccount *model.UserAccount) (*model.UserResponeMessage, error)
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{db}
}

// Disini adalah layer Repository dari User-Service, untuk berkomunikasi dengan database
func (r UserRepository) Login(username, status string) (*model.UserAccount, error) {
	results := r.db.QueryRow("SELECT * FROM user_account WHERE username=? AND status_account=?", username, status)
	var userAccount model.UserAccount

	err := results.Scan(&userAccount.Id, &userAccount.Username, &userAccount.Password, &userAccount.StatusAccount)
	if err != nil {
		return nil, err
	}

	return &userAccount, nil
}

func (r UserRepository) RegisterNewUser(m *model.UserAccount) (*model.UserResponeMessage, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	stmnt1, _ := tx.Prepare("INSERT INTO user_account(username, password) VALUE (?,?)")
	stmnt2, _ := tx.Prepare("INSERT INTO user_profile(user_account_id, firstname, lastname, gender, phone_number, email) VALUE (?,?,?,?,?,?)")
	stmnt3, _ := tx.Prepare("INSERT INTO user_location(user_account_id) VALUE(?)")

	result, err := stmnt1.Exec(m.Username, m.Password)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	lastInsertID, _ := result.LastInsertId()
	_, err = stmnt2.Exec(lastInsertID, m.Profile.Firstname, m.Profile.Lastname, m.Profile.Gender, m.Profile.PhoneNumber, m.Profile.Email)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	_, err = stmnt3.Exec(lastInsertID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &model.UserResponeMessage{Response: "Inserting New User Success", Code: "200", Result: m}, nil
}

func (r UserRepository) GetUserProfileById(userId int32) (*model.UserResponeMessage, error) {
	var userAccount model.UserAccount

	result := r.db.QueryRow("SELECT * FROM user_account WHERE id=?", userId)
	err := result.Scan(&userAccount.Id, &userAccount.Username, &userAccount.Password, &userAccount.StatusAccount)
	if err != nil {
		return nil, err
	}

	var mp model.UserProfile
	result2 := r.db.QueryRow("SELECT * FROM user_profile WHERE user_account_id=?", userId)
	err = result2.Scan(&mp.Id, &mp.Firstname, &mp.Lastname, &mp.Gender, &mp.PhoneNumber, &mp.Email, &mp.ImageURL, &mp.DateUpdated, &mp.DateCreated)
	if err != nil {
		return nil, err
	}
	userAccount.Profile = &mp

	var ml model.UserLocation
	result3 := r.db.QueryRow("SELECT * FROM user_location WHERE user_account_id=? ", userId)
	err = result3.Scan(&mp.Id, &ml.Latitude, &ml.Longitude, &ml.DateUpdated)
	if err != nil {
		return nil, err
	}
	userAccount.Profile.Location = &ml

	return &model.UserResponeMessage{Response: "Get user Profile Success", Code: "200", Result: &userAccount}, nil
}

func (r UserRepository) UpdateUserProfilePicture(userProfile *model.UserProfile) (*model.UserResponeMessage, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	stmnt1, _ := tx.Prepare("UPDATE user_profile SET imageURL = ? WHERE user_account_id = ?")
	_, err = stmnt1.Exec(userProfile.ImageURL, userProfile.Id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &model.UserResponeMessage{Response: "Updating Profile Picture Success", Code: "200", Result: &model.UserAccount{
		Profile: &model.UserProfile{Id: userProfile.Id, ImageURL: userProfile.ImageURL},
	}}, nil
}

func (r UserRepository) UpdateUserProfileByID(mp *model.UserProfile) (*model.UserResponeMessage, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	stmnt1, _ := tx.Prepare("UPDATE user_profile SET firstname=?,lastname=?,gender=?,phone_number=?,email=? WHERE user_account_id = ?")
	_, err = stmnt1.Exec(mp.Firstname, mp.Lastname, mp.Gender, mp.PhoneNumber, mp.Email, mp.Id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &model.UserResponeMessage{Response: "Updating User Profile Success", Code: "200", Result: &model.UserAccount{
		Profile: mp,
	}}, nil
}

func (r UserRepository) UpdateUserLocation(mp *model.UserProfile) (*model.UserResponeMessage, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	stmnt1, _ := tx.Prepare("UPDATE user_location SET latitude=?,longitude=? WHERE user_account_id = ?")
	_, err = stmnt1.Exec(mp.Location.Latitude, mp.Location.Longitude, mp.Id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &model.UserResponeMessage{Response: "Updating User Location Success", Code: "200", Result: &model.UserAccount{
		Profile: mp,
	}}, nil
}

func (r UserRepository) DeleteUserByID(userAccount *model.UserAccount) (*model.UserResponeMessage, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	stmnt1, _ := tx.Prepare("UPDATE user_account SET status_account = ? WHERE id = ?")
	_, err = stmnt1.Exec("N", userAccount.Id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &model.UserResponeMessage{Response: "Deactivated User Success", Code: "200"}, nil
}
