package dao

import (
	"backend/db"
	"database/sql"
)

// CreateJob inserts a new job into the database.
func CreateJob(name, employerId, jobTypeId, roleId, questionSetId, description string) error {
	db := db.GetDB()
	_, err := db.Exec(
		"INSERT INTO Jobs (name, employerId, jobTypeId, roleId, questionSetId, description) VALUES ($1, $2, $3, $4, $5, $6)",
		name,
		employerId,
		jobTypeId,
		roleId,
		questionSetId,
		description,
	)
	return err
}

func CreateJobType(name string) error {
	db := db.GetDB()
	_, err := db.Exec(
		"INSERT INTO JobTypes (name) VALUES ($1)",
		name,
	)
	return err
}

func GetAllJobs(offset, limit int64) (*sql.Rows, error) {
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
		OFFSET $2
    `

	rows, err := db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func GetJobTypes() (*sql.Rows, error) {
	db := db.GetDB()
	query := `
		SELECT
			id,
			name
		FROM JobTypes
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}