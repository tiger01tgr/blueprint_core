package models

import (
	"fmt"
)

type QuestionSet struct {
	id         uint64
	employerId uint64
	role       string
	interviewType string
	questions  []uint64
	created_at string
	deleted    bool
}

func (qs QuestionSet) String() string {
	return "QuestionSet{" +
		"id: " + fmt.Sprint(qs.id) +
		", employerId: " + fmt.Sprint(qs.employerId) +
		", role: " + qs.role +
		", interviewType: " + qs.interviewType +
		", questions: " + fmt.Sprint(qs.questions) +
		", created_at: " + qs.created_at +
		"}"
}