package user

import (
	userDao "backend/internal/db/dao/user"
	models "backend/internal/db/models"
	"time"
	"database/sql"
)

// will support pagination: should go by id?
func GetUsers() {

}

func GetUserWithId(id int) (*models.User, error) {
	row, err := userDao.ReadUserWithId(id)
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
	row, err := userDao.ReadUserWithEmail(email)
	if err != nil {
		return nil, err
	}
	u, err := fillUserHelper(row)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func CreateUser(firstname string, middlename string, lastname string, email string, typeOfUser string, ) error {
	u := models.User{
		FirstName:  firstname,
		MiddleName: middlename,
		LastName:   lastname,
		Email:      email,
		TypeOfUser: typeOfUser,
		CreatedAt: time.Now(),
		LastLogin: time.Now(),
		Deleted: false,
	}
	_, err := userDao.CreateUser(u)
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
		&u.TypeOfUser,
		&u.CreatedAt,
		&u.LastLogin,
		&u.Deleted,
	)
	if err != nil {
		return u, err
	}
	return u, nil
}
