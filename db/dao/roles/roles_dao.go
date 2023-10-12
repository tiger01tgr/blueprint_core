package dao

import (
	"backend/db"
	"database/sql"
)

func GetRoles() (*sql.Rows, error) {
	db := db.GetDB()
	rows, err := db.Query("SELECT * FROM Roles")
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func CreateRole(name string) (*sql.Result, error) {
	db := db.GetDB()
	res, err := db.Exec(
		"INSERT INTO Roles (name) VALUES ($1)",
		name,
	)
	return &res, err
}

func UpdateRole(id int, name string) error {
	db := db.GetDB()
	_, err := db.Exec("UPDATE Roles SET name = $1 WHERE id = $2", name, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRole(id int) error {
	db := db.GetDB()
	_, err := db.Exec("DELETE FROM Roles WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
