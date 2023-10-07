package dao

import (
	"backend/db"
	"database/sql"
)

// CreateQuestion inserts a new question into the database.
func CreateQuestion(questionSetID int, text string, timelimit int) (*sql.Result, error) {
	db := db.GetDB()
	res, err := db.Exec(
		"INSERT INTO Questions (questionSetId, text, timelimit, deleted) VALUES ($1, $2, $3, $4)",
		questionSetID,
		text,
		timelimit,
		false,
	)
	return &res, err
}

func GetQuestionByID(id int) (*sql.Row, error) {
	db := db.GetDB()
	row := db.QueryRow("SELECT id, questionSetId, text, timelimit, deleted FROM Questions WHERE id = $1", id)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return row, nil
}

// UpdateQuestion updates an existing question in the database.
func UpdateQuestion(id int, text string, timelimit int) error {
	db := db.GetDB()
	_, err := db.Exec("UPDATE Questions SET text = $1, timelimit = $2 WHERE id = $3", text, timelimit, id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteQuestion marks a question as deleted in the database.
func DeleteQuestion(id int) error {
	db := db.GetDB()
	// Prepare and execute the SQL UPDATE statement to set the "deleted" flag to true
	_, err := db.Exec("UPDATE Questions SET deleted = true WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
