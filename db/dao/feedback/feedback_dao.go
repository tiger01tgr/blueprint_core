package dao

import (
	"backend/db"
	"github.com/lib/pq"
	// "backend/db/models"
	"database/sql"
	// "fmt"
	// "log"
	"time"
)

func CreateFeedback(userId int64, questionSetId int64, practiceSessionId int64, questionIds []int64, videoUrls []string, feedbacks []string, created_at time.Time, seen bool) error {
    db := db.GetDB()
    tx, err := db.Begin()
    if err != nil {
        return err
    }

    stmt, err := tx.Prepare("INSERT INTO Feedback (userId, questionSetId, practiceSessionId, questionIds, videoUrls, feedbacks, created_at, seen) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(userId, questionSetId, practiceSessionId, pq.Array(questionIds), pq.Array(videoUrls), pq.Array(feedbacks), created_at, seen)
    if err != nil {
        tx.Rollback()
        return err
    }

    err = tx.Commit()
    return err
}

func GetFeedback(userId int64) (*sql.Rows, error){
	db := db.GetDB()
	rows, err := db.Query("SELECT id, userId, questionSetId, practiceSessionId, questionIds, videoUrls, feedbacks, created_at, seen FROM Feedback WHERE userId = $1", userId)
	return rows, err
}