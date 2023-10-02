package user

import (
	"backend/internal/db"
	"backend/internal/db/models"
	"database/sql"
)

func CreateUser(u models.User) (models.User, error) {
	database := db.GetDB()
	_, err := database.Exec(
		"INSERT INTO Users (first_name, middle_name, last_name, email, type_of_user, created_at, last_login)"+
			" VALUES ($1, $2, $3, $4, $5, $6, $7)",
		u.FirstName,
		u.MiddleName,
		u.LastName,
		u.Email,
		u.TypeOfUser,
		u.CreatedAt,
		u.LastLogin,
		u.Deleted,
	)
	if err != nil {
		return u, err
	}
	return u, nil
}

func ReadUserWithId(id int) (models.User, error) {
	database := db.GetDB()
	u, err := fillUserHelper(database.QueryRow("SELECT * FROM Users WHERE id = $1", id))
	return u, err
}

func ReadUserWithEmail(email string) (models.User, error) {
	database := db.GetDB()
	u, err := fillUserHelper(database.QueryRow("SELECT * FROM Users WHERE email = $1", email))
	return u, err
}

func UpdateUserLastLogin(u models.User) (models.User, error) {
	database := db.GetDB()
	_, err := database.Exec("UPDATE Users SET last_login = $1 WHERE id = $2", u.LastLogin, u.ID)
	if err != nil {
		return u, err
	}
	return u, nil
}

func DeleteUser(u models.User) (models.User, error) {
	database := db.GetDB()
	_, err := database.Exec("UPDATE Users SET deleted = $1 WHERE id = $2", true, u.ID)
	if err != nil {
		return u, err
	}
	return u, nil
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
