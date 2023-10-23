package practicesessions

import (
	dao "backend/db/dao/practice_sessions"
	"backend/db/models"
	"backend/services/questions"
	"database/sql"
	"errors"
	"mime/multipart"
	"backend/services/s3"
	"github.com/google/uuid"
	"time"
	"fmt"
)

func CreatePracticeSession(userId int64, questionSetId int64) error {
	row := dao.GetPracticeSession(userId, questionSetId)
	var ps models.PracticeSession
	err := row.Scan(&ps.ID, &ps.UserId, &ps.QuestionSetId, &ps.Status, &ps.LastAnsweredQuestionId, &ps.CompletedAt)
	if err != nil && err != sql.ErrNoRows {
		return err
	} else if err == nil && ps.Status == "in_progress" {
		return errors.New("User and Practice set pair already exists")
	}
	ps = models.NewPracticeSession(userId, questionSetId)
	return dao.CreatePracticeSession(ps.UserId, ps.QuestionSetId, ps.Status)
}

func GetPracticeSession(userId int64, questionSetId int64) (*models.PracticeSession, error) {
	row := dao.GetPracticeSession(userId, questionSetId)
	if row.Err() == sql.ErrNoRows {
		return nil, errors.New("Practice session does not exist")
	}
	var ps models.PracticeSession
	var lastAnsweredQuestionID sql.NullInt64
	var completedAt sql.NullTime
	err := row.Scan(&ps.ID, &ps.UserId, &ps.QuestionSetId, &ps.Status, &lastAnsweredQuestionID, &completedAt)
	if err != nil {
		return nil, err
	}
	if lastAnsweredQuestionID.Valid {
		ps.LastAnsweredQuestionId = lastAnsweredQuestionID.Int64
		question, err := questions.GetNextQuestion(ps.QuestionSetId, ps.LastAnsweredQuestionId)
		if err != nil {
			return nil, err
		}
		ps.CurrentQuestionId = question.ID
	} else {
		ps.LastAnsweredQuestionId = 0
		ps.CurrentQuestionId = 0
	}
	if completedAt.Valid {
		ps.CompletedAt = completedAt.Time
	}
	return &ps, err
}

func CreatePracticeSubmission(userId int64, questionSetId int64, practiceSessionId int64, questionId int64, video *multipart.File) error {
	videoUrl, err := uploadUserSubmissionVideo(video)
	if err != nil {
		return err
	}
	ps, err := GetPracticeSession(userId, questionSetId)
	if err != nil {
		return errors.New("No practice session found with user and question set pair")
	}
	if ps.ID != practiceSessionId {
		return errors.New("Practice session id does not exist")
	}
	questionIdValidate, err := questions.IfQuestionIdExistsInQuestionSet(questionId, questionSetId)
	if err != nil {
		return err
	}
	if !questionIdValidate {
		return errors.New("Question id does not exist in question set")
	}
	if ps.Status == "completed" {
		return errors.New("Practice session already completed")
	}

	var completedAt sql.NullTime
	p, err := questions.GetNextQuestion(questionSetId, ps.LastAnsweredQuestionId)
	if err != nil && err != sql.ErrNoRows {
		return err
	} else if err == sql.ErrNoRows {
		ps.Status = "completed"
		ps.CompletedAt = time.Now()
		completedAt = sql.NullTime{Time: ps.CompletedAt, Valid: true}
	} else if p.ID != questionId {
		fmt.Println("p.ID: ", p.ID)
		return errors.New("Question id does not match current question id")
	}
	err = dao.CreatePracticeSubmission(userId, questionSetId, practiceSessionId, questionId, videoUrl, ps.Status, completedAt)
	fmt.Println("hello from practice_sessions_basic.go")
	return err
}

func uploadUserSubmissionVideo(video *multipart.File) (string, error) {
	uuid := uuid.New()
	return s3.UploadUserSubmissionVideo(uuid.String(), video)
}