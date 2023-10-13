package models

import (
	"fmt"
)

// Question represents the Questions table in the database.
type Question struct {
	ID            int64
	QuestionSetId int64
	Text          string
	TimeLimit     int64
}

func (q Question) String() string {
	return "Question{" +
		"id: " + fmt.Sprint(q.ID) +
		", questionSetId: " + fmt.Sprint(q.QuestionSetId) +
		", text: " + q.Text +
		", timeLimit: " + fmt.Sprint(q.TimeLimit) +
		"}"
}
