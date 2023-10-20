package dao

import (
	"backend/db"
	"database/sql"
	"fmt"
)

func CustomQueryForRow(query string, args ...any) (*sql.Row, error) {
	db := db.GetDB()
	var row *sql.Row
	if args == nil {
		fmt.Println("args is nil")
		row = db.QueryRow(query)
	} else {
		fmt.Println("args is not nil")
		row = db.QueryRow(query, args...)
	}
	if row.Err() != nil {
		return nil, row.Err()
	}
	return row, nil
}

func CustomQueryForRows(query string) (*sql.Rows, error) {
	db := db.GetDB()
	fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
