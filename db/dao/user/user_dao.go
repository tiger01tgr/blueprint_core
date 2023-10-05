package dao

import (
	"backend/db"
	"backend/db/models"
	"database/sql"
)

func CreateUser(u models.User) (*sql.Result, error) {
	database := db.GetDB()
	res, err := database.Exec(
		"INSERT INTO Users (first_name, middle_name, last_name, email, type_of_user, created_at, last_login, deleted)"+
			" VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		u.FirstName,
		u.MiddleName,
		u.LastName,
		u.Email,
		u.TypeOfUser,
		u.CreatedAt,
		u.LastLogin,
		u.Deleted,
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

func UpdateUserLastLogin(u models.User) (*sql.Result, error) {
	database := db.GetDB()
	res, err := database.Exec("UPDATE Users SET last_login = $1 WHERE id = $2", u.LastLogin, u.ID)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

func DeleteUser(u models.User) (*sql.Result, error) {
	database := db.GetDB()
	res, err := database.Exec("UPDATE Users SET deleted = $1 WHERE id = $2", true, u.ID)
	if err != nil {
		return &res, err
	}
	return &res, nil
}
