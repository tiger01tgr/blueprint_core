package models

import (
	"time"
)

type PracticeSession struct {
	ID            int64
	QuestionSetId int64
	UserId        int64
	CurrentQuestionId int64
	LastAnsweredQuestionId int64
	Status		  string
	CompletedAt	  time.Time
}

func NewPracticeSession(userId int64, questionSetId int64) PracticeSession {
	return PracticeSession{
		QuestionSetId: questionSetId,
		UserId:        userId,
		CurrentQuestionId: 0,
		LastAnsweredQuestionId: 0,
		Status:		   "in_progress",
	}
}