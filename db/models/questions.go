package models

import (
	"fmt"
)

// Question represents the Questions table in the database.
type Question struct {
	id         uint64
	questionSetId uint64
	text       string
	timeLimit  uint64
	deleted    bool
}

func (q Question) String() string {
	return "Question{" +
			"id: " + fmt.Sprint(q.id) +
			", questionSetId: " + fmt.Sprint(q.questionSetId) +
			", text: " + q.text +
			", timeLimit: " + fmt.Sprint(q.timeLimit) +
			", deleted: " + fmt.Sprint(q.deleted) +
		"}"
}