package models

import "time"

type Feedback struct {
	ID                int64     `json:"id"`
	UserId            int64     `json:"user_id"`
	QuestionSetId     int64     `json:"question_set_id"`
	PracticeSessionId int64     `json:"practice_session_id"`
	QuestionIds       []int64   `json:"question_ids"`
	VideoUrls         []string  `json:"url"`
	Feedback          []string  `json:"feedback"`
	CreatedAt         time.Time `json:"created_at"`
	Seen              bool      `json:"seen"`
}
