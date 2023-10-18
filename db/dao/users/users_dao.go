package dao

import (
	"backend/db"
	"database/sql"
	"time"
	"backend/db/models"
)

func CreateUser(firstName, middleName, lastName, email, userType string) (*sql.Result, error) {
	database := db.GetDB()
	res, err := database.Exec(
		"INSERT INTO Users (first_name, middle_name, last_name, email, type_of_user, created_at, last_login, deleted)"+
			" VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		firstName,
		middleName,
		lastName,
		email,
		userType,
		time.Now(),
		time.Now(),
		false,
	)
	return &res, err
}
/*
		&u.ID,
		&u.FirstName,
		&u.MiddleName,
		&u.LastName,
		&u.Email,
		&u.UserType,
		&u.CreatedAt,
		&u.LastLogin,
		&u.Deleted,
*/

func ReadUserWithId(id int64) (*sql.Row, error) {
	database := db.GetDB()
	row := database.QueryRow("SELECT id, first_name, middle_name, last_name, email, type_of_user, created_at, last_login, deleted FROM Users WHERE id = $1", id)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return row, nil
}

func ReadUserIdWithEmail(email string) (*sql.Row, error) {
	database := db.GetDB()
	row := database.QueryRow("SELECT id FROM Users WHERE email = $1", email)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return row, nil
}

func UpdateUserLastLogin(id int64) (*sql.Result, error) {
	database := db.GetDB()
	res, err := database.Exec("UPDATE Users SET last_login = $1 WHERE id = $2", time.Now(), id)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

func ReadUserProfile(id int64) (*sql.Row, error) {
	database := db.GetDB()
	row := database.QueryRow("SELECT id, first_name, middle_name, last_name, email, type_of_user, school, major, employer, position, phone, resume FROM Users WHERE id = $1", id)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return row, nil
}

func DeleteUser(id int64) (*sql.Result, error) {
	database := db.GetDB()
	res, err := database.Exec("UPDATE Users SET deleted = $1 WHERE id = $2", true, id)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

func UpdateUserProfile(user *models.User) (*sql.Result, error) {
	database := db.GetDB()
	res, err := database.Exec("UPDATE Users SET first_name = $1, middle_name = $2, last_name = $3, email = $4, type_of_user = $5, school = $6, major = $7, employer = $8, position = $9, phone = $10, resume = $11 WHERE id = $12", user.FirstName, user.MiddleName, user.LastName, user.Email, user.UserType, user.School, user.Major, user.Employer, user.Position, user.Phone, user.Resume, user.ID)
	if err != nil {
		return &res, err
	}
	return &res, nil
}
