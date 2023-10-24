package feedback

import (
	dao "backend/db/dao/feedback"
	"backend/db/models"
	sessionsService "backend/services/practice_sessions"
	"errors"
	"fmt"
	"log"
	"time"
)

func CreateFeedback(sessionId int64, submissionIds []int64, feedback []string) error {
	practiceSession, err := sessionsService.GetPracticeSessionWithId(sessionId)
	if err != nil {
		log.Print("Error getting practice session")
		log.Println(err.Error())
		return err
	}
	if practiceSession.Status != "completed" {
		log.Println("Practice session is in progress or closed")
		return errors.New("Practice session is in progress or closed")
	}
	questionSetId := practiceSession.QuestionSetId
	userId := practiceSession.UserId
	submissions, err := sessionsService.GetCompletedPracticeSessionSubmissions(sessionId)
	if err != nil {
		log.Print("Error getting submissions")
		log.Println(err.Error())
		return err
	}
	if len(submissions) != len(submissionIds) {
		log.Println("Number of submissions does not match number of submission ids")
		return err
	}
	questionIds := make([]int64, len(submissions))
	urls := make([]string, len(submissions))
	for i, submission := range submissions {
		questionIds[i] = submission.QuestionId
		urls[i] = submission.Url
	}
	err = dao.CreateFeedback(userId, questionSetId, sessionId, questionIds, urls, feedback, time.Now(), false)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func GetAllFeedback(userId int64) ([]models.Feedback, error) {
	// Get feedback
	feedback, err := dao.GetAllFeedback(userId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var feedbacks []models.Feedback
	for feedback.Next() {
		var fb models.Feedback
		err := feedback.Scan(&fb.ID, &fb.UserId, &fb.QuestionSetId, &fb.PracticeSessionId, &fb.CreatedAt, &fb.Seen)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		feedbacks = append(feedbacks, fb)
	}
	return feedbacks, nil
}

func MarkFeedbackAsRead(feedbackId int64) {
	err := dao.MarkFeedbackAsRead(feedbackId)
	if err != nil {
		log.Println(err.Error())
	}
}

func GetFeedback(userId int64, feedbackId int64) (*models.Feedback, error) {
	row, err := dao.GetFeedback(feedbackId)
	var feedback models.Feedback
	err = row.Scan(&feedback.ID, &feedback.UserId, &feedback.QuestionSetId, &feedback.PracticeSessionId, &feedback.CreatedAt, &feedback.Seen)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if feedback.UserId != userId {
		log.Println("User does not have access to this feedback")
		return nil, err
	}
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &feedback, nil

}

func GetFeedbackEntries(userId int64, feedbackId int64) ([]models.FeedbackEntries, error) {
	if err := isAuthorized(userId, feedbackId); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	rows, err := dao.GetFeedbackEntries(feedbackId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var feedbackEntries []models.FeedbackEntries
	for rows.Next() {
		var fb models.FeedbackEntries
		err := rows.Scan(&fb.ID, &fb.FeedbackId, &fb.QuestionId, &fb.VideoUrl, &fb.Feedback)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		feedbackEntries = append(feedbackEntries, fb)
	}
	return feedbackEntries, nil
}

func isAuthorized(userId int64, feedbackId int64) error {
	feedback, err := GetFeedback(userId, feedbackId)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	if feedback == nil {
		return errors.New(fmt.Sprintf("Feedback id=%d not found", feedbackId))
	}
	if feedback.UserId != userId {
		return errors.New("Access Denied")
	}
	return nil
}
