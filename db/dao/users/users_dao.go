package dao

import (
	"backend/db"
	"database/sql"
	"time"
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

func ReadUserWithId(id int) (*sql.Row, error) {
	database := db.GetDB()
	row := database.QueryRow("SELECT * FROM Users WHERE id = $1", id)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return row, nil
}

func ReadUserWithEmail(email string) (*sql.Row, error) {
	database := db.GetDB()
	row := database.QueryRow("SELECT * FROM Users WHERE email = $1", email)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return row, nil
}

func UpdateUserLastLogin(id int) (*sql.Result, error) {
	database := db.GetDB()
	res, err := database.Exec("UPDATE Users SET last_login = $1 WHERE id = $2", time.Now(), id)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

func DeleteUser(id int) (*sql.Result, error) {
	database := db.GetDB()
	res, err := database.Exec("UPDATE Users SET deleted = $1 WHERE id = $2", true, id)
	if err != nil {
		return &res, err
	}
	return &res, nil
}
