package dao

import (
	"backend/db"
	"database/sql"
	"time"
)

func CreateEmployer(name, logo string, industryId int) (*sql.Result, error) {
	db := db.GetDB()
	res, err := db.Exec(
		"INSERT INTO Employers (name, logo, industryId, created_at, deleted) VALUES ($1, $2, $3, $4, $5)",
		name,
		logo,
		industryId,
		time.Now(),
		false,
	)
	return &res, err
}

func GetEmployerByID(id string) (*sql.Row, error) {
	db := db.GetDB()

	// SQL Join to get the industry name
	row := db.QueryRow(
		"SELECT" + 
		" Employers.id" + 
		", Employers.name" +
		", Employers.logo" +
		", Industries.name" +
		", Employers.industryId" +
		", Employers.created_at" +
		", Employers.deleted" +
		" FROM Employers" +
		" INNER JOIN Industries ON Employers.industryId = Industries.id" +
		" WHERE Employers.id = $1", 
		id,
	)

	if row.Err() != nil {
		return nil, row.Err()
	}
	return row, nil
}

func GetAllEmployers() (*sql.Rows, error) {
	db := db.GetDB()
	
	// Select ALL from Employers and join with Industries
	rows, err := db.Query(
		"SELECT" +
		" Employers.id" +
		", Employers.name" +
		", Employers.logo" +
		", Industries.name" +
		", Employers.industryId" + 
		", Employers.created_at" +
		", Employers.deleted" +
		" FROM Employers" +
		" INNER JOIN Industries ON Employers.industryId = Industries.id",
	)

	if err != nil {
		return nil, err
	}
	return rows, nil
}

func UpdateEmployer(id, name, logo, industry string) error {
	db := db.GetDB()
	_, err := db.Exec("UPDATE Employers SET name = $1, logo = $2, industryId = $3, WHERE id = $4", name, logo, industry, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEmployer(id string) error {
	db := db.GetDB()
	_, err := db.Exec("UPDATE Employers SET deleted = true WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}