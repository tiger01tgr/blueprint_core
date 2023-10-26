package dao

import (
	"backend/db"
	// "backend/db/models"
	"database/sql"
	// "fmt"
	"log"
	"time"

	"github.com/pkg/errors"
)

func CreateFeedback(userId int64, questionSetId int64, practiceSessionId int64, questionIds []int64, videoUrls []string, feedbacks []string, created_at time.Time, seen bool) error {
	db := db.GetDB()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// Defer a function to handle the transaction based on the success or failure of the operations
	defer func() {
		if p := recover(); p != nil {
			// A panic occurred, so we should rollback the transaction
			log.Println("Panic occurred 1")
			tx.Rollback()
		} else if err != nil {
			// An error occurred, so we should rollback the transaction
			log.Println("Panic occurred 2")
			tx.Rollback()
		} else {
			// All operations were successful, so we commit the transaction
			err = tx.Commit()
			if err != nil {
				// Failed to commit the transaction
				log.Println("Panic occurred 3")
				err = errors.Wrap(err, "Failed to commit transaction")
			}
			log.Println("Success!")
		}
	}()

	// Insert into Feedback
	var feedbackID int64
	err = tx.QueryRow("INSERT INTO Feedback (userId, questionSetId, practiceSessionId, created_at, seen) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		userId, questionSetId, practiceSessionId, created_at, seen).Scan(&feedbackID)
	if err != nil {
		return errors.Wrap(err, "Failed to insert into Feedback or returning id")
	}

	// Insert into FeedbackEntries
	for i := range questionIds {
		_, err = tx.Exec("INSERT INTO FeedbackEntries (feedbackId, questionId, videoUrl, feedbackText) VALUES ($1, $2, $3, $4)",
			feedbackID, questionIds[i], videoUrls[i], feedbacks[i])
		if err != nil {
			return errors.Wrap(err, "Failed to insert into FeedbackEntries")
		}
	}

	// Mark practice session as closed
	_, err = tx.Exec("UPDATE PracticeSessions SET status = 'closed' WHERE id = $1", practiceSessionId)
	if err != nil {
		return errors.Wrap(err, "Failed to update PracticeSessions")
	}

	return nil // If everything is successful, the defer block will handle the transaction commit

}

func GetAllFeedback(userId int64) (*sql.Rows, error) {
	db := db.GetDB()
	rows, err := db.Query(
		`SELECT 
			Feedback.id, 
			Feedback.userId, 
			Feedback.questionSetId, 
			Feedback.practiceSessionId, 
			Feedback.created_at, 
			Feedback.seen,
			QuestionSets.name AS questionset_name,
			QuestionSets.interviewType as interview_type,
			Employers.name AS employer_name,
    		Employers.logo AS employer_logo
			FROM Feedback
			JOIN QuestionSets ON Feedback.questionSetId = QuestionSets.id
			JOIN Employers ON QuestionSets.employerId = Employers.id
			WHERE userId = $1 
			ORDER BY created_at ASC`,
		userId)
	return rows, err
}

func GetFeedback(feedbackId int64) (*sql.Row, error) {
	db := db.GetDB()
	row := db.QueryRow("SELECT id, userId, questionSetId, practiceSessionId, created_at, seen FROM Feedback WHERE id = $1", feedbackId)
	return row, nil
}

func GetFeedbackEntries(feedbackId int64) (*sql.Rows, error) {
	db := db.GetDB()
	rows, err := db.Query(`
		SELECT 
			FeedbackEntries.id, 
			FeedbackEntries.feedbackId, 
			FeedbackEntries.questionId, 
			FeedbackEntries.videoUrl, 
			FeedbackEntries.feedbackText,
			Questions.text as question_text,
			Questions.timeLimit as question_time_limit
		FROM FeedbackEntries 
		JOIN Questions ON FeedbackEntries.questionId = Questions.id
		WHERE feedbackId = $1`,
		feedbackId)
	return rows, err
}

func MarkFeedbackAsRead(feedbackId int64) error {
	db := db.GetDB()
	_, err := db.Exec(
		"UPDATE Feedback SET seen = true WHERE id = $1",
		feedbackId,
	)
	return err
}
