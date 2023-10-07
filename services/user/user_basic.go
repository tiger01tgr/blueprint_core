package user

import (
	dao "backend/db/dao/users"
	models "backend/db/models"
	"database/sql"
)

// will support pagination: should go by id?
func GetUsers() {

}

func GetUserWithId(id int) (*models.User, error) {
	row, err := dao.ReadUserWithId(id)
	if err != nil {
		return nil, err
	}
	u, err := fillUserHelper(row)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func GetUserWithEmail(email string) (*models.User, error) {
	row, err := dao.ReadUserWithEmail(email)
	if err != nil {
		return nil, err
	}
	u, err := fillUserHelper(row)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func CreateUser(firstname string, middlename string, lastname string, email string, userType string, ) error {
	_, err := dao.CreateUser(firstname, middlename, lastname, email, userType)
	if err != nil {
		return err
	}
	return nil
}

func fillUserHelper(row *sql.Row) (models.User, error) {
	var u models.User
	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.MiddleName,
		&u.LastName,
		&u.Email,
		&u.UserType,
		&u.CreatedAt,
		&u.LastLogin,
		&u.Deleted,
	)
	if err != nil {
		return u, err
	}
	return u, nil
}
