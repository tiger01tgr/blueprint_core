package models

import "time"

type Feedback struct {
	ID                int64     `json:"id"`
	UserId            int64     `json:"user_id"`
	QuestionSetId     int64     `json:"question_set_id"`
	QuestionSetName   string    `json:"question_set_name"`
	PracticeSessionId int64     `json:"practice_session_id"`
	EmployerName      string    `json:"employer_name"`
	EmployerLogo      string    `json:"employer_logo"`
	CreatedAt         time.Time `json:"created_at"`
	Seen              bool      `json:"seen"`
}

type FeedbackEntries struct {
	ID           int64  `json:"id"`
	FeedbackId   int64  `json:"feedback_id"`
	QuestionId   int64  `json:"question_id"`
	QuestionText string `json:"question_text"`
	TimeLimit    int64  `json:"time_limit"`
	VideoUrl     string `json:"video_url"`
	Feedback     string `json:"feedback"`
}
