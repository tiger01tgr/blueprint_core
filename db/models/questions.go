package models

import (
	"fmt"
)

// Question represents the Questions table in the database.
type Question struct {
	ID            uint64
	QuestionSetId uint64
	Text          string
	TimeLimit     uint64
}

func (q Question) String() string {
	return "Question{" +
		"id: " + fmt.Sprint(q.ID) +
		", questionSetId: " + fmt.Sprint(q.QuestionSetId) +
		", text: " + q.Text +
		", timeLimit: " + fmt.Sprint(q.TimeLimit) +
		"}"
}
