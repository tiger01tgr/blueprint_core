package dao

import (
	"backend/db"
	"database/sql"
)

func GetIndustries() (*sql.Rows, error) {
	db := db.GetDB()
	rows, err := db.Query("SELECT * FROM Industries")
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func CreateIndustry(name string) (*sql.Result, error) {
	db := db.GetDB()
	res, err := db.Exec(
		"INSERT INTO Industries (name) VALUES ($1)",
		name,
	)
	return &res, err
}

func UpdateIndustry(id int, name string) error {
	db := db.GetDB()
	_, err := db.Exec("UPDATE Industries SET name = $1 WHERE id = $2", name, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteIndustry(id int) error {
	db := db.GetDB()
	_, err := db.Exec("DELETE FROM Industries WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
