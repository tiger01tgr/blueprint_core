package dao

import (
	"backend/db"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

// CreatePracticeSession inserts a new practice session into the database.
func CreatePracticeSession(userId int64, questionSetId int64, status string) error {
	db := db.GetDB()
	_, err := db.Exec(
		"INSERT INTO PracticeSessions (userId, questionSetId, status) VALUES ($1, $2, $3)",
		userId,
		questionSetId,
		status,
	)
	return err
}

func GetPracticeSession(userId int64, questionSetId int64) *sql.Row {
	db := db.GetDB()
	row := db.QueryRow(
		"SELECT id, userId, questionSetId, status, lastAnsweredQuestionId, completedAt FROM PracticeSessions WHERE userId = $1 AND questionSetId = $2 AND status = 'in_progress'",
		userId,
		questionSetId,
	)
	return row
}

func CreatePracticeSubmission(userId int64, questionSetId int64, practiceSessionId int64, questionId int64, videoUrl string, status string, completedAt sql.NullTime) error {
    db := db.GetDB()

    // Start a transaction
    tx, err := db.Begin()
    if err != nil {
        return errors.Wrap(err, "Failed to start transaction")
    }

    // Defer a function to handle the transaction based on the success or failure of the operations
    defer func() {
        if p := recover(); p != nil {
            // A panic occurred, so we should rollback the transaction
			fmt.Println("Panic occurred 1")
            tx.Rollback()
        } else if err != nil {
            // An error occurred, so we should rollback the transaction
			fmt.Println("Panic occurred 2")
            tx.Rollback()
        } else {
            // All operations were successful, so we commit the transaction
			fmt.Println()
            err = tx.Commit()
            if err != nil {
                // Failed to commit the transaction
				fmt.Println("Panic occurred 3")
                err = errors.Wrap(err, "Failed to commit transaction")
            }
			fmt.Println("Success!")
        }
    }()

    // Insert a record into PracticeSessionSubmissions
    _, err = tx.Exec(
        "INSERT INTO PracticeSessionSubmissions (practiceSessionId, questionId, url) VALUES ($1, $2, $3)",
        practiceSessionId,
        questionId,
        videoUrl,
    )
    if err != nil {
        return errors.Wrap(err, "Failed to insert into PracticeSessionSubmissions")
    }
	if completedAt.Valid {
		// Update PracticeSession table here
		_, err = tx.Exec(
			"UPDATE PracticeSessions SET lastAnsweredQuestionId = $1, status = $2, completedAt = $3 WHERE id = $4",
			questionId,
			status,
			completedAt.Time,
			practiceSessionId,
		)
		fmt.Println(status)
	} else {
		_, err = tx.Exec(
			"UPDATE PracticeSessions SET lastAnsweredQuestionId = $1, status = $2 WHERE id = $3",
			questionId,
			status,
			practiceSessionId,
		)
		fmt.Println(status)
	}
	if err != nil {
		return errors.Wrap(err, "Failed to update PracticeSession")
	}

    return nil // If everything is successful, the defer block will handle the transaction commit
}