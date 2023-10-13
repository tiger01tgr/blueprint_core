package dao

import (
	"backend/db"
	"database/sql"
	"time"
)

// CreateQuestion inserts a new question into the database.
func CreateQuestionSet(name, interviewType string, employerId, roleId int64) (sql.Result, error) {
	db := db.GetDB()
	res, err := db.Exec(
		"INSERT INTO QuestionSets (name, interviewType, employerId, roleId, created_at, deleted) VALUES ($1, $2, $3, $4, $5, $6)",
		name,
		interviewType,
		employerId,
		roleId,
		time.Now(),
		false,
	)
	return res, err
}

func GetAllQuestionSets(limit int64) (*sql.Rows, error) {
	db := db.GetDB()

	query := `
        SELECT
            qs.id,
            qs.name,
            qs.interviewType,
            qs.employerId,
            qs.roleId,
            qs.created_at,
            qs.deleted,
			e.logo as logo,
            r.name AS role_name,
            e.name AS employer_name,
			i.id AS industry_id,
            i.name AS industry_name
        FROM QuestionSets AS qs
        JOIN Roles AS r ON qs.roleId = r.id
        JOIN Employers AS e ON qs.employerId = e.id
        JOIN Industries AS i ON e.industryId = i.id
		LIMIT $1
    `

	rows, err := db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func GetQuestionSetsByID(id int64) (*sql.Row, error) {
	db := db.GetDB()
	row := db.QueryRow("SELECT id, name, interviewType, employerId, roleId, created_at, deleted FROM QuestionSets WHERE id = $1", id)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return row, nil
}

func GetQuestionSetsByName(name string) (*sql.Row, error) {
	db := db.GetDB()
	row := db.QueryRow("SELECT id, name, interviewType, employerId, roleId, created_at, deleted FROM QuestionSets WHERE name = $1", name)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return row, nil
}

// UpdateQuestion updates an existing question in the database.
func UpdateQuestionSet(id int64, name, interviewType string, employerId, roleId int64) error {
	db := db.GetDB()
	_, err := db.Exec("UPDATE QuestionSets SET name = $1, interviewType = $2, employerId = $3, roleId = $4 WHERE id = $5", name, interviewType, employerId, roleId, id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteQuestion marks a question as deleted in the database.
func DeleteQuestionSet(id int64) error {
	db := db.GetDB()
	// Prepare and execute the SQL UPDATE statement to set the "deleted" flag to true
	_, err := db.Exec("UPDATE QuestionSets SET deleted = true WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteQuestionSetRow(id int64) error {
	db := db.GetDB()
	// Prepare and execute the SQL UPDATE statement to set the "deleted" flag to true
	_, err := db.Exec("DELETE FROM QuestionSets WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func CustomQueryForRow(query string, args ...any) (*sql.Row, error) {
	db := db.GetDB()
	row := db.QueryRow(query, args)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return row, nil
}
