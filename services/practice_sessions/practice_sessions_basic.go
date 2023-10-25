package practicesessions

import (
	dao "backend/db/dao/practice_sessions"
	questionsService "backend/services/questions"
	"backend/db/models"
	"backend/services/questions"
	"backend/services/s3"
	"database/sql"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
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

func GetPracticeSessionWithId(sessionId int64) (*models.PracticeSession, error) {
	row := dao.GetPracticeSessionWithId(sessionId)
	if row.Err() == sql.ErrNoRows {
		return nil, errors.New("Practice session does not exist")
	}
	var ps models.PracticeSession
	var lastAnsweredQuestionID sql.NullInt64
	var completedAt sql.NullTime
	err := row.Scan(&ps.ID, &ps.UserId, &ps.QuestionSetId, &ps.Status, &lastAnsweredQuestionID, &completedAt)
	return &ps, err
}

func GetCompletedPracticeSessions(page int64, limit int64) ([]models.PracticeSession, error) {
	offset := (page - 1) * limit
	rows, err := dao.GetCompletedPracticeSessions(offset, limit)
	defer rows.Close()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var practiceSessions []models.PracticeSession
	for rows.Next() {
		var ps models.PracticeSession
		var completedAt sql.NullTime
		err := rows.Scan(&ps.ID, &ps.UserId, &ps.QuestionSetId, &completedAt)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		if completedAt.Valid {
			ps.CompletedAt = completedAt.Time
		}
		practiceSessions = append(practiceSessions, ps)
	}
	return practiceSessions, nil
}

func GetCompletedPracticeSessionsPagination(limit int64) (int64, int64, error) {
	row, err := dao.GetNumberOfCompletedSessions()
	if err != nil {
		return 0, 0, err
	}
	var count int64
	row.Scan(&count)
	return count, (count / limit) + 1, nil
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
	if ps.Status == "completed" {
		return errors.New("Practice session already completed")
	}

	questionIdValidate, err := questions.IfQuestionIdExistsInQuestionSet(questionId, questionSetId)
	if err != nil {
		return err
	}
	if !questionIdValidate {
		return errors.New("Question id does not exist in question set")
	}

	var completedAt sql.NullTime

	// Check if the submitting question is the next question
	p, err := questions.GetNextQuestion(questionSetId, ps.LastAnsweredQuestionId)
	if err != nil {
		fmt.Println(err.Error())
		return err
	} else if p != nil && p.ID != questionId {
		// If the submitting question is not the next question, return an error
		fmt.Println("p.ID: ", p.ID)
		return errors.New("Question id does not match current question id")
	}
	// Passed questionId checks

	// Check if session is now complete
	p, err = questions.GetNextQuestion(questionSetId, questionId)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err.Error())
		return err
	} else if err == sql.ErrNoRows {
		fmt.Println(err.Error())
		ps.Status = "completed"
		ps.CompletedAt = time.Now()
		completedAt = sql.NullTime{Time: ps.CompletedAt, Valid: true}
	}
	err = dao.CreatePracticeSubmission(userId, questionSetId, practiceSessionId, questionId, videoUrl, ps.Status, completedAt)
	fmt.Println("hello from practice_sessions_basic.go")
	return err
}

func GetCompletedPracticeSessionSubmissions(sessionId int64) ([]models.PracticeSessionSubmission, error) {
	rows, err := dao.GetCompletedPracticeSessionSubmissions(sessionId)
	defer rows.Close()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var practiceSessionSubmissions []models.PracticeSessionSubmission
	for rows.Next() {
		var ps models.PracticeSessionSubmission
		err := rows.Scan(&ps.ID, &ps.PracticeSessionId, &ps.QuestionId, &ps.Url)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		question, err := questionsService.GetQuestionWithId(ps.QuestionId)
		ps.QuestionText = question.Text
		ps.TimeLimit = question.TimeLimit
		practiceSessionSubmissions = append(practiceSessionSubmissions, ps)
	}
	return practiceSessionSubmissions, nil
}


func uploadUserSubmissionVideo(video *multipart.File) (string, error) {
	uuid := uuid.New()
	return s3.UploadUserSubmissionVideo(uuid.String(), video)
}
