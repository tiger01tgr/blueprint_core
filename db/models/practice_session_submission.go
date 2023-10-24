package models

type PracticeSessionSubmission struct {
	ID                int64  `json:"id"`
	PracticeSessionId int64  `json:"practice_session_id"`
	QuestionId        int64  `json:"question_id"`
	Url               string `json:"url"`
}
