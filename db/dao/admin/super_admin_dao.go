package dao

import (
	"backend/db"
	"database/sql"
)

func GetSuperAdminByEmail(email string) (*sql.Row, error) {
	db := db.GetDB()

	row := db.QueryRow(
		"SELECT * FROM SuperAdmins WHERE email = $1",
		email,
	)

	if row.Err() != nil {
		return nil, row.Err()
	}
	return row, nil
}
