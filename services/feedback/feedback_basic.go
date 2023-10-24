package feedback

import (
	dao "backend/db/dao/feedback"
	"backend/db/models"
	sessionsService "backend/services/practice_sessions"
	"log"
	"time"

	"github.com/lib/pq"
)

func CreateFeedback(sessionId int64, submissionIds []int64, feedback []string) error {
	practiceSession, err := sessionsService.GetPracticeSessionWithId(sessionId)
	if err != nil {
		log.Print("Error getting practice session")
		log.Println(err.Error())
		return err
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

func GetFeedback(userId int64) ([]models.Feedback, error) {
	// Get feedback
	feedback, err := dao.GetFeedback(userId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var feedbacks []models.Feedback
	for feedback.Next() {
		var fb models.Feedback
		var questionIds []int64
		var videoUrls []string
		var feedbackArr []string
		err := feedback.Scan(&fb.ID, &fb.UserId, &fb.QuestionSetId, &fb.PracticeSessionId, pq.Array(&questionIds), pq.Array(&videoUrls), pq.Array(&feedbackArr), &fb.CreatedAt, &fb.Seen)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		fb.QuestionIds = questionIds
		fb.VideoUrls = videoUrls
		fb.Feedback = feedbackArr
		feedbacks = append(feedbacks, fb)
	}
	return feedbacks, nil
}
